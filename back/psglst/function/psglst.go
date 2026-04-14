package fncPsglst

import (
	fncApndix "back/apndix/function"
	mdlApndix "back/apndix/model"
	mdlPsglst "back/psglst/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FncPsglstPsglstPrcess(rspPsglst []mdlPsglst.MdlPsglstPsgdtlDtbase, fllist mdlApndix.MdlApndixFllistDtbase,
	nowObjtkn mdlSbrapi.MdlSbrapiMsghdrParams, objParams mdlSbrapi.MdlSbrapiMsghdrApndix,
	sycPnrcde, sycChrter, sycFrbase, sycFrtaxs, sycFlhour, sycMilege,
	idcFrbase, idcFrtaxs, sycErrlog, sycProvnc *sync.Map,
	slcHfbalv []mdlApndix.MdlApndixHfbalvDtbase,
	mapCurrcv map[string]mdlApndix.MdlApndixCurrcvDtbase,
	mapClslvl map[string]mdlApndix.MdlApndixClsslvDtbase, errErignr, errPrmkey *string) (
	[]mongo.WriteModel, []mongo.WriteModel, []mongo.WriteModel,
	[]mongo.WriteModel, []mongo.WriteModel, []mongo.WriteModel, []mongo.WriteModel) {
	sycWgroup, sycClrpsg, sycNulpsg := &sync.WaitGroup{}, &sync.Map{}, &sync.Map{}
	totPsgdtl := len(rspPsglst)
	for _, psglst := range rspPsglst {

		// Get null data
		if psglst.Tktnvc == "" || psglst.Pnrcde == "" {
			err := fncSbrapi.FncSbrapiPsgdtaMainob(nowObjtkn, mapClslvl, &psglst)
			if err != nil || psglst.Tktnvc == "" || psglst.Pnrcde == "" {
				fncApndix.FncApndixUpdateSlcstr(&psglst.Noterr, "PSGDATA NIL")
			} else {
				fncApndix.FncApndixUpdateSlcstr(&psglst.Noteup, "PSGDATA GET")
			}
		}

		// Running concurency every psglst
		sycWgroup.Add(1)
		go FncPslgstRsvpnrMainpg(psglst,
			sycClrpsg, sycNulpsg, sycPnrcde, sycChrter, sycMilege,
			mapCurrcv,
			sycWgroup, nowObjtkn,
			psglst.Airlfl, psglst.Pnrcde, "")
	}

	// Looping null and wait all goroutine finish
	fnlPrvpsg := map[string]map[string]mdlPsglst.MdlPsglstPsgdtlDtbase{}
	sycWgroup.Wait()
	sycNulpsg.Range(func(key, val any) bool {
		if mtcPsglst, mtc := val.(mdlPsglst.MdlPsglstPsgdtlDtbase); mtc {
			if _, ist := sycClrpsg.Load(key); !ist {
				slcPntitl := strings.Split(mtcPsglst.Pnritl, "|")
				for _, pntitl := range slcPntitl {
					if pntitl == "" {
						continue
					}
					arrPnritl := strings.Split(pntitl, "*")
					nowAirlfl := arrPnritl[0]
					nowPnrcde := arrPnritl[1]
					if fnlPrvpsg[nowAirlfl] == nil {
						fnlPrvpsg[nowAirlfl] = make(map[string]mdlPsglst.MdlPsglstPsgdtlDtbase)
					}
					if slices.Contains([]string{"JT", "ID", "IW", "IU", "OD", "SL"}, nowAirlfl) {
						fnlPrvpsg[nowAirlfl][nowPnrcde+"|"+mtcPsglst.Prmkey] = mtcPsglst
					} else {
						sycClrpsg.Store(mtcPsglst.Prmkey, mtcPsglst)
					}
				}
			}
		}
		return true
	})

	// Looping exist data
	for airlfl, fstPrvpsg := range fnlPrvpsg {
		newObjtkn, er1 := fncSbrapi.FncSbrapiCrtssnMainob(airlfl)
		tmpSycwgp := &sync.WaitGroup{}
		if er1 != nil {
			for i := 0; i < 3; i++ {
				getObjtkn, er3 := fncSbrapi.FncSbrapiCrtssnMainob(airlfl)
				if er3 == nil && getObjtkn.Bsttkn != "" {
					newObjtkn = getObjtkn
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
		for rawkey, scdPrvpsg := range fstPrvpsg {
			slckey := strings.Split(rawkey, "|")
			pnrcde, prmkey := slckey[0], slckey[1]
			if _, ist := sycClrpsg.Load(prmkey); !ist {
				tmpSycwgp.Add(1)
				go FncPslgstRsvpnrMainpg(scdPrvpsg,
					sycClrpsg, sycNulpsg, sycPnrcde, sycChrter, sycMilege,
					mapCurrcv,
					tmpSycwgp, newObjtkn,
					airlfl, pnrcde, "last")
			}
		}
		tmpSycwgp.Wait()
		fncSbrapi.FncSbrapiClsssnMainob(newObjtkn)
	}

	// Declare empty variable final
	totClrpsg := 0
	mgoFrbase, mgoFrtaxs := []mongo.WriteModel{}, []mongo.WriteModel{}
	mgoFlhour, mgoMilege := []mongo.WriteModel{}, []mongo.WriteModel{}
	mgoPsgdtl, mgoPsgsmr := []mongo.WriteModel{}, []mongo.WriteModel{}
	mgoProvnc := []mongo.WriteModel{}
	fnlPsglst := []mdlPsglst.MdlPsglstPsgdtlDtbase{}
	mapPaidbt := map[string]int{}
	mapQntybt := map[string]int{}
	mapWghtbt := map[string]int{}
	mapFbavbt := map[string]int{}
	totSmmary := mdlPsglst.MdlPsglstPsgsmrDtbase{
		Prmkey: objParams.Airlfl + objParams.Flnbfl + objParams.Depart +
			strconv.Itoa(int(objParams.Datefl)),
		Airlfl: objParams.Airlfl,
		Depart: objParams.Depart,
		Flnbfl: objParams.Flnbfl,
		Ndayfl: objParams.Ndayfl,
		Datefl: objParams.Datefl,
		Mnthfl: objParams.Mnthfl,
		Flstat: fllist.Flstat,
		Seatcn: fllist.Seatcn,
		Airtyp: fllist.Airtyp,
		Flhour: fllist.Flhour,
	}

	// Semi final loop and push to final
	sycClrpsg.Range(func(key, val any) bool {
		if mtcPsglst, mtc := val.(mdlPsglst.MdlPsglstPsgdtlDtbase); mtc {

			// Get segment now
			if mtcPsglst.Segtkt != "" {
				prvTimefl, prvRoutfl, slcSegtkt, mtcSegtkt := "", "", []string{}, false
				sptSegtkt := strings.Split(mtcPsglst.Segtkt, "|")
				fstDepart := strings.Split(sptSegtkt[0], "-")[1]
				lstArrivl := strings.Split(sptSegtkt[len(sptSegtkt)-1], "-")[2]
				istRoutpp := fstDepart == lstArrivl
				for _, segtkt := range sptSegtkt {
					cpntkt := strings.Split(segtkt, "-")
					timecp := strings.Split(cpntkt[0], ":")
					if intime, _ := strconv.Atoi(timecp[0]); int64(intime) == mtcPsglst.Timefl ||
						(mtcPsglst.Routvc == cpntkt[1]+"-"+cpntkt[2] && mtcPsglst.Flnbfl == cpntkt[5]) {
						mtcSegtkt = true
					}

					// Gate logic
					if prvTimefl == "" {
						slcSegtkt = append(slcSegtkt, segtkt)
					} else {
						fmtprv, _ := time.Parse("0601021504", prvTimefl)
						fmtnow, _ := time.Parse("0601021504", timecp[0])
						fmtdif := fmtprv.Sub(fmtnow)
						if fmtdif.Hours() > 24 || (prvRoutfl == cpntkt[1]+"-"+cpntkt[2] && istRoutpp) {
							if mtcSegtkt {
								break
							}
							slcSegtkt = []string{}
						} else {
							slcSegtkt = append(slcSegtkt, segtkt)
						}

					}

					// Prev time flight or arrival
					prvRoutfl = cpntkt[2] + "-" + cpntkt[1]
					prvTimefl = timecp[0]
					if timecp[1] != "0101010000" {
						prvTimefl = timecp[1]
					}
				}

				// Get highest fba
				strSegtkt := strings.Join(slcSegtkt, "|")
				slcMaxfba := []int{int(mtcPsglst.Fbavbt)}
				bolStpsrc := true
				for _, nowSegtkt := range slcSegtkt {
					for _, hfbalv := range slcHfbalv {

						// Regex Airline
						nowAirlfl := hfbalv.Airlfl
						regAirlfl := regexp.MustCompile("-(" + nowAirlfl + ")-")
						lgcAirlfl := nowAirlfl == "ALL" || regAirlfl.MatchString(nowSegtkt)

						// Regex class flown
						nowClssfl := hfbalv.Clssfl
						regClssfl := regexp.MustCompile(`-(` + nowClssfl + `)$`)
						lgcClssfl := nowClssfl == "ALL" || regClssfl.MatchString(nowSegtkt)

						// Regex route flown
						strRoutrg := func(routfl string) string {
							fnlRoutfl, slcRoutfl := "", strings.Split(routfl, "-")
							for _, dstrct := range slcRoutfl {
								if dstrct == "ALL" {
									fnlRoutfl += "-[A-Z]{3}"
									continue
								}
								fnlRoutfl += "-(" + dstrct + ")"
							}
							if !strings.Contains(routfl, "ALL") {
								fnlRoutfl += "-|-(" + slcRoutfl[0] + ")-[A-Z]{3}-.+"
								fnlRoutfl += "-[A-Z]{3}-(" + slcRoutfl[1] + ")"
							}
							return fnlRoutfl + "-"
						}(hfbalv.Routfl)
						regRoutfl := regexp.MustCompile(strRoutrg)
						lgcRoutfl := regRoutfl.MatchString(strSegtkt)

						// Final result
						if lgcAirlfl && lgcClssfl && lgcRoutfl && bolStpsrc {
							slcMaxfba = append(slcMaxfba, int(hfbalv.Hfbabt))
							if hfbalv.Source == "VCR" {
								bolStpsrc = false
								slcMaxfba = []int{int(mtcPsglst.Fbavbt)}
							}
							break
						}
					}
				}
				if mtcPsglst.Hfbabt == 0 && len(slcMaxfba) > 0 {
					mtcPsglst.Hfbabt = int32(slices.Max(slcMaxfba))
				}
			}
			fnlPsglst = append(fnlPsglst, mtcPsglst)

			// Total group summary bg and ae
			if mtcPsglst.Groupc != "-" && mtcPsglst.Groupc != "" {
				mapPaidbt[mtcPsglst.Groupc] += int(mtcPsglst.Paidbt)
				mapQntybt[mtcPsglst.Groupc] += int(mtcPsglst.Qntybt)
				mapWghtbt[mtcPsglst.Groupc] += int(mtcPsglst.Wghtbt)
				mapFbavbt[mtcPsglst.Groupc] += int(mtcPsglst.Hfbabt)
			}
		}
		return true
	})

	// Looping again final
	for _, psglst := range fnlPsglst {
		if val, ist := mapPaidbt[psglst.Groupc]; ist {
			psglst.Ptotbt = int32(val)
		} else {
			psglst.Ptotbt = psglst.Paidbt
		}
		if val, ist := mapQntybt[psglst.Groupc]; ist {
			psglst.Qtotbt = int32(val)
		} else {
			psglst.Qtotbt = psglst.Qntybt
		}
		if val, ist := mapWghtbt[psglst.Groupc]; ist {
			psglst.Wtotbt = int32(val)
		} else {
			psglst.Wtotbt = psglst.Wghtbt
		}
		if val, ist := mapFbavbt[psglst.Groupc]; ist {
			psglst.Ftotbt = int32(val)
		} else {
			psglst.Ftotbt = psglst.Hfbabt

		}

		// Manage route
		fnlRoutac := psglst.Routfl
		for _, routll := range []string{psglst.Routvc, psglst.Routfr} {
			if len(routll) >= 7 {
				slcRoutmx := strings.Split(psglst.Routmx, "-")
				slcRoutac := []string{}
				varDeprfl, varDeprvc := psglst.Depart, ""
				varArvlfl, varArvlvc := psglst.Arrivl, ""
				if psglst.Isitfl == "F" && len(routll) >= 7 {
					varDeprvc = routll[:3]
					varArvlvc = routll[4:]
				}
				for _, dstrct := range slcRoutmx {
					if varDeprfl == dstrct || varDeprvc == dstrct || len(slcRoutac) != 0 {
						slcRoutac = append(slcRoutac, dstrct)
					}
					if dstrct == varArvlfl || dstrct == varArvlvc {
						break
					}
				}
				if len(fnlRoutac) >= len(strings.Join(slcRoutac, "-")) {
					continue
				}
				fnlRoutac = strings.Join(slcRoutac, "-")
				psglst.Routac = strings.Join(slcRoutac, "-")
			}

			// Get route actual
			slcRoutac := strings.Split(fnlRoutac, "-")
			if slcPstion := slices.Index(slcRoutac, psglst.Depart); slcPstion != -1 {
				if slcPstion+1 < len(slcRoutac) {
					psglst.Routfl = strings.Join(slcRoutac[slcPstion:slcPstion+2], "-")
				}
			}
			if psglst.Routfl == "" {
				psglst.Routfl = psglst.Depart + "-" + psglst.Arrivl
			}
		}

		// Get rout from fare calc and segment
		slcRoutfr, cekRoutfr, lstRoutsg := []string{}, false, ""
		if len(psglst.Routfr) >= 7 {
			regRoutfr := regexp.MustCompile(fmt.Sprintf("%s.+%s", psglst.Routfr[:3],
				psglst.Routfr[4:]))
			if res := regRoutfr.MatchString(psglst.Routsg); res {

				// Combine route
				slcRoutmx := strings.Split(psglst.Routmx, "-")
				nowRoutmx := slcRoutmx[0] + "-" + slcRoutmx[len(slcRoutmx)-1]
				segFullrt := psglst.Routsg
				if strings.Contains(psglst.Routsg, nowRoutmx) {
					segFullrt = strings.Replace(psglst.Routsg, nowRoutmx, psglst.Routmx, 1)
				}

				// Looping full rout max
				for _, routsg := range strings.Split(segFullrt, "-") {
					if strings.Contains(psglst.Routfx, lstRoutsg+"-"+routsg) {
						if len(slcRoutfr) > 0 && slcRoutfr[len(slcRoutfr)-1] == lstRoutsg {
							slcRoutfr = slcRoutfr[:len(slcRoutfr)-1]
							continue
						}
					}
					if lstRoutsg+"-"+routsg == psglst.Routfr {
						slcRoutfr = []string{lstRoutsg, routsg}
						break
					}
					if psglst.Routfr[4:] == routsg && len(slcRoutfr) > 0 {
						cekRoutfr = true
						slcRoutfr = append(slcRoutfr, routsg)
						break
					}
					lstRoutsg = routsg
					if psglst.Routfr[:3] == routsg || len(slcRoutfr) > 0 {
						slcRoutfr = append(slcRoutfr, routsg)
					}
				}

				// Last push route actual from route fare
				if cekRoutfr && len(slcRoutfr) >= 2 {
					strRoutfr := strings.Join(slcRoutfr, "-")
					if len(strRoutfr) > len(fnlRoutac) {
						psglst.Routac = strRoutfr
						fnlRoutac = strRoutfr
					}
				}
			}
		}

		// Get farebase and faretaxes if flow
		if psglst.Isitfl == "F" {

			// Looping to get farebase vcr and flown
			mapRoutfb := map[string]string{"routfl": psglst.Airlfl + psglst.Routfl}
			frbRoutvc := psglst.Routvc
			if psglst.Ntafvc == 0 && psglst.Isitnr == "" && psglst.Frbcde != "HB" {
				if psglst.Routvc == "" {
					frbRoutvc = psglst.Depart + "-" + psglst.Arrivl
				}
				mapRoutfb["routvc"] = psglst.Airlfl + frbRoutvc
				if slices.Contains([]string{"JT", "ID", "IW", "IU", "OD", "SL"}, psglst.Airlvc) {
					mapRoutfb["routvc"] = psglst.Airlvc + frbRoutvc
				}
			}
			for keyfst, valfst := range mapRoutfb {
				if len(valfst) != 9 {
					continue
				}

				// Get flight hour from API
				if _, istfst := idcFrbase.Load(valfst); !istfst {
					objParams := mdlSbrapi.MdlSbrapiMsghdrApndix{
						Airlfl: valfst[:2], Depart: valfst[2:5], Arrivl: valfst[6:], Routfl: valfst[2:]}
					nowmgo, err := fncSbrapi.FncSbrapiFrbaseMainob(nowObjtkn, objParams, sycFrbase, mapClslvl)
					if err == nil {
						mgoFrbase = append(mgoFrbase, nowmgo...)
						idcFrbase.Store(valfst, true)
					}
				}

				// Looping
				for keyscd, valscd := range func() []string {
					slcKeytax := []string{valfst + psglst.Frbcde}
					strKeycls := valfst + psglst.Clssfl
					if strings.Contains(psglst.Frbcde, "RT") {
						strKeycls = valfst + psglst.Clssfl + "RT"
					}
					slcKeytax = append(slcKeytax, strKeycls)
					return slcKeytax
				}() {

					// Get farebase from sync
					istFrbase, ist := sycFrbase.Load(valscd)
					if mtcFrbase, mtc := istFrbase.(mdlApndix.MdlApndixFrbaseDtbase); ist && mtc {
						if keyfst == "routfl" {
							psglst.Ntaffl = mtcFrbase.Frbnta
						} else {
							psglst.Ntafvc = float64(mtcFrbase.Frbnta)
							psglst.Srcfrb = "CLSSFL"
							if keyscd == 0 {
								psglst.Srcfrb = "FRBCDE"
							}
						}
						break
					}
				}
			}

			// Looping to get faretaxes vcr and flown
			mapRoutax := map[string]string{"routfl": psglst.Airlfl + psglst.Routfl}
			if psglst.Yqtxvc == 0 && psglst.Isitnr == "" && psglst.Frbcde != "HB" {
				taxRoutvc := psglst.Routvc
				if psglst.Routvc == "" {
					taxRoutvc = psglst.Depart + "-" + psglst.Arrivl
				}
				if psglst.Ntafvc != 0 && psglst.Routfr != "" {
					taxRoutvc = psglst.Routfr
				}
				taxArilvc := psglst.Airlfl
				if slices.Contains([]string{"JT", "ID", "IW", "IU", "OD", "SL"}, psglst.Airlvc) {
					taxArilvc = psglst.Airlvc
				}
				mapRoutax["routvc"] = taxArilvc + taxRoutvc
			}
			for keyfst, valfst := range mapRoutax {
				if len(valfst) != 9 {
					continue
				}

				// Get primary key and time book or issued
				nowKeytax := valfst + psglst.Cbinvc
				nowClscbn := psglst.Cbinvc
				if keyfst == "routfl" || nowClscbn == "" {
					nowKeytax = valfst + psglst.Cbinfl
					nowClscbn = psglst.Cbinfl
				}
				nowDatemc := strconv.Itoa(int(psglst.Datefl))
				if len(strconv.Itoa(int(psglst.Timeis))) == 10 {
					nowDatemc = strconv.Itoa(int(psglst.Timeis))[:6]
				}
				if len(strconv.Itoa(int(psglst.Timecr))) == 10 {
					nowDatemc = strconv.Itoa(int(psglst.Timecr))[:6]
				}
				intDatemc, _ := strconv.Atoi(nowDatemc)

				// Looping cabin
				slcKeytax := []string{"Y"}
				if psglst.Bookdc != 0 {
					slcKeytax = append(slcKeytax, "C")
				}

				// Get flight hour from API
				if _, ist := idcFrtaxs.Load(valfst); !ist {
					for _, valscd := range slcKeytax {
						objParams := mdlSbrapi.MdlSbrapiMsghdrApndix{
							Airlfl: valfst[:2], Depart: valfst[2:5], Arrivl: valfst[6:], Routfl: valfst[2:]}
						nowmgo, err := fncSbrapi.FncSbrapiFrtaxsMainob(nowObjtkn, objParams, sycFrtaxs, valscd)
						if err == nil {
							mgoFrtaxs = append(mgoFrtaxs, nowmgo...)
							idcFrtaxs.Store(valfst, true)
						}
					}
				}

				// Get farebase from sync
				istFrtaxs, ist := sycFrtaxs.Load(nowKeytax)
				if mtcFrtaxs, mtc := istFrtaxs.(mdlApndix.MdlApndixFrtaxsDtbase); ist && mtc {
					if keyfst == "routfl" {
						psglst.Yqtxfl = mtcFrtaxs.Ftfuel
					} else {
						psglst.Yqtxvc = float64(mtcFrtaxs.Ftfuel)
						slcHstory := strings.Split(mtcFrtaxs.Hstory, "|")
						if mtcFrtaxs.Datend <= int32(intDatemc) {
							continue
						} else if len(slcHstory) > 0 && mtcFrtaxs.Hstory != "" {
							lenHstory := len(slcHstory) - 1
							for idxtrd, valtrd := range slcHstory {
								slcValtrd := strings.Split(valtrd, "@")
								intDatend, _ := strconv.Atoi(slcValtrd[0])
								if intDatend <= intDatemc || lenHstory == idxtrd {
									slcFrtaxs := strings.Split(slcValtrd[1], "/")
									for _, valfrt := range slcFrtaxs {
										slcValfrt := strings.Split(valfrt, ":")
										strTaxcde := slcValfrt[0]
										intFrtaxs, _ := strconv.Atoi(slcValfrt[1])
										if strTaxcde == "yq" && intFrtaxs != 0 {
											psglst.Yqtxvc = float64(intFrtaxs)
										}
									}
									break
								}
							}
						}
					}
				}
			}
		}

		// Get final price
		if fnlRoutac == "" {
			fnlRoutac = psglst.Depart + "-" + psglst.Arrivl
		}
		slcRoutac, totMilege, nowMilege := strings.Split(fnlRoutac, "-"), float64(0), float64(0)
		for i := 0; i < len(slcRoutac)-1; i++ {

			// Reuse func milege
			fncMilege := func(nowRoutac string) bool {
				istMilege, ist := sycMilege.Load(nowRoutac)
				if mtcMilege, mtc := istMilege.(mdlApndix.MdlApndixMilegeDtbase); mtc && ist {
					if slcRoutac[i] == psglst.Depart {
						nowMilege = float64(mtcMilege.Milege)
					}
					totMilege += float64(mtcMilege.Milege)
				}
				return ist
			}

			// Route milege hit API Sabre if null data
			nowRoutac := slcRoutac[i] + "-" + slcRoutac[i+1]
			if rspfnc := fncMilege(nowRoutac); !rspfnc {
				rspMilege, err := fncSbrapi.FncSbrapiMilegeMainob(nowObjtkn, fnlRoutac)
				if err == nil {
					for _, milege := range rspMilege {
						if _, ist := sycMilege.Load(milege.Routfl); !ist {
							sycMilege.Store(milege.Routfl, milege)
							mgoMilege = append(mgoMilege, mongo.NewUpdateOneModel().
								SetFilter(bson.M{"routfl": milege.Routfl}).
								SetUpdate(bson.M{"$set": milege}).SetUpsert(true))
						}
					}
					fncMilege(nowRoutac)
				}
			}
		}

		// Final rate and price adjustment
		psglst.Frrate = nowMilege / totMilege
		valChrter := float64(1)
		if psglst.Isitct == "CT" {
			valChrter = 0
		}
		psglst.Ntaffl = int32(float64(psglst.Ntaffl) * valChrter)
		psglst.Yqtxfl = int32(float64(psglst.Yqtxfl) * valChrter)
		psglst.Ntafvc = psglst.Ntafvc * psglst.Frrate * valChrter
		psglst.Yqtxvc = psglst.Yqtxvc * psglst.Frrate * valChrter
		psglst.Fareae = psglst.Fareae * psglst.Frrate * valChrter

		// Push summary
		if psglst.Isitfl == "F" {
			totSmmary.Totnta += psglst.Ntafvc
			totSmmary.Tottyq += psglst.Yqtxvc
			totSmmary.Totpax += 1
			totSmmary.Totfae += psglst.Fareae
			totSmmary.Totqfr += psglst.Qsrcvc
			totSmmary.Totrph += (psglst.Ntafvc + psglst.Yqtxvc) / fllist.Flhour
		}

		// Get province
		cekProvnc := true
		totSmmary.Routfl = psglst.Routfl
		getProvnc, istProvnc := sycProvnc.Load(psglst.Routfl)
		if strProvnc, mtcProvnc := getProvnc.(string); istProvnc && mtcProvnc {
			psglst.Provnc = strProvnc
			totSmmary.Provnc = strProvnc
			if strProvnc != "" {
				cekProvnc = false
			}
		}

		// Push to error log
		if psglst.Isitfl == "F" {
			FncPsglstErrlogManage(mdlPsglst.MdlPsglstErrlogDtbase{
				Erpart: "provnc", Ersrce: "dtbase", Erdvsn: "SLSRPT",
				Dateup: int32(objParams.Dateup), Timeup: int64(objParams.Timeup),
				Datefl: int32(objParams.Datefl), Airlfl: objParams.Airlfl,
				Flnbfl: objParams.Flnbfl, Routfl: objParams.Routfl, Worker: 1,
			}, cekProvnc, sycErrlog, errErignr, errPrmkey)

			// Push to database provnc blank
			if !istProvnc {
				varProvnc := mdlApndix.MdlApndixProvncDtbase{
					Routfl: objParams.Routfl, Provnc: "", Updtby: ""}
				sycProvnc.Store(objParams.Routfl, varProvnc)
				mgoProvnc = append(mgoProvnc, mongo.NewUpdateOneModel().
					SetFilter(bson.M{"routfl": objParams.Routfl}).
					SetUpdate(bson.M{"$set": varProvnc}).SetUpsert(true))
			}
		}

		// Push final to database
		mgoPsgdtl = append(mgoPsgdtl, mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": psglst.Prmkey}).
			SetUpdate(bson.M{"$set": psglst}).
			SetUpsert(true))
		totClrpsg++
	}

	// Push to error log
	FncPsglstErrlogManage(mdlPsglst.MdlPsglstErrlogDtbase{
		Erpart: "psgdtl", Ersrce: "sbrapi", Erdvsn: "MNFEST",
		Dateup: int32(objParams.Dateup), Timeup: int64(objParams.Timeup),
		Datefl: int32(objParams.Datefl), Airlfl: objParams.Airlfl,
		Flnbfl: objParams.Flnbfl, Routfl: objParams.Routfl, Worker: 1,
		Paxdif: fmt.Sprintf("%d/%d", totClrpsg, totPsgdtl),
	}, totClrpsg != totPsgdtl, sycErrlog, errErignr, errPrmkey)

	// Return final data
	mgoPsgsmr = append(mgoPsgsmr, mongo.NewUpdateOneModel().
		SetFilter(bson.M{"prmkey": totSmmary.Prmkey}).
		SetUpdate(bson.M{"$set": totSmmary}).
		SetUpsert(true))
	return mgoPsgdtl, mgoPsgsmr, mgoFrbase, mgoFrtaxs, mgoFlhour, mgoMilege, mgoProvnc
}
