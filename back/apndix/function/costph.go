package fncApndix

import (
	mdlApndix "back/apndix/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Get Sync map data LC and PUN prev day
func FncApndixCostphMapobj() map[string]int64 {

	// Inisialisasi variabel
	fnldta := make(map[string]int64)

	// Select database and collection
	tablex := Client.Database(Dbases).Collection("apndix_costph")
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
		var object mdlApndix.MdlApndixCostphDtbase
		datarw.Decode(&object)
		fnldta[object.Airlfl] = object.Costph
	}

	// return data
	return fnldta
}
