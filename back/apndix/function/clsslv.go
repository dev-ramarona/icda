package fncApndix

import (
	mdlApndix "back/apndix/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Get Sync map data LC and PUN prev day
func FncApndixClssvlMapobj() map[string]mdlApndix.MdlApndixClsslvDtbase {

	// Inisialisasi variabel
	fnldta := map[string]mdlApndix.MdlApndixClsslvDtbase{}

	// Select database and collection
	tablex := Client.Database(Dbases).Collection("apndix_clsslv")
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
		var object mdlApndix.MdlApndixClsslvDtbase
		datarw.Decode(&object)
		fnldta[object.Clssfl] = object
	}

	// return data
	return fnldta
}
