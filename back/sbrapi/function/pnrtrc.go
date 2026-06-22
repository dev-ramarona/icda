package fncSbrapi

import (
	mdlSbrapi "back/sbrapi/model"
	"strconv"
	"strings"
	"time"
)

// Get data LC, PUN, LDN Raw from sabre
func FncSbrapiPnrtrcMainob(actcde string, unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	params mdlSbrapi.MdlSbrapiMsghdrApndix) (map[string]mdlSbrapi.MdlSbrapiPnrtrcDtbase, error) {

	// Declare variable
	rawDatefl, _ := time.Parse("060102", strconv.Itoa(int(params.Datefl)))
	ddmDatefl := strings.ToUpper(rawDatefl.Format("02Jan"))
	strComand := actcde + params.Flnbfl + "/" + ddmDatefl + params.Depart

	// Isi struktur data
	fnlLcnpun := map[string]mdlSbrapi.MdlSbrapiPnrtrcDtbase{}
	strOutput, err := FncSbrapiCmdscrMainob(unqhdr, strComand)
	if err != nil {
		return fnlLcnpun, err
	}

	//  Final data
	fnlLcnpun = FncSbrapiPnrtrcPrcess(actcde, strOutput, params)
	return fnlLcnpun, nil
}

// Function Treatment for API LC AND PUN
func FncSbrapiPnrtrcPrcess(actcde, output string, params mdlSbrapi.MdlSbrapiMsghdrApndix,
) map[string]mdlSbrapi.MdlSbrapiPnrtrcDtbase {

	// Declare first output
	var fnlResult = map[string]mdlSbrapi.MdlSbrapiPnrtrcDtbase{}
	var fnlRoutfl = ""
	var rawTimenw = time.Now().Format("0601021504")
	var intTimenw, _ = strconv.Atoi(rawTimenw)

	// Looping data
	outlne := strings.Split(output, "\n")
	for _, outrow := range outlne {

		// Get route
		if len(outrow) == 6 {
			fnlRoutfl = outrow[:3] + "-" + outrow[3:]
			continue
		}

		// Skip end
		if outrow == "END" {
			continue
		}

		// Split data
		slcrow := strings.Split(outrow, ".")
		clnrow := []string{}
		for _, row := range slcrow {
			if strings.TrimSpace(row) != "" {
				clnrow = append(clnrow, row)
			}
		}
		getPnrcde := clnrow[3]
		getTotpax, _ := strconv.Atoi(strings.TrimSpace(clnrow[0][3:6]))
		getClssbk := clnrow[2]
		getAgtnme := strings.TrimSpace(clnrow[0][6:])
		cekIssued := ""
		if actcde == "LC" && strings.Contains(clnrow[1], "/") {
			cekIssued = "issued"
		} else if actcde == "LX" {
			getClssbk = clnrow[1]
		}

		fnlResult[getPnrcde] = mdlSbrapi.MdlSbrapiPnrtrcDtbase{
			Actcde: actcde,
			Agtnme: getAgtnme,
			Depart: params.Depart,
			Routfl: fnlRoutfl,
			Flnbfl: params.Flnbfl,
			Timenw: int64(intTimenw),
			Pnrcde: getPnrcde,
			Issued: cekIssued,
			Totpax: int32(getTotpax),
			Clssbk: getClssbk,
		}
		// // Push to database LC AND PUN
		// if len(clnrow) >= 3 {
		// 	totpax, _ := strconv.Atoi(strings.TrimSpace(clnrow[0][3:6]))
		// 	agtnme := strings.TrimSpace(clnrow[0][6:len(clnrow[0])])
		// 	pnrcde := clnrow[2]
		// 	clssfl := clnrow[1][:1]
		// 	if len(clnrow) == 4 {
		// 		pnrcde = clnrow[3]
		// 		clssfl = clnrow[2]
		// 	}

		// 	// Assign data
		// 	if _, ist := allpnr[pnrcde]; !ist {
		// 		allpnr[pnrcde] = true
		// 		result = append(result, mdlSbrapi.MdlSbrapiPnrtrcDtbase{
		// 			Prmkey: lcnpun + params.Airlfl + params.Flnbfl + params.Depart +
		// 				strconv.Itoa(int(params.Datefl)) + pnrcde + timenw[0:6],
		// 			Airlfl: params.Airlfl, Lcrpun: lcnpun, Totpax: totpax,
		// 			Flnbfl: params.Flnbfl, Depart: params.Depart,
		// 			Routfl: params.Routfl, Clssfl: strings.TrimSpace(clssfl),
		// 			Datefl: params.Datefl, Dateup: int32(intDatenw),
		// 			Timeup: int64(intTimenw), Agtnme: agtnme, Pnrcde: pnrcde})
		// 	}
		// }
	}

	// Return final data
	return fnlResult
}
