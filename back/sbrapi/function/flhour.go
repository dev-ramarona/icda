package fncSbrapi

import (
	fncApndix "back/apndix/function"
	mdlApndix "back/apndix/model"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Get data LC, PUN, LDN Raw from sabre
func FncSbrapiFlhourMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	sycFlhour *sync.Map,
	params mdlSbrapi.MdlSbrapiMsghdrApndix,
) ([]mdlApndix.MdlApndixFlhourDtbase, error) {

	// Declare variable
	rawDatefl, _ := time.Parse("060102", strconv.Itoa(int(params.Datefl)))
	ddmDatefl := strings.ToUpper(rawDatefl.Format("02Jan"))
	strComand := fmt.Sprintf("V*%s/%s", params.Flnbfl, ddmDatefl)

	// Isi struktur data
	fnlLcnpun := []mdlApndix.MdlApndixFlhourDtbase{}
	strOutput, err := FncSbrapiCmdxmlMainob(unqhdr, strComand)
	if err != nil {
		return fnlLcnpun, err
	}

	//  Final data
	fnlLcnpun, err = FncSbrapiFlhourPrcess(strOutput, sycFlhour, params)
	if err != nil {
		return fnlLcnpun, err
	}
	return fnlLcnpun, nil
}

// Treatment data raw flight hour
func FncSbrapiFlhourPrcess(rawxml []byte, sycFlhour *sync.Map,
	apndix mdlSbrapi.MdlSbrapiMsghdrApndix) (
	[]mdlApndix.MdlApndixFlhourDtbase, error) {
	fnlFlhour := []mdlApndix.MdlApndixFlhourDtbase{}

	// Parsing XML ke dalam struktur Go
	rspFlhour := mdlSbrapi.MdlSbrapiFlhourRspenv{}
	err := xml.Unmarshal([]byte(rawxml), &rspFlhour)
	if err != nil {
		return fnlFlhour, err
	}

	// Looping all flight list
	xmlFlhour := rspFlhour.Body.SabreCommandLLSRS.XML_Content
	for _, objFlhour := range xmlFlhour.AIRAALSADSKED0.SKD001 {

		// Initialisasi data
		nowDepart := objFlhour.BoardPoint
		nowArrivl := objFlhour.DestinationAirportCode
		nowRoutef := nowDepart + "-" + nowArrivl
		nowPrmkey := apndix.Airlfl + apndix.Flnbfl + nowRoutef
		nowFlhour := objFlhour.ElapsedTime
		nowTimefl := objFlhour.ScheduledDepartureTime
		nowTimerv := objFlhour.ArrivalTime

		// Convert str 12.55 time to decimal
		floFlhour, err := fncApndix.FncApndixConvrtFlhour(nowFlhour)
		if err != nil {
			return fnlFlhour, err
		}

		// Convert Time flight 920A / 1230P to string decimal time
		strTimefl, err := fncApndix.FncApndixConvrtFltime(nowTimefl)
		if err != nil {
			return fnlFlhour, err
		}
		strTimerv, err := fncApndix.FncApndixConvrtFltime(nowTimerv)
		if err != nil {
			return fnlFlhour, err
		}
		intDaterv := func() int32 {
			intDaterv := apndix.Datefl
			intTimefl, _ := strconv.Atoi(strTimefl)
			intTimerv, _ := strconv.Atoi(strTimerv)
			if intTimefl > intTimerv {
				fmtDaterv, _ := time.Parse("060102", strconv.Itoa(int(apndix.Datefl)))
				strDaterv, _ := strconv.Atoi(fmtDaterv.AddDate(0, 0, +1).Format("060102"))
				intDaterv = int32(strDaterv)

			}
			return intDaterv
		}()
		intTimefl, _ := strconv.Atoi(strconv.Itoa(int(apndix.Datefl)) + strTimefl)
		intTimerv, _ := strconv.Atoi(strconv.Itoa(int(intDaterv)) + strTimerv)
		intTimenw, _ := strconv.Atoi(time.Now().Format("0601021504"))
		intDatenw, _ := strconv.Atoi(time.Now().Format("060102"))

		// Get air miles
		strAirmls := strings.Trim(objFlhour.AirMilesFlown, " ")
		intAirmls, err := strconv.Atoi(strAirmls)
		if err != nil {
			intAirmls = 0
		}

		// Check now than prev flight hour
		var nowDatend = int32(intDatenw)
		var nowHstory = string("")
		if val, ist := sycFlhour.Load(apndix.Airlfl + apndix.Flnbfl + nowRoutef); ist {
			if get, mtc := val.(mdlApndix.MdlApndixFlhourDtbase); mtc {
				nowDatend, nowHstory = fncApndix.FncApndixFormatHstory(get.Flhour,
					floFlhour, get.Hstory, get.Datend, int32(intDatenw))
			}
		}

		// Push to the final data
		var outFlhour = mdlApndix.MdlApndixFlhourDtbase{
			Prmkey: nowPrmkey,
			Airlfl: apndix.Airlfl,
			Routfl: nowRoutef,
			Flnbfl: apndix.Flnbfl,
			Flhour: floFlhour,
			Timefl: int64(intTimefl),
			Timerv: int64(intTimerv),
			Timeup: int64(intTimenw),
			Dateup: int32(intDatenw),
			Datend: nowDatend,
			Airtyp: objFlhour.EquipmentCode,
			Airmls: int32(intAirmls),
			Hstory: nowHstory,
		}
		fnlFlhour = append(fnlFlhour, outFlhour)
	}

	// Final return
	return fnlFlhour, nil
}
