package fncPsglst

import (
	fncApndix "back/apndix/function"
	mdlApndix "back/apndix/model"
	mdlPsglst "back/psglst/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

func FncPslgstRsvpnrMainpg(psglst mdlPsglst.MdlPsglstPsgdtlDtbase,
	sycClrpsg, sycNulpsg, sycPnrcde, sycChrter, sycMilege *sync.Map,
	mapCurrcv map[string]mdlApndix.MdlApndixCurrcvDtbase,
	sycWgroup *sync.WaitGroup, objtkn mdlSbrapi.MdlSbrapiMsghdrParams,
	airlfl, pnrcde, lstvar string) {
	var cekLstvar, cekChrter, cekIsflwn, cekNonrev, cekTcktng bool
	if sycWgroup != nil {
		defer sycWgroup.Done()
	}

	// DEBUGING
	if lstvar == "scd" || lstvar == "fst" {
		psglst.Source += "|INTERLINE"
		if lstvar == "scd" {
			cekLstvar = true
		}
	}
	if objtkn.Bsttkn == "" {
		psglst.Source += "|TOKEN NIL" + airlfl
	}

	// Check on name XXDHC pax non rev and isit flown
	cekIsflwn = psglst.Isitfl == "F"
	if psglst.Nmefst == "XXDHC" || psglst.Nmelst == "XXDHC" ||
		psglst.Nmefst == "XXSNY" || psglst.Nmelst == "XXSNY" ||
		psglst.Nmefst == "XDHC" || psglst.Nmelst == "XDHC" {
		psglst.Isitnr = "CREW"
		cekNonrev = true
	}

	// Get ticketing from PNR
	slcSbarea := []string{"ITINERARY", "RECORD_LOCATOR", "ANCILLARY", "TICKETING"}

	// Get reservation
	nowRsvpnr := mdlSbrapi.MdlSbrapiRsvpnrRsprsv{}
	if istTcktng, ist := sycPnrcde.Load(pnrcde + airlfl); ist {
		if mtcTcktng, mtc := istTcktng.(mdlSbrapi.MdlSbrapiRsvpnrRsprsv); mtc {
			nowRsvpnr = mtcTcktng
		}
	} else {
		getTcktng, err := fncSbrapi.FncSbrapiRsvpnrMainob(objtkn, pnrcde, slcSbarea)
		if err != nil {
			return
		}
		nowRsvpnr = getTcktng
		sycPnrcde.Store(pnrcde+airlfl, getTcktng)
	}

	// If data not null
	if nowRsvpnr.BookingDetails.RecordLocator != "" {

		// Date formating PNR book PNR Create date
		varTimecr := nowRsvpnr.BookingDetails.SystemCreationTimestamp
		if pnrTimecr, err := time.Parse("2006-01-02T15:04:05", varTimecr); err == nil {
			rawTimerw, _ := strconv.Atoi(pnrTimecr.Format("0601021504"))
			nowAgtdcr := nowRsvpnr.BookingDetails.CreationAgentID
			psglst.Timecr = int64(rawTimerw)
			if nowAgtdcr != "" {
				psglst.Agtdcr = nowAgtdcr
			}
		}

		// Get PNR interline
		objPnritl := nowRsvpnr.POS.Source.TTYRecordLocator
		if objPnritl.RecordLocator != "" {
			nowPnritl := objPnritl.CRSCode + "*" + objPnritl.RecordLocator
			psglst.Pnritl = fncApndix.FncApndixUpdateSlcstr(&psglst.Pnritl, nowPnritl)
		}

		// Get PNR interline and itinerary
		slcItinry := nowRsvpnr.PassengerReservation.Segments.Segment
		slcSegpnr := []string{}
		slcRoutsg, lstArrivl := []string{}, ""
		if len(slcItinry) != 0 {
			for idx, itinry := range slcItinry {
				if !slices.Contains([]string{"JT", "ID", "IW", "IU", "OD", "SL"},
					itinry.Air.OperatingAirlineCode) {
					continue
				}

				// PNR Interline
				rawPnritl := itinry.Air.AirlineRefId
				if len(rawPnritl) > 5 {
					psglst.Pnritl = fncApndix.FncApndixUpdateSlcstr(&psglst.Pnritl, rawPnritl[2:])
				}

				// Get time flown
				rawTimefl := itinry.Air.DepartureDateTime
				fmtTimefl, _ := time.Parse("2006-01-02T15:04:05", rawTimefl)
				strTimefl := fmtTimefl.Format("0601021504")

				// Itinerary segment
				rawDepart := itinry.Air.DepartureAirport
				rawArrivl := itinry.Air.ArrivalAirport
				rawActncd := itinry.Air.ActionCode
				mktAirlfl := itinry.Air.MarketingAirlineCode
				optAirlfl := itinry.Air.OperatingAirlineCode
				mktFlnbfl := itinry.Air.MarketingFlightNumber
				optFlnbfl := itinry.Air.OperatingFlightNumber
				mktClssfl := itinry.Air.MarketingClassOfService
				optClssfl := itinry.Air.OperatingClassOfService
				fmtSegpnr := fmt.Sprintf("%s-%s-%s-%s-MKT-%s-%s-%s-OPT-%s-%s-%s",
					rawDepart, rawArrivl, rawActncd, strTimefl,
					mktAirlfl, mktFlnbfl, mktClssfl,
					optAirlfl, optFlnbfl, optClssfl)
				lstArrivl = rawArrivl
				if len(slcRoutsg) == 0 ||
					(idx >= 1 && slcRoutsg[len(slcRoutsg)-1] != rawDepart) {
					slcSegpnr = append(slcSegpnr, fmtSegpnr)
					slcRoutsg = append(slcRoutsg, rawDepart)
				}
			}
			slcRoutsg = append(slcRoutsg, lstArrivl)
			psglst.Routsg = strings.Join(slcRoutsg, "-")
			psglst.Segpnr = strings.Join(slcSegpnr, "|")
		}

		// Get ticketing detail for issued date
		var slcTcktng = nowRsvpnr.PassengerReservation.TicketingInfo.TicketDetails
		var mapEmdnae = map[string]bool{}
		var getTktnvc = ""
		if len(slcTcktng) != 0 {
			for _, tcktng := range slcTcktng {

				// Logical gate for ticket number
				strFmtnme := (psglst.Nmelst + "     ")[:5]
				strLstnme := (psglst.Nmefst + " ")[:1]
				cncFulln1 := strFmtnme + "/" + strLstnme
				cncFulln2 := psglst.Nmelst + "/" + strLstnme
				if cncFulln1 == tcktng.PassengerName ||
					cncFulln2 == tcktng.PassengerName {

					// Get ticket number blank and emd
					if tcktng.TicketNumber[3:4] != "4" {
						getTktnvc = tcktng.TicketNumber[:13]
						if psglst.Psgrid == "FB79CD940001" {
							fmt.Println(pnrcde, tcktng.TicketNumber[:13])
						}
					} else if tcktng.TicketNumber[3:4] == "4" {
						mapEmdnae[tcktng.TicketNumber[:13]] = true
					}
				}
			}
		}
		if getTktnvc != "" {
			psglst.Tktnvc = getTktnvc
		}

		// Get ancillary
		if len(nowRsvpnr.OpenReservationElements) > 0 {
			for _, elm := range nowRsvpnr.OpenReservationElements {
				delete(mapEmdnae, psglst.Emdnae)
				minNmefst := math.Min(20, float64(len(elm.NameAssociationList.FirstName)))
				minNmelst := math.Min(30, float64(len(elm.NameAssociationList.LastName)))
				if elm.ActionCode != "HI" ||
					elm.NameAssociationList.FirstName[:int(minNmefst)] != psglst.Nmefst ||
					elm.NameAssociationList.LastName[:int(minNmelst)] != psglst.Nmelst {
					continue
				}

				// Looping segment assoc
				nowRoutae := ""
				for _, cpn := range elm.SegmentAssociationList {
					nowDepart, nowArrivl := cpn.BoardPoint, cpn.OffPoint
					cpnTimefl, _ := time.Parse("2006-01-02", cpn.DepartureDate)
					nowTimefl, _ := time.Parse("0601021504", strconv.Itoa(int(psglst.Timefl)))
					difTimefl := math.Abs(cpnTimefl.Sub(nowTimefl).Hours() / 24)

					// Compare to flown data
					if (nowDepart == psglst.Depart || nowArrivl == psglst.Arrivl) && difTimefl < 1 {
						nowRoutae = cpn.BoardPoint + "-" + cpn.OffPoint
					}

					// Compare to route vcr
					if len(psglst.Routvc) >= 7 {
						if (nowDepart == psglst.Routvc[:3] || nowArrivl == psglst.Routvc[4:]) && difTimefl < 1 {
							nowRoutae = cpn.BoardPoint + "-" + cpn.OffPoint
						}
					}
				}

				// Push final data if assoc
				nowPaidbt := 1
				regDescae := regexp.MustCompile(`\d+K|\d+ K`)
				rslDescae := regDescae.FindAllString(elm.CommercialName, -1)
				if len(rslDescae) > 0 {
					regDescae := regexp.MustCompile(`\d+`)
					rslDescae := regDescae.FindAllString(rslDescae[0], -1)
					rawPaidbt := rslDescae[0]
					intPaidbt, _ := strconv.Atoi(rawPaidbt)
					nowPaidbt = intPaidbt * elm.NumberOfItems
				}

				// If get route assoc
				if nowRoutae != "" {
					fncApndix.FncApndixUpdateSlcstr(&psglst.Gpcdae, elm.GroupCode)
					fncApndix.FncApndixUpdateSlcstr(&psglst.Sbcdae, elm.RficSubcode)
					fncApndix.FncApndixUpdateSlcstr(&psglst.Descae, elm.CommercialName)
					psglst.Wgbgae += int32(nowPaidbt)
					psglst.Qtbgae += int32(elm.NumberOfItems)
					psglst.Routae = nowRoutae

					// Get emd number
					fncApndix.FncApndixUpdateSlcstr(&psglst.Emdnae, elm.EMDNumber)
					if elm.EMDNumber != "" {
						for _, tps := range elm.TravelPortions {
							if tps.BoardPoint+"-"+tps.OffPoint == nowRoutae {
								fncApndix.FncApndixUpdateSlcstr(&psglst.Emdnae, tps.EMDNumber)
							}
						}
					}

					// Fare manage
					psglst.Currae = elm.OriginalBasePrice.Currency
					if psglst.Currae != "IDR" {
						if vlx, ist := mapCurrcv[psglst.Currae]; ist {
							cnvFareae := elm.OriginalBasePrice.Price / vlx.Crrate
							psglst.Fareae += cnvFareae
						}
					} else {
						psglst.Fareae += elm.OriginalBasePrice.Price
					}

					// Check group code
					fstGroupc := elm.GroupCode == "BG"
					scdGroupc := elm.GroupCode == "UP"
					trdGroupc := elm.GroupCode == "TS" && psglst.Airlfl != "SL"
					if fstGroupc || scdGroupc || trdGroupc {
						if len(rslDescae) > 0 {
							psglst.Paidbt += int32(nowPaidbt)
						} else {
							psglst.Paidbt += int32(elm.NumberOfItems)
						}
					}

				}
			}

			// Cek emd non exist
			if len(mapEmdnae) > 0 {
				fncApndix.FncApndixUpdateSlcstr(&psglst.Noterr, "EMD DETAIL NIL")
			}
		} else if len(mapEmdnae) > 0 {
			fncApndix.FncApndixUpdateSlcstr(&psglst.Noterr, "EMD DETAIL NIL")
		}
	}

	if psglst.Pnritl == "" {
		cekLstvar = true
	}

	// Get ticketing document
	if psglst.Psgrid == "FB79CD940001" {
		fmt.Println("Tktnvc:", psglst.Tktnvc, "pnrcde:", pnrcde, "airlfl:", airlfl)
	}
	if psglst.Tktnvc != "" || (cekLstvar && psglst.Tktnfl != "") {
		getTktnow := psglst.Tktnvc
		if psglst.Tktnvc == "" {
			getTktnow = psglst.Tktnfl
		}
		if psglst.Psgrid == "FB79CD940001" {
			fmt.Println("masuk")
		}
		err := fncSbrapi.FncSbrapiGettktMainob(objtkn, airlfl, &psglst, getTktnow, mapCurrcv)
		if err != nil {
			fncApndix.FncApndixUpdateSlcstr(&psglst.Noterr, err.Error())
		}

		// Split farecalc and check non revenue
		tmpNonrev := FncPsglstFrcalcSplitd(&psglst, mapCurrcv, sycMilege, objtkn)
		if tmpNonrev {
			cekNonrev = true
			if psglst.Isitnr == "" {
				psglst.Isitnr = "ZEROFB"
			}
		}
		psglst.Source += "|GETTKT"
		if psglst.Statvc == "OK" || psglst.Statvc == "USED" {
			cekTcktng = true
			if psglst.Tktnvc == "" {
				psglst.Tktnvc = psglst.Tktnfl
			}
		}
	}

	// Check if data clear or not
	if cekChrter || !cekIsflwn || cekNonrev || cekLstvar || cekTcktng {
		istStlerr := true
		mapSuberr := map[string]bool{}
		fncFnlcek := func(params any, noterr, suberr string) {
			valFloatx, mtcFloatx := params.(float64)
			valString, mtcString := params.(string)
			if (mtcFloatx && valFloatx < 1000) || (mtcString && valString == "") {
				istStlerr = false
				mapSuberr[suberr] = true
				fncApndix.FncApndixUpdateSlcstr(&psglst.Noterr, noterr+" NIL")
			}
		}

		// Final check tktnfl
		fncFnlcek(psglst.Tktnvc, "TKTNVC", "MNFEST")
		fncFnlcek(psglst.Pnrcde, "PNRCDE", "MNFEST")
		fncFnlcek(psglst.Timeis, "TIMEIS", "MNFEST")
		fncFnlcek(psglst.Routvc, "ROUTVC", "MNFEST")
		fncFnlcek(psglst.Ntaffl, "NTAFFl", "SLSRPT")
		fncFnlcek(psglst.Ntafvc, "NTAFVC", "SLSRPT")
		fncFnlcek(psglst.Curncy, "CURNCY", "SLSRPT")
		if !istStlerr && cekIsflwn {
			for suberr := range mapSuberr {
				if suberr == "SLSRPT" && !cekNonrev {
					psglst.Slsrpt = "NOT CLEAR"
				}
				if suberr == "MNFEST" {
					psglst.Mnfest = "NOT CLEAR"
				}
			}
		}
		sycClrpsg.Store(psglst.Prmkey, psglst)
	} else {
		sycNulpsg.Store(psglst.Prmkey, psglst)
	}
}
