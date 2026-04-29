package fncSbrapi

import (
	mdlSbrapi "back/sbrapi/model"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get data LC, PUN, LDN Raw from sabre
func FncSbrapiViehstMainob(unqhdr mdlSbrapi.MdlSbrapiMsghdrParams,
	params mdlSbrapi.MdlSbrapiMsghdrApndix) ([]mongo.WriteModel, error) {
	mgoRawhst := []mongo.WriteModel{}

	// Declare variable
	rawDatefl, _ := time.Parse("060102", strconv.Itoa(int(params.Datefl)))
	ddmDatefl := strings.ToUpper(rawDatefl.Format("02Jan"))

	// Isi struktur data
	strComand := fmt.Sprintf("VIH%v/%v", params.Flnbfl, ddmDatefl)
	strOutput, err := FncSbrapiCmdscrMainob(unqhdr, strComand)
	if err != nil {
		return mgoRawhst, err
	}

	// Final data
	mgoRawhst = FncSbrapiViehstPrcess(strOutput, params)
	return mgoRawhst, nil
}

// Function Treatment for API LC AND PUN
func FncSbrapiViehstPrcess(output string, params mdlSbrapi.MdlSbrapiMsghdrApndix,
) []mongo.WriteModel {

	// Declare first output
	var mgoResult []mongo.WriteModel

	// Looping data
	outlne := strings.Split(output, "\n")
	for seqnce, outrow := range outlne {
		strDatefl := strconv.Itoa(int(params.Datefl))
		strSeqnce := fmt.Sprintf("%06d", seqnce)
		strPrmkey := strDatefl + params.Airlfl + params.Flnbfl + "NO" + strSeqnce
		mgoResult = append(mgoResult, mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": strPrmkey}).
			SetUpdate(bson.M{"$set": mdlSbrapi.MdlViehstRawtmpDtbase{
				Prmkey: strPrmkey,
				Datefl: params.Datefl,
				Airlfl: params.Airlfl,
				Flnbfl: params.Flnbfl,
				Rawdta: outrow,
			}}).SetUpsert(true))
	}

	// Return final data
	return mgoResult
}
