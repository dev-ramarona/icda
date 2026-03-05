package fncPsglst

import (
	fncApndix "back/apndix/function"
	mdlApndix "back/apndix/model"
	mdlPsglst "back/psglst/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Running process hit passanggerlist daily
func FncPsglstPrcessMainpg(c *gin.Context) {

	// protect single run
	if fncApndix.Status.Sbrapi != 0.0 {
		return
	}
	fncApndix.Status.Sbrapi = 0.01

	// Bind JSON Body input to variable
	inpErrlog := mdlPsglst.MdlPsglstErrlogDtbase{} //save
	if err := c.BindJSON(&inpErrlog); err != nil {
		panic(err)
	}
	errErignr := &inpErrlog.Erignr
	errPrmkey := &inpErrlog.Prmkey

	// Declare date format
	strTimenw := time.Now().Format("0601021504")
	intTimenw, _ := strconv.Atoi(strTimenw)
	intDatenw, _ := strconv.Atoi(strTimenw[0:6])
	strDatepv := time.Now().AddDate(0, 0, -1)
	if inpErrlog.Datefl != 0 {
		tmpDatefl := strconv.Itoa(int(inpErrlog.Datefl))
		strDatepv, _ = time.Parse("060102", tmpDatefl)
	}
	strDatefl := strDatepv.Format("060102")
	intDatefl, _ := strconv.Atoi(strDatefl)

	// Declare airline
	slcAirlfl := []string{inpErrlog.Airlfl}
	if inpErrlog.Airlfl == "" {
		slcAirlfl = []string{"JT", "ID", "IW", "IU", "OD", "SL"}
	}

	// Declare Depart
	slcDepart := []string{inpErrlog.Depart}
	if inpErrlog.Depart == "" {
		slcDepart = fncApndix.FncApndixDstrctGetslc()
	}

	// Declare Flight number
	slcFlnbfl := []string{inpErrlog.Flnbfl}
	if inpErrlog.Flnbfl == "" {
		slcFlnbfl = []string{"All"}
	}

	// Indicator done data
	var totWorker = inpErrlog.Worker
	var mapClslvl = fncApndix.FncApndixClssvlMapobj()
	var mapProvnc = fncApndix.FncApndixProvncMapobj()
	var slcHfbalv = fncApndix.FncApndixHfbalvMapobj()
	var sycFlhour = fncApndix.FncApndixFlhourSycmap()
	var sycFlnbfl = fncApndix.FncApndixFlnbflSycmap()
	var sycFrbase = fncApndix.FncApndixFrbaseSycmap()
	var sycFrtaxs = fncApndix.FncApndixFrtaxsSycmap()
	var sycMilege = fncApndix.FncApndixMilegeSycmap()
	var sycErrlog = FncPsglstErrlogSycmap(int32(intDatefl))
	var sycCurrcv = &sync.Map{}
	var sycChrter = &sync.Map{}
	var sycPnrcde = &sync.Map{}
	var sycPrgrss = &sync.Map{}
	sycPrgrss.Store("nowfln", float64(0))
	sycPrgrss.Store("maxfln", float64(1))
	sycPrgrss.Store("lenair", float64(len(slcAirlfl)))

	// Looping slice airlines
	for _, airlfl := range slcAirlfl {
		fmt.Println("Processing airline:", airlfl, totWorker)

		var idcFlhour = &sync.Map{}
		var idcFrbase = &sync.Map{}
		var idcFrtaxs = &sync.Map{}
		var idcFlnbfl = &sync.Map{}

		// Get Multiple API sessions/tokens
		slcRspssn, err := fncSbrapi.FncSbrapiCrtssnMultpl(airlfl, int(totWorker))
		lgcRspssn := err != nil || slcRspssn[0].Bsttkn == "" || len(slcRspssn) < 1
		FncPsglstErrlogManage(mdlPsglst.MdlPsglstErrlogDtbase{
			Erpart: "sssion", Ersrce: "sbrapi", Erdvsn: "MNFEST",
			Dateup: int32(intDatenw), Timeup: int64(intTimenw),
			Datefl: int32(intDatefl), Airlfl: airlfl, Worker: 5,
		}, lgcRspssn, sycErrlog, errErignr, errPrmkey)
		if lgcRspssn {
			continue
		}

		// Prepare job queue
		jobFllist := make(chan mdlPsglst.MdlPsglstFllistDtbase, 1500)
		var swg sync.WaitGroup

		// Launch 10 workers using 10 tokens
		for i := 0; i < int(totWorker); i++ {
			if len(slcRspssn) >= i+1 {
				if slcRspssn[i].Bsttkn != "" {
					swg.Add(1)
					fmt.Println("Success Token-", i)
					go FncPsglstPrcessWorker(slcRspssn[i],
						&swg,
						jobFllist,
						mapClslvl, mapProvnc, slcHfbalv,
						sycFlhour, sycFrbase, sycFrtaxs, sycErrlog, sycFlnbfl, sycChrter,
						sycCurrcv, sycPnrcde, sycMilege, sycPrgrss,
						idcFlhour, idcFrbase, idcFrtaxs, idcFlnbfl,
						strTimenw, errErignr, errPrmkey)
					continue
				}
				fmt.Println("Failed Token-", i)
			}
		}

		// Looping slice departures
		for _, depart := range slcDepart {

			// Get API Flight List data
			rawFllist, err := fncSbrapi.FncSbrapiFllistMainob(slcRspssn[0],
				mdlSbrapi.MdlSbrapiMsghdrApndix{Datefl: int32(intDatefl),
					Airlfl: airlfl, Depart: depart})
			FncPsglstErrlogManage(mdlPsglst.MdlPsglstErrlogDtbase{
				Erpart: "fllstl", Ersrce: "sbrapi", Erdvsn: "MNFEST",
				Dateup: int32(intDatenw), Timeup: int64(intTimenw),
				Datefl: int32(intDatefl), Airlfl: airlfl, Worker: 5,
				Depart: depart,
			}, err != nil, sycErrlog, errErignr, errPrmkey)
			if err != nil {
				continue
			}

			// Counting all maximal data progress
			if valmax, istmax := sycPrgrss.Load("maxfln"); istmax {
				if valFltmax, mtc := valmax.(float64); mtc {
					sycPrgrss.Store("maxfln", valFltmax+float64(len(rawFllist)))
				}
			}

			// Looping Flight List
			for _, fllist := range rawFllist {

				// Only accept on this route
				if slices.Contains(slcFlnbfl, fllist.Flnbfl) ||
					slices.Contains(slcFlnbfl, "All") {
					jobFllist <- fllist
				}
			}
		}

		// Finish
		close(jobFllist)
		swg.Wait()
		fmt.Printf("Done airline:%s time:%s \n", airlfl,
			time.Now().Format("06-Jan-02/15:04:05"))
		fncSbrapi.FncSbrapiClsssnMultpl(slcRspssn)

		// Reduce total airline progress
		if valair, istair := sycPrgrss.Load("lenair"); istair {
			if valFltair, mtcair := valair.(float64); mtcair {
				sycPrgrss.Store("lenair", valFltair-1)
			}
		}
	}

	// Detect error and count it
	statdt := "Clear"
	sycErrlog.Range(func(key, value any) bool {
		statdt = "Pending"
		return false
	})

	// Final put log action
	rsupdt := fncApndix.FncApndixBulkdbSingle(
		[]mongo.WriteModel{mongo.NewUpdateOneModel().
			SetFilter(bson.M{"datefl": intDatefl}).
			SetUpdate(bson.M{"$set": mdlPsglst.MdlPsglstActlogDtbase{
				Dateup: int32(intDatenw),
				Datefl: int32(intDatefl),
				Timeup: int64(intTimenw),
				Statdt: statdt,
			}}).
			SetUpsert(true)}, "psglst_logact")
	fncApndix.Status.Sbrapi = 0
	if rsupdt != nil {
		panic("Error Insert/Update to DB:" + rsupdt.Error())
	}
	if *errErignr != "" || *errPrmkey != "" {
		c.JSON(500, gin.H{"status": "Failed"})
		return
	}
	c.JSON(200, gin.H{"status": "Done"})
}

// Running process psglst
func FncPsglstPrcessWorker(
	nowObjtkn mdlSbrapi.MdlSbrapiMsghdrParams,
	swg *sync.WaitGroup,
	jobFllist <-chan mdlPsglst.MdlPsglstFllistDtbase,
	mapClslvl map[string]mdlApndix.MdlApndixClsslvDtbase,
	mapProvnc map[string]string,
	slcHfbalv []mdlApndix.MdlApndixHfbalvDtbase,
	sycFlhour, sycFrbase, sycFrtaxs, sycErrlog, sycFlnbfl,
	sycChrter, sycCurrcv, sycPnrcde, sycMilege, sycPrgrss,
	idcFlhour, idcFrbase, idcFrtaxs, idcFlnbfl *sync.Map,
	strTimenw string, errErignr, errPrmkey *string) {

	// Declare global variable
	defer swg.Done()
	var mgoFllist, mgoFlhour []mongo.WriteModel
	var mgoFrbase, mgoFrtaxs []mongo.WriteModel
	var mgoMilege, mgoFlnbfl []mongo.WriteModel
	var mgoPsgsmr, mgoPsgdtl []mongo.WriteModel

	// Get currency
	mapCurrcv := map[string]mdlApndix.MdlApndixCurrcvDtbase{}
	if getCurrcv, ist := sycCurrcv.Load("currcv"); ist {
		if mtcCurrcv, mtc := getCurrcv.(map[string]mdlApndix.MdlApndixCurrcvDtbase); mtc {
			mapCurrcv = mtcCurrcv
		}
	} else {
		var getCurrcv, err = fncSbrapi.FncSbrapiCurrcvMainob(nowObjtkn)
		if err == nil {
			mapCurrcv = getCurrcv
			sycCurrcv.Store("currcv", getCurrcv)
		}
	}

	// iterate jobs
	cntdta := 0
	for fllist := range jobFllist {
		cntdta++

		// prepare locals
		var nowStartm = time.Now()
		var intDatefl = fllist.Datefl
		var dbsFlnbfl, dbsDepart, dbsArrivl = fllist.Flnbfl, fllist.Depart, fllist.Arrivl
		var dbsRoutfl, dbsAirlfl = fllist.Routfl, fllist.Airlfl
		var objParams = mdlSbrapi.MdlSbrapiMsghdrApndix{
			Airlfl: dbsAirlfl, Datefl: intDatefl, Depart: dbsDepart,
			Arrivl: dbsArrivl, Flnbfl: dbsFlnbfl, Routfl: dbsRoutfl}

		// Conver String and int date
		rawNdayfl, _ := time.Parse("060102", strconv.Itoa(int(intDatefl)))
		strNdayfl := rawNdayfl.Format("Mon")
		objParams.Ndayfl = strNdayfl
		strMnthfl := rawNdayfl.Format("0601")
		intMnthfl, _ := strconv.Atoi(strMnthfl)
		objParams.Mnthfl = int32(intMnthfl)
		intDatenw, _ := strconv.Atoi(strTimenw[0:6])
		objParams.Dateup = int32(intDatenw)
		intTimenw, _ := strconv.Atoi(strTimenw)
		objParams.Timeup = int64(intTimenw)

		// Get flight detail
		func() {
			err := fncSbrapi.FncSbrapiFldtilMainob(nowObjtkn, objParams, &fllist)
			if fllist.Flstat == "PDC" {
				FncPsglstErrlogManage(mdlPsglst.MdlPsglstErrlogDtbase{
					Erpart: "fldtil", Ersrce: "sbrapi", Erdvsn: "SLSRPT",
					Dateup: int32(intDatenw), Timeup: int64(intTimenw),
					Datefl: int32(intDatefl), Airlfl: dbsAirlfl,
					Flnbfl: dbsFlnbfl, Routfl: dbsRoutfl, Worker: 1,
				}, err != nil || fllist.Routmx == "", sycErrlog, errErignr, errPrmkey)
			}
		}()

		// Get flight hour
		keyFlhour := dbsAirlfl + dbsFlnbfl + dbsRoutfl
		nulFlhour := true
		if _, ist := idcFlhour.Load(keyFlhour); !ist {
			idcFlhour.Store(keyFlhour, true)
			rspFlhour, err := fncSbrapi.FncSbrapiFlhourMainob(nowObjtkn, sycFlhour, objParams)
			if err == nil && len(rspFlhour) > 0 {

				// Looping all flight hour
				for _, flhour := range rspFlhour {
					if flhour.Flhour == 0 {
						continue
					}

					// Push new data flight to database
					sycFlhour.Store(flhour.Prmkey, flhour)
					mgoFlhour = append(mgoFlhour, mongo.NewUpdateOneModel().
						SetFilter(bson.M{"prmkey": flhour.Prmkey}).
						SetUpdate(bson.M{"$set": flhour}).
						SetUpsert(true))
					nulFlhour = false

					// Push data flight hour if isset
					if flhour.Routfl[:3] == dbsDepart {
						fllist.Flhour = flhour.Flhour
					}
				}
			}
		}

		// Get from syc flight hour if empty
		if nulFlhour {
			if getFlhour, ist := sycFlhour.Load(keyFlhour); ist {
				if mtcFlhour, mtc := getFlhour.(mdlApndix.MdlApndixFlhourDtbase); mtc {
					fllist.Flhour = mtcFlhour.Flhour
					nulFlhour = false
				}
			}
		}

		// If doesn't get flight hour API
		if fllist.Flstat == "PDC" {
			FncPsglstErrlogManage(mdlPsglst.MdlPsglstErrlogDtbase{
				Erpart: "flhour", Ersrce: "sbrapi", Erdvsn: "SLSRPT",
				Dateup: int32(intDatenw), Timeup: int64(intTimenw),
				Datefl: int32(intDatefl), Airlfl: dbsAirlfl,
				Flnbfl: dbsFlnbfl, Routfl: dbsRoutfl, Worker: 1,
			}, nulFlhour, sycErrlog, errErignr, errPrmkey)
		}

		// Get fare component
		func() {

			// Make combination all route
			slcRoutfl := []string{dbsRoutfl}
			slcRoutmx := strings.Split(fllist.Routmx, "-")
			lenRoutmx := len(slcRoutmx)
			for i := 0; i < lenRoutmx-1; i++ {
				for e := i + 1; e < lenRoutmx; e++ {
					mowRoutfl := slcRoutmx[i] + "-" + slcRoutmx[e]
					if !slices.Contains(slcRoutfl, mowRoutfl) {
						slcRoutfl = append(slcRoutfl, mowRoutfl)
					}
				}
			}

			// Looping all route combination
			for _, routfl := range slcRoutfl {
				keyFrball := dbsAirlfl + routfl
				nowObjprm := objParams
				nowObjprm.Depart = routfl[:3]
				nowObjprm.Arrivl = routfl[4:]
				nowObjprm.Routfl = routfl

				// Get farebase
				if _, ist := idcFrbase.Load(keyFrball); !ist {
					nowmgo, err := fncSbrapi.FncSbrapiFrbaseMainob(nowObjtkn, nowObjprm, sycFrbase, mapClslvl)
					if err == nil {
						mgoFrbase = append(mgoFrbase, nowmgo...)
						idcFrbase.Store(keyFrball, true)
					}
				}

				// Get faretaxes
				if _, ist := idcFrtaxs.Load(keyFrball); !ist {

					// Declare looping economy and bisnis
					slcClscbn := []string{"Y"}
					if fllist.Autrzc != 0 {
						slcClscbn = []string{"Y", "C"}
					}
					for _, clscbn := range slcClscbn {
						nowmgo, err := fncSbrapi.FncSbrapiFrtaxsMainob(nowObjtkn, nowObjprm, sycFrtaxs, clscbn)
						if err == nil {
							mgoFrtaxs = append(mgoFrtaxs, nowmgo...)
							idcFrtaxs.Store(keyFrball, true)
						}
					}
				}
			}
		}()

		// Push final flight list
		mgoFllist = append(mgoFllist, mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": fllist.Prmkey}).
			SetUpdate(bson.M{"$set": fllist}).
			SetUpsert(true))
		FncPsglstErrlogManage(mdlPsglst.MdlPsglstErrlogDtbase{
			Erpart: "fllist", Ersrce: "sbrapi", Erdvsn: "MNFEST",
			Dateup: int32(intDatenw), Timeup: int64(intTimenw),
			Datefl: int32(intDatefl), Airlfl: dbsAirlfl, Worker: 1,
			Depart: dbsDepart, Flnbfl: dbsFlnbfl, Routfl: dbsRoutfl, Flstat: fllist.Flstat,
		}, fllist.Flstat != "PDC" && fllist.Flstat != "CANCEL", sycErrlog, errErignr, errPrmkey)

		// Handle PDC flight
		if fllist.Flstat == "PDC" {

			// Push final flightnumber
			var prmkey = dbsAirlfl + dbsFlnbfl
			if _, ist := idcFlnbfl.Load(prmkey); !ist {
				tmpFlnbfl := fncApndix.FncApndixFlnbflPrcess(sycFlnbfl, objParams, prmkey, fllist.Routmx)
				mgoFlnbfl = append(mgoFlnbfl, tmpFlnbfl...)
				idcFlnbfl.Store(prmkey, true)
			}

			// Get passangger list
			rspPsglst, err := fncSbrapi.FncSbrapiPsglstMainob(nowObjtkn, objParams, mapCurrcv, fllist, mapClslvl)
			FncPsglstErrlogManage(mdlPsglst.MdlPsglstErrlogDtbase{
				Erpart: "psglst", Ersrce: "dtbase", Erdvsn: "MNFEST",
				Dateup: int32(intDatenw), Timeup: int64(intTimenw),
				Datefl: int32(intDatefl), Airlfl: dbsAirlfl,
				Flnbfl: dbsFlnbfl, Routfl: dbsRoutfl, Worker: 1,
			}, err != nil, sycErrlog, errErignr, errPrmkey)
			tmpPsgdtl, tmpPsgsmr, tmpFrbase, tmpFrtaxs, tmpFlhour, tmpMilege :=
				FncPsglstPsglstPrcess(rspPsglst, fllist,
					nowObjtkn, objParams,
					sycPnrcde, sycChrter, sycFrbase, sycFrtaxs, sycFlhour, sycMilege,
					idcFrbase, idcFrtaxs, sycErrlog, slcHfbalv,
					mapProvnc, mapCurrcv, mapClslvl, *errErignr, *errPrmkey)
			mgoPsgsmr = append(mgoPsgsmr, tmpPsgsmr...)
			mgoPsgdtl = append(mgoPsgdtl, tmpPsgdtl...)
			mgoMilege = append(mgoMilege, tmpMilege...)
			mgoFlhour = append(mgoFlhour, tmpFlhour...)
			mgoFrbase = append(mgoFrbase, tmpFrbase...)
			mgoFrtaxs = append(mgoFrtaxs, tmpFrtaxs...)
			fncApndix.FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
				"psglst_psgsmr": &mgoPsgsmr,
				"psglst_psgdtl": &mgoPsgdtl,
				"apndix_milege": &mgoMilege,
				"apndix_flhour": &mgoFlhour,
				"apndix_frbase": &mgoFrbase,
				"apndix_frtaxs": &mgoFrtaxs,
			}, 200)
		}

		// Indicator end process
		nowEnddtm := time.Now()
		nowDifftm := nowEnddtm.Sub(nowStartm)
		fmtDifftm := fmt.Sprintf("%02d:%02d:%02d", int(nowDifftm.Hours()),
			int(nowDifftm.Minutes())%60, int(nowDifftm.Seconds())%60)
		fmt.Println("End", fllist.Depart+fllist.Airlfl+fllist.Flnbfl, cntdta, "-",
			dbsAirlfl, dbsFlnbfl, intDatefl, dbsRoutfl, fmtDifftm)

		// Percentage all done data progress
		if valmax, istmax := sycPrgrss.Load("maxfln"); istmax {
			if valFltmax, mtcmax := valmax.(float64); mtcmax {
				if valair, istair := sycPrgrss.Load("lenair"); istair {
					if valFltair, mtcair := valair.(float64); mtcair {
						if valnow, istnow := sycPrgrss.Load("nowfln"); istnow {
							if valFltnow, mtcnow := valnow.(float64); mtcnow {
								donevl := valFltnow + 1
								sycPrgrss.Store("nowfln", donevl)
								fncApndix.Status.Sbrapi = donevl / valFltmax * 100 / valFltair
							}
						}
					}
				}
			}
		}
	}

	// Push if ist data
	fncApndix.FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
		"apndix_fllist": &mgoFllist,
		"apndix_flnbls": &mgoFlnbfl,
		"psglst_psgsmr": &mgoPsgsmr,
		"psglst_psgdtl": &mgoPsgdtl,
		"apndix_milege": &mgoMilege,
		"apndix_flhour": &mgoFlhour,
		"apndix_frbase": &mgoFrbase,
		"apndix_frtaxs": &mgoFrtaxs,
	}, 0)
}
