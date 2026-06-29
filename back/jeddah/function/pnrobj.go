package fncJeddah

import (
	fncApndix "back/apndix/function"
	mdlJeddah "back/jeddah/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Get Sync map data Cost per hour
func FncJeddahPnrobjMapobj() map[string]mdlJeddah.MdlJeddahPnrsmrCmpare {

	// Inisialisasi variabel
	fnldta := make(map[string]mdlJeddah.MdlJeddahPnrsmrCmpare)

	// Select database and collection
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("jeddah_flnbls")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{})
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var object mdlJeddah.MdlJeddahPnrsmrDtbase
		datarw.Decode(&object)
		if _, ist := fnldta[object.Pnrcde]; !ist {
			fnldta[object.Pnrcde] = mdlJeddah.MdlJeddahPnrsmrCmpare{}
		}
		fnldta[object.Pnrcde] = mdlJeddah.MdlJeddahPnrsmrCmpare{
			Flnbsg: object.Flnbsg,
			Routsg: object.Routsg,
			Clsssg: object.Clsssg,
			Timesg: object.Timesg,
			Totisd: object.Totisd,
			Totbok: object.Totbok,
			Totcxl: object.Totcxl,
		}
	}

	// return data
	return fnldta
}
