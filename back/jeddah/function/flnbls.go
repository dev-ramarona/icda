package fncJeddah

import (
	fncApndix "back/apndix/function"
	mdlJeddah "back/jeddah/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Get Sync map data Cost per hour
func FncJeddahFlnblsMapobj() map[string]map[string]mdlJeddah.MdlJeddahFlnblsDtbase {

	// Inisialisasi variabel
	fnldta := make(map[string]map[string]mdlJeddah.MdlJeddahFlnblsDtbase)

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
		var object mdlJeddah.MdlJeddahFlnblsDtbase
		datarw.Decode(&object)
		if _, ist := fnldta[object.Airlfl]; !ist {
			fnldta[object.Airlfl] = map[string]mdlJeddah.MdlJeddahFlnblsDtbase{}
		}
		prmkey := fmt.Sprintf("%v%v%v%v", object.Airlfl, object.Flnbfl, object.Depart, object.Datefl)
		fnldta[object.Airlfl][prmkey] = object
	}

	// return data
	return fnldta
}
