package fncSbrapi

import (
	fncApndix "back/apndix/function"
	mdlSbrapi "back/sbrapi/model"
	"encoding/xml"
	"fmt"
	"time"
)

// Login Sabre and create Session API
func FncSbrapiCrtssnMainob(carrierCode string) (mdlSbrapi.MdlSbrapiMsghdrParams, error) {

	// Declare first output
	convid := fncApndix.FncApndixCreateCduuid()
	mssgid := fncApndix.FncApndixCreateCduuid()
	timefm := time.Now().UTC().Format(time.RFC3339)
	var objhdr = mdlSbrapi.MdlSbrapiMsghdrParams{
		Convid: fmt.Sprintf("V1@%s@%s", convid, fncApndix.Pcckey),
		Mssgid: fmt.Sprintf("mid:%s@clientofsabre.com", mssgid),
		Timefm: timefm,
		Bsttkn: "Failed",
	}

	// Isi struktur data
	bdyCrtssn := mdlSbrapi.MdlSbrapiCrtssnReqenv{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: mdlSbrapi.MdlSbrapiCrtssnReqhdr{
			MessageHeader: FncSbrapiMsghdrMainob(fncApndix.Pcckey,
				"Create Session API", "SessionCreateRQ", objhdr),
			Security: mdlSbrapi.MdlSbrapiCrtssnReqscr{
				UsernameToken: mdlSbrapi.MdlSbrapiCrtssnRequsr{
					Username: fncApndix.Usrnme, Password: fncApndix.Psswrd,
					Organization: carrierCode, Domain: carrierCode,
				},
				XmlnsWsse: "http://schemas.xmlsoap.org/ws/2002/12/secext",
			},
		},
		Body: mdlSbrapi.MdlSbrapiCrtssnReqbdy{
			SessionCreateRQ: mdlSbrapi.MdlSbrapiCrtssnReqcrt{
				Version: "1.0.0",
				POS: mdlSbrapi.MdlSbrapiCrtssnReqpos{Source: mdlSbrapi.MdlSbrapiCrtssnReqsrc{
					PseudoCityCode: fncApndix.Pcckey}},
				Xmlns: "http://webservices.sabre.com",
			},
		},
	}

	// Read response
	rspssn, err := FncSbrapiMsghdrXmldta(bdyCrtssn)
	if err != nil {
		return objhdr, err
	}

	// Parsing XML ke dalam struktur Go
	var envssn mdlSbrapi.MdlSbrapiCrtssnRspenv
	err = xml.Unmarshal([]byte(rspssn), &envssn)
	if err != nil {
		return objhdr, err
	}

	// Return non error data
	objhdr.Bsttkn = envssn.Header.Security.BinarySecurityToken.Token
	return objhdr, nil

}

// Get multiple session
func FncSbrapiCrtssnMultpl(airline string, countx int) ([]mdlSbrapi.MdlSbrapiMsghdrParams, error) {
	sessions := make([]mdlSbrapi.MdlSbrapiMsghdrParams, 0, countx)
	for i := 0; i < countx; i++ {
		objhdr, err := FncSbrapiCrtssnMainob(airline)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, objhdr)
	}
	return sessions, nil
}
