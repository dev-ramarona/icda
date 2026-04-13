package fncSbrapi

import (
	fncApndix "back/apndix/function"
	mdlApndix "back/apndix/model"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Comamand macro Sabre API Sreen
func FncSbrapiFrtaxsMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix, sycFrtaxs *sync.Map,
	clscbn string,
) ([]mongo.WriteModel, error) {

	// Isi struktur data
	tmpFrbase := 100000000
	rawDatefl := time.Now()
	dmyDatefl := rawDatefl.Format("2006-01-02T00:00:00")
	fnlFrtaxs := []mongo.WriteModel{}
	bdyFrtaxs := mdlSbrapi.MdlSbrapiFrtaxsReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiFrtaxsReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncApndix.Pcckey,
				"Calculate Air Tax For Itinerary", "OTA_AirTaxRQ", unqhdr),
			Security: mdlSbrapi.MdlSbrapiFrtaxsReqscr{
				BinarySecurityToken: mdlSbrapi.MdlSbrapiFrtaxsReqbst{
					ValueType: "String", EncodingType: "wsse:Base64Binary", Token: unqhdr.Bsttkn,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiFrtaxsReqbdy{
			AirTaxRQ: mdlSbrapi.MdlSbrapiFrtaxsReqatx{
				Xmlns:   "http://webservices.sabre.com/sabreXML/2003/07",
				Version: "2.0.3",
				ItineraryInfos: mdlSbrapi.MdlSbrapiFrtaxsReqits{
					RPH:                "2",
					SalePseudoCityCode: "JKT",
					TicketingCarrier:   apndix.Airlfl,
					ValidatingCarrier:  apndix.Airlfl,
					FlightSegment: mdlSbrapi.MdlSbrapiFrtaxsReqfls{
						ArrivalDateTime:   dmyDatefl,
						DepartureDateTime: dmyDatefl,
						FlightNumber:      "0",
						ResBookDesigCode:  clscbn,
						DepartureAirport: mdlSbrapi.MdlSbrapiFrtaxsReqdpa{
							CodeContext:  "IATA",
							LocationCode: apndix.Depart},
						ArrivalAirport: mdlSbrapi.MdlSbrapiFrtaxsReqarv{
							CodeContext:  "IATA",
							LocationCode: apndix.Arrivl},
						Equipment:        mdlSbrapi.MdlSbrapiFrtaxsReqeqp{AirEquipType: ""},
						MarketingAirline: mdlSbrapi.MdlSbrapiFrtaxsReqcde{Code: apndix.Airlfl},
						OperatingAirline: mdlSbrapi.MdlSbrapiFrtaxsReqcde{Code: apndix.Airlfl},
					},
					AirFareInfo: mdlSbrapi.MdlSbrapiFrtaxsReqafi{
						PTC_FareBreakdown: mdlSbrapi.MdlSbrapiFrtaxsReqfbk{
							PassengerType: mdlSbrapi.MdlSbrapiFrtaxsReqcde{
								Code: "ADT"},
							PassengerFare: mdlSbrapi.MdlSbrapiFrtaxsReqpsf{
								BaseFare: mdlSbrapi.MdlSbrapiFrtaxsReqbsf{
									Amount: strconv.Itoa(tmpFrbase)}},
						},
					},
				},
			},
		},
	}

	// Treatment APO Session
	raw, err := FncSbrapiMsghdrXmldta(bdyFrtaxs)
	if err != nil {
		return fnlFrtaxs, err
	}

	// Parsing XML ke dalam struktur Go
	rspFrtaxs := mdlSbrapi.MdlSbrapiFrtaxsRspenv{}
	err = xml.Unmarshal([]byte(raw), &rspFrtaxs)
	if err != nil {
		return fnlFrtaxs, err
	}

	// Return String
	rawFrtaxs := rspFrtaxs.Body.AirTaxRS.ItineraryInfos.ItineraryInfo.TaxInfo
	fnlFrtaxs = FncSbrapiFrtaxsTrtmnt(rawFrtaxs, apndix, sycFrtaxs, clscbn, tmpFrbase)
	return fnlFrtaxs, nil
}

// Treatment data raw flight list
func FncSbrapiFrtaxsTrtmnt(rawxml mdlSbrapi.MdlSbrapiFrtaxsRsptxi,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix, sycFrtaxs *sync.Map,
	clscbn string, tmpFrbase int,
) []mongo.WriteModel {

	// Declare first output
	var taxPrmkey = apndix.Airlfl + apndix.Routfl + clscbn
	var tmpOthers = []string{}
	var tmpFrtxid = []int{}
	var slcFtothr = []bson.D{}
	var fnlFrtaxs = mdlApndix.MdlApndixFrtaxsDtbase{
		Prmkey: taxPrmkey,
		Airlfl: apndix.Airlfl,
		Cbinfl: clscbn,
		Depart: apndix.Depart,
		Routfl: apndix.Routfl,
		Ftothr: "",
	}

	// Return non error data
	for _, frtaxs := range rawxml.TaxDetails.Tax {
		valFrtaxs, _ := strconv.Atoi(frtaxs.Amount)
		switch frtaxs.TaxCode {
		case "YQF":
			fnlFrtaxs.Ftfuel = int32(valFrtaxs)
		case "D5":
			fnlFrtaxs.Ftaptx = int32(valFrtaxs)
		case "P4":
			fnlFrtaxs.Ftiwjr = int32(valFrtaxs)
		case "YRI":
			fnlFrtaxs.Ftaxyr = int32(valFrtaxs)
		default:
			if frtaxs.Type == "F" {
				if frtaxs.Amount != "0" {
					txtFrtaxs := frtaxs.TaxCode + ":" + frtaxs.Amount
					tmpOthers = append(tmpOthers, txtFrtaxs)
				}
			} else {
				if frtaxs.TaxCode == "ID" {
					tmpFrtxid = append(tmpFrtxid, valFrtaxs)
				} else {
					slcFtothr = append(slcFtothr, bson.D{
						{Key: frtaxs.TaxCode, Value: valFrtaxs}})
				}
			}
		}
	}

	// Last treatment taxes ID
	var now = fnlFrtaxs
	if len(tmpFrtxid) > 0 {
		taxRateid := math.Round((float64(tmpFrtxid[0]) / float64(now.Ftfuel+now.Ftaxyr+int32(tmpFrbase))) * 100)
		strRateid := strconv.Itoa(int(taxRateid))
		fltRateid, _ := strconv.ParseFloat(strRateid, 32)
		fnlFrtaxs.Ftppnx = fltRateid / 100
		now.Ftppnx = fnlFrtaxs.Ftppnx
	}

	// Treatment frtax other
	if len(tmpOthers) > 0 {
		if len(slcFtothr) > 0 {
			fstKeytax := slcFtothr[0][0].Key
			fstValtax := slcFtothr[0][0].Value.(int)
			fstPrcntg := float64(fstValtax) / float64(now.Ftfuel+now.Ftaxyr+int32(tmpFrbase)) * 100
			fstStrval := fstKeytax + ":" + strconv.Itoa(int(fstPrcntg)) + "%"
			if len(slcFtothr) == 2 {
				scdKeytax := slcFtothr[1][0].Key
				scdValtax := slcFtothr[1][0].Value.(int)
				scdPrcntg := float64(scdValtax) / float64(now.Ftfuel) * 100
				scdStrval := scdKeytax + ":" + strconv.Itoa(int(scdPrcntg)) + "%"
				tmpOthers = append(tmpOthers, scdStrval)
				fstPrcntg := float64(fstValtax) / float64(tmpFrbase) * 100
				fstStrval = fstKeytax + ":" + strconv.Itoa(int(fstPrcntg)) + "%"
			} else {
				tmpOthers = append(tmpOthers, fstStrval)
			}
		}
		fnlFrtaxs.Ftothr = strings.Join(tmpOthers, "/")
	}

	// Check now than prev frbase
	var intDatenw, _ = strconv.Atoi(time.Now().Format("060102"))
	if val, ist := sycFrtaxs.Load(taxPrmkey); ist {
		if prv, mtc := val.(mdlApndix.MdlApndixFrtaxsDtbase); mtc {
			prvTaxstr := fmt.Sprintf("yq:%v/apt:%v/p4:%v/yr:%v/id:%v",
				prv.Ftfuel, prv.Ftaptx, prv.Ftiwjr, prv.Ftaxyr, prv.Ftppnx)
			nowTaxstr := fmt.Sprintf("yq:%v/apt:%v/p4:%v/yr:%v/id:%v",
				now.Ftfuel, now.Ftaptx, now.Ftiwjr, now.Ftaxyr, now.Ftppnx)
			fnlFrtaxs.Datend, fnlFrtaxs.Hstory = fncApndix.FncApndixFormatHstory(prvTaxstr,
				nowTaxstr, prv.Hstory, prv.Datend, int32(intDatenw))
		}
	} else {
		fnlFrtaxs.Datend = int32(intDatenw)
	}

	// Return final data
	sycFrtaxs.Store(taxPrmkey, fnlFrtaxs)
	return []mongo.WriteModel{mongo.NewUpdateOneModel().
		SetFilter(bson.M{"prmkey": fnlFrtaxs.Prmkey}).
		SetUpdate(bson.M{"$set": fnlFrtaxs}).
		SetUpsert(true)}
}
