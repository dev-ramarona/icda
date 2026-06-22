package fncJeddah

import (
	fncApndix "back/apndix/function"
	mdlJeddah "back/jeddah/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Running process hit passanggerlist daily
func FncJeddahPrcessMainpg(c *gin.Context) {

	// protect single run
	if fncApndix.Status.Sbrapi != 0.0 {
		return
	}

	// Bind JSON Body input to variable
	inpErrlog := mdlJeddah.MdlJeddahPramsInputx{} //save
	if err := c.BindJSON(&inpErrlog); err != nil {
		panic(err)
	}
	fncApndix.Status.Sbrapi = 0.01
	var mapFlnbls = FncJeddahFlnblsMapobj()
	var prvPnrobj = FncJeddahPnrobjMapobj()
	var sycFlnbls = &sync.Map{}
	for airlfl, slices := range mapFlnbls {
		totWorker := 1
		var sycPnrcde = &sync.Map{}
		var sycWgroup sync.WaitGroup
		slcRspssn, err := fncSbrapi.FncSbrapiCrtssnMultpl(airlfl, int(totWorker))
		if err != nil {
			fmt.Println("airlfl" + airlfl + "failed")
			continue
		}

		jobPnrtrc := make(chan mdlJeddah.MdlJeddahFlnblsDtbase, 1500)
		for i := 0; i < int(totWorker); i++ {
			if len(slcRspssn) >= i+1 {
				if slcRspssn[i].Bsttkn != "" {
					sycWgroup.Add(1)
					fmt.Println("Success Token-", i)
					go FncPnrtrcPrcessWorker(&sycWgroup, jobPnrtrc, sycPnrcde, sycFlnbls, slcRspssn[i], prvPnrobj, slices)
					continue
				}
				fmt.Println("Failed Token-", i)
			}
		}

		for _, object := range slices {
			if inpErrlog.Flnbfl_jeddah == "" || inpErrlog.Flnbfl_jeddah == object.Flnbfl {
				jobPnrtrc <- object
			}
		}

		close(jobPnrtrc)
		sycWgroup.Wait()
	}

	// Done
	fncApndix.Status.Sbrapi = 0
	c.JSON(200, gin.H{"status": "Done"})
}

func FncPnrtrcPrcessWorker(sycWgroup *sync.WaitGroup,
	jobPnrtrc <-chan mdlJeddah.MdlJeddahFlnblsDtbase, sycPnrcde, sycFlnbls *sync.Map,
	nowObjtkn mdlSbrapi.MdlSbrapiMsghdrParams, prvPnrobj map[string]mdlJeddah.MdlJeddahPnrsmrCmpare,
	slices map[string]mdlJeddah.MdlJeddahFlnblsDtbase) {
	defer sycWgroup.Done()
	var mgoFlnbls []mongo.WriteModel
	var mgoPnrtrc []mongo.WriteModel
	for slcPnrtrc := range jobPnrtrc {
		var objParams = mdlSbrapi.MdlSbrapiMsghdrApndix{
			Airlfl: slcPnrtrc.Airlfl, Datefl: slcPnrtrc.Datefl, Depart: slcPnrtrc.Depart,
			Flnbfl: slcPnrtrc.Flnbfl}
		var mapPnrtrc = map[string]mdlJeddah.MdlJeddahPnrsmrDtbase{}
		rspIssued, errIssued := fncSbrapi.FncSbrapiPnrtrcMainob("LC", nowObjtkn, objParams)
		if errIssued == nil {

			//
			rspBooked, errBooked := fncSbrapi.FncSbrapiPnrtrcMainob("PUN", nowObjtkn, objParams)
			if errBooked == nil {
				for pnrcde, val := range rspIssued {
					if _, ist := mapPnrtrc[val.Pnrcde]; !ist {
						mapPnrtrc[val.Pnrcde] = mdlJeddah.MdlJeddahPnrsmrDtbase{
							Pnrcde: val.Pnrcde,
							Pnrsrc: val.Pnrcde,
							Totpax: val.Totpax}
					}
					if _, ist := rspBooked[pnrcde]; !ist && val.Issued != "" {
						objPnrtrc := mapPnrtrc[val.Pnrcde]
						objPnrtrc.Timefl = val.Timefl
						objPnrtrc.Routsg = val.Routfl
						objPnrtrc.Flnbsg = objParams.Airlfl + "-" + val.Flnbfl
						objPnrtrc.Clssbk = val.Clssbk
						objPnrtrc.Agtnme = val.Agtnme
						objPnrtrc.Totisd = val.Totpax
						mapPnrtrc[val.Pnrcde] = objPnrtrc
					} else {
						objPnrtrc := mapPnrtrc[val.Pnrcde]
						objPnrtrc.Timefl = val.Timefl
						objPnrtrc.Routsg = val.Routfl
						objPnrtrc.Flnbsg = objParams.Airlfl + "-" + val.Flnbfl
						objPnrtrc.Clssbk = val.Clssbk
						objPnrtrc.Agtnme = val.Agtnme
						objPnrtrc.Totbok = val.Totpax
						mapPnrtrc[val.Pnrcde] = objPnrtrc
					}
				}
			}

			rspCancel, errCancel := fncSbrapi.FncSbrapiPnrtrcMainob("LX", nowObjtkn, objParams)
			if errCancel == nil {
				for pnrcde, val := range rspCancel {
					if _, ist := mapPnrtrc[val.Pnrcde]; !ist {
						mapPnrtrc[val.Pnrcde] = mdlJeddah.MdlJeddahPnrsmrDtbase{
							Pnrcde: val.Pnrcde,
							Pnrsrc: val.Pnrcde,
							Totpax: val.Totpax}
					}
					if _, ist := rspIssued[pnrcde]; !ist {
						objPnrtrc := mapPnrtrc[val.Pnrcde]
						objPnrtrc.Timefl = val.Timefl
						objPnrtrc.Routsg = val.Routfl
						objPnrtrc.Flnbsg = objParams.Airlfl + "-" + val.Flnbfl
						objPnrtrc.Clssbk = val.Clssbk
						objPnrtrc.Agtnme = val.Agtnme
						objPnrtrc.Totcxl = val.Totpax
						mapPnrtrc[val.Pnrcde] = objPnrtrc
					}
				}
			}
		}

		//
		nowswg := sync.WaitGroup{}
		for pnrcde, pnrObject := range mapPnrtrc {
			if _, ist := sycPnrcde.Load(pnrcde); !ist {
				nowswg.Add(1)
				go func() {
					defer nowswg.Done()
					sycPnrcde.Store(pnrcde, true)
					slcSbarea, idx := []string{"ITINERARY", "REMARKS"}, 0
					if pnrObject.Totcxl > 0 {
						slcSbarea = []string{}
					}
					getRsvpnr, err := fncSbrapi.FncSbrapiRsvpnrMainob(nowObjtkn, pnrcde, slcSbarea)
					if err == nil {

						// Get divided
						pnrTotori := pnrObject.Totpax
						for _, slices := range getRsvpnr.BookingDetails.DivideSplitDetails.Itemslice {
							valNamevl := slices.XMLName.Local
							if strings.Contains(valNamevl, "Split") {
								if slices.XMLName.Local == "SplitFromRecord" {
									difTotpax := slices.OriginalNumberOfPax - slices.CurrentNumberOfPax
									pnrTotori += int32(difTotpax)
									fmtSplits := fmt.Sprintf("%v:%v", slices.RecordLocator, difTotpax)
									fncApndix.FncApndixUpdateSlcstr(&pnrObject.Spltfr, fmtSplits)
									if idx == 0 {
										pnrObject.Pnrsrc = slices.RecordLocator
									}
									idx++
								}
								if slices.XMLName.Local == "SplitToRecord" {
									difTotpax := slices.OriginalNumberOfPax - slices.CurrentNumberOfPax
									pnrTotori += int32(difTotpax)
									fmtSplits := fmt.Sprintf("%v:%v", slices.RecordLocator, difTotpax)
									fncApndix.FncApndixUpdateSlcstr(&pnrObject.Spltto, fmtSplits)
								}
							}
						}
						pnrObject.Totori = pnrTotori

						// Get segment
						if len(getRsvpnr.PassengerReservation.Segments.Segment) > 0 {
							slcClssbk, slcFlnbsg, slcRoutsg, slcTimefl := []string{}, []string{}, []string{}, []int64{}
							for _, segmnt := range getRsvpnr.PassengerReservation.Segments.Segment {
								strAirlsg := segmnt.Air.MarketingAirlineCode
								rawFlnbsg := segmnt.Air.MarketingFlightNumber
								intFlnbsg, _ := strconv.Atoi(rawFlnbsg)
								strFlnbsg := strconv.Itoa(intFlnbsg)
								strDepart := segmnt.Air.DepartureAirport
								strArrivl := segmnt.Air.ArrivalAirport
								rawTimefl := segmnt.Air.DepartureDateTime
								fmtTimefl, _ := time.Parse("2006-01-02T15:04:05", rawTimefl)
								intTimefl, _ := strconv.Atoi(fmtTimefl.Format("0601021504"))
								intDatefl, _ := strconv.Atoi(fmtTimefl.Format("060102"))
								slcTimefl = append(slcTimefl, int64(intTimefl))
								slcFlnbsg = append(slcFlnbsg, strAirlsg+"-"+strFlnbsg)
								slcClssbk = append(slcClssbk, segmnt.Air.MarketingClassOfService)
								slcRoutsg = append(slcRoutsg, strDepart+"-"+strArrivl)

								// Get for database flight number list
								if _, istsyc := sycFlnbls.Load(strFlnbsg + strDepart); !istsyc {
									if _, istslc := slices[strFlnbsg+strDepart]; !istslc {
										sycFlnbls.Store(strFlnbsg+strDepart, true)
										nowPrmkey := fmt.Sprintf("%v%v%v%v", strAirlsg, strFlnbsg, strDepart, intDatefl)
										mgoFlnbls = append(mgoFlnbls, mongo.NewUpdateOneModel().
											SetFilter(bson.M{"prmkey": nowPrmkey}).
											SetUpdate(bson.M{"$set": mdlJeddah.MdlJeddahFlnblsDtbase{
												Prmkey: nowPrmkey,
												Datefl: int32(intDatefl),
												Airlfl: strAirlsg,
												Flnbfl: strFlnbsg,
												Depart: strDepart,
											}}).SetUpsert(true))
										fncApndix.FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
											"jeddah_flnbls": &mgoFlnbls,
										}, 200)
									}
								}
							}
							pnrObject.Clssbk = strings.Join(slcClssbk, "|")
							pnrObject.Flnbsg = strings.Join(slcFlnbsg, "|")
							pnrObject.Timefl = slcTimefl[0]
							pnrObject.Timerv = slcTimefl[len(slcTimefl)-1]
							pnrObject.Routsg = strings.Join(slcRoutsg, "|")
						}

						// If cancel data
						if pnrObject.Totcxl > 0 {
							slcClssbk, slcFlnbsg, slcRoutsg, slcTimefl := []string{}, []string{}, []string{}, []int64{}
							for _, itnrxs := range getRsvpnr.History[0].ItineraryHistory {
								strAirlsg := itnrxs.MarketingAirlineCode
								strDepart := itnrxs.DepartureAirport
								strArrivl := itnrxs.ArrivalAirport
								rawFlnbsg := itnrxs.MarketingFlightNumber
								intFlnbsg, _ := strconv.Atoi(rawFlnbsg)
								strFlnbsg := strconv.Itoa(intFlnbsg)
								rawTimefl := itnrxs.DepartureDateTime
								fmtTimefl, _ := time.Parse("2006-01-02T15:04:05", rawTimefl)
								intTimefl, _ := strconv.Atoi(fmtTimefl.Format("0601021504"))
								slcTimefl = append(slcTimefl, int64(intTimefl))
								slcFlnbsg = append(slcFlnbsg, strAirlsg+"-"+strFlnbsg)
								slcClssbk = append(slcClssbk, itnrxs.ClassOfService)
								slcRoutsg = append(slcRoutsg, strDepart+"-"+strArrivl)
							}
							pnrObject.Clssbk = strings.Join(slcClssbk, "|")
							pnrObject.Flnbsg = strings.Join(slcFlnbsg, "|")
							pnrObject.Timefl = slcTimefl[0]
							pnrObject.Timerv = slcTimefl[len(slcTimefl)-1]
							pnrObject.Routsg = strings.Join(slcRoutsg, "|")
						}

					}

					mgoPnrtrc = append(mgoPnrtrc, mongo.NewUpdateOneModel().
						SetFilter(bson.M{"pnrcde": pnrObject.Pnrcde}).
						SetUpdate(bson.M{"$set": pnrObject}).SetUpsert(true))
					fncApndix.FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
						"jeddah_pnrsmr": &mgoPnrtrc,
					}, 200)
				}()
			}
		}

		// Final
		nowswg.Wait()
		fncApndix.FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
			"jeddah_pnrsmr": &mgoPnrtrc,
			"jeddah_flnbls": &mgoFlnbls,
		}, 0)
	}
}
