package fncPsglst

import (
	fncApndix "back/apndix/function"
	mdlApndix "back/apndix/model"
	mdlPsglst "back/psglst/model"
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getcurr() map[string]mdlApndix.MdlApndixCurrcvDtbase {
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("apndix_currcv")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cursor, err := tablex.Find(contxt, bson.M{})
	if err != nil {
		fmt.Println("Err")
	}
	final := map[string]mdlApndix.MdlApndixCurrcvDtbase{}
	for cursor.Next(contxt) {
		semifinal := mdlApndix.MdlApndixCurrcvDtbase{}
		cursor.Decode(&semifinal)
		final[semifinal.Crcode] = semifinal
	}

	return final

}

func getsmr(datefl int32) *sync.Map {
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cursor, err := tablex.Find(contxt, bson.M{"datefl": datefl, "airlfl": "SL"})
	if err != nil {
		fmt.Println("Err")
	}
	final := &sync.Map{}
	for cursor.Next(contxt) {
		semifinal := mdlPsglst.MdlPsglstPsgsmrDtbase{}
		cursor.Decode(&semifinal)
		final.Store(semifinal.Prmkey, semifinal)
	}

	return final

}

func UpdateManualSL(c *gin.Context) {
	mapCurrency := getcurr()
	for _, v := range []int32{260801, 260802, 260803, 260804, 260805, 260806, 260807, 260808, 260809} {
		sycPsgsmr := getsmr(v)
		tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgdtl")
		contxt, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		cursor, err := tablex.Find(contxt, bson.M{"airlfl": "SL", "datefl": v, "isitfl": "F", "isitnr": "", "yrtxvc": 0})
		if err != nil {
			fmt.Println("Err")
		}
		var mgoUpdateDtl []mongo.WriteModel
		var mgoUpdateSmr []mongo.WriteModel
		for cursor.Next(contxt) {
			var slcDtaset mdlPsglst.MdlPsglstPsgdtlDtbase
			cursor.Decode(&slcDtaset)
			// fmt.Println("Rwtxvc: ", slcDtaset.Rwtxvc, "Yrtxvc:", slcDtaset.Yrtxvc, "Yqtxvc:", slcDtaset.Yqtxvc, "Prmkey:", slcDtaset.Prmkey)
			if slcDtaset.Rwtxvc != "" {
				brkdown := strings.Split(slcDtaset.Rwtxvc, "|")
				segment := strings.Split(slcDtaset.Segtkt, "|")

				sumAmountYQ := 0.0
				sumAmountYR := 0.0
				fnlRate := 0.0
				for _, per := range brkdown {
					p := strings.Split(per, ":")
					if len(p) == 3 {
						curr := p[0]
						taxcode := p[1][:2]
						amount, err := strconv.ParseFloat(p[2], 64)
						if err == nil {
							if taxcode == "YQ" || taxcode == "YR" {
								fnlRate := mapCurrency[curr].Crrate
								newamount := amount
								if fnlRate != 0 {
									newamount = amount / fnlRate
								}
								if taxcode == "YQ" {
									sumAmountYQ += newamount
								}
								if taxcode == "YR" {
									sumAmountYR += newamount
								}
							}
						}
					}
				}

				if len(segment) == 1 {
					if slcDtaset.Yrtxvc == 0 || slcDtaset.Yqtxvc == 0 {
						slcDtaset.Taxcrt = fnlRate
						slcDtaset.Srcyqf = "GETTKT"
					}

					// YR
					if slcDtaset.Yrtxvc == 0 {
						slcDtaset.Oldyrtxvc = slcDtaset.Yrtxvc
						slcDtaset.Yrtxvc = sumAmountYR
						prmkey := fmt.Sprintf("%v%v%v%v", slcDtaset.Airlfl, slcDtaset.Flnbfl, slcDtaset.Depart, slcDtaset.Datefl)
						if valsyc, istsyc := sycPsgsmr.Load(prmkey); istsyc {
							if valFltsyc, mtc := valsyc.(mdlPsglst.MdlPsglstPsgsmrDtbase); mtc {
								valFltsyc.Tottyr += sumAmountYR
								sycPsgsmr.Store(prmkey, valFltsyc)
							}
						}
					}

					// YQ
					if slcDtaset.Yqtxvc == 0 {
						slcDtaset.Oldyqtxvc = slcDtaset.Yqtxvc
						slcDtaset.Yqtxvc = sumAmountYQ
						prmkey := fmt.Sprintf("%v%v%v%v", slcDtaset.Airlfl, slcDtaset.Flnbfl, slcDtaset.Depart, slcDtaset.Datefl)
						if valsyc, istsyc := sycPsgsmr.Load(prmkey); istsyc {
							if valFltsyc, mtc := valsyc.(mdlPsglst.MdlPsglstPsgsmrDtbase); mtc {
								valFltsyc.Tottyq += sumAmountYQ
								sycPsgsmr.Store(prmkey, valFltsyc)
							}
						}
					}
				} else {

					// YR
					if slcDtaset.Yrtxvc == 0 && slcDtaset.Yrtxfl != 0 {
						if sumAmountYR != 0 {
							slcDtaset.Oldyrtxvc = slcDtaset.Yrtxvc
							slcDtaset.Yrtxvc = float64(slcDtaset.Yrtxfl)
							prmkey := fmt.Sprintf("%v%v%v%v", slcDtaset.Airlfl, slcDtaset.Flnbfl, slcDtaset.Depart, slcDtaset.Datefl)
							if valsyc, istsyc := sycPsgsmr.Load(prmkey); istsyc {
								if valFltsyc, mtc := valsyc.(mdlPsglst.MdlPsglstPsgsmrDtbase); mtc {
									valFltsyc.Tottyr += slcDtaset.Yrtxvc
									sycPsgsmr.Store(prmkey, valFltsyc)
								}
							}
						}
					}

					// YQ
					if slcDtaset.Yqtxvc == 0 && slcDtaset.Yqtxfl != 0 {
						if sumAmountYQ != 0 {
							slcDtaset.Oldyqtxvc = slcDtaset.Yqtxvc
							slcDtaset.Yqtxvc = float64(slcDtaset.Yqtxfl)
							prmkey := fmt.Sprintf("%v%v%v%v", slcDtaset.Airlfl, slcDtaset.Flnbfl, slcDtaset.Depart, slcDtaset.Datefl)
							if valsyc, istsyc := sycPsgsmr.Load(prmkey); istsyc {
								if valFltsyc, mtc := valsyc.(mdlPsglst.MdlPsglstPsgsmrDtbase); mtc {
									valFltsyc.Tottyq += slcDtaset.Yqtxvc
									sycPsgsmr.Store(prmkey, valFltsyc)
								}
							}
						}
					}
				}
			}

			updPsgdtl := mongo.NewUpdateManyModel().
				SetFilter(bson.M{"prmkey": slcDtaset.Prmkey}).
				SetUpdate(bson.M{"$set": slcDtaset})
			mgoUpdateDtl = append(mgoUpdateDtl, updPsgdtl)
			fncApndix.FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
				"psglst_psgdtl": &mgoUpdateDtl,
			}, 200)
		}
		fncApndix.FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
			"psglst_psgdtl": &mgoUpdateDtl,
		}, 0)

		// Update psgsmr
		sycPsgsmr.Range(func(key, val any) bool {
			if psgsmr, mtc := val.(mdlPsglst.MdlPsglstPsgsmrDtbase); mtc {
				updPsgsmr := mongo.NewUpdateManyModel().
					SetFilter(bson.M{"prmkey": psgsmr.Prmkey}).
					SetUpdate(bson.M{"$set": psgsmr})
				mgoUpdateSmr = append(mgoUpdateSmr, updPsgsmr)
			}
			return true
		})
		fncApndix.FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
			"psglst_psgsmr": &mgoUpdateSmr,
		}, 0)

		fmt.Println("done", v)
	}
}
