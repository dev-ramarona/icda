package fncApndix

import (
	mdlApndix "back/apndix/model"
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Get Sync map data LC and PUN prev day
func FncApndixFlhourSycmap() *sync.Map {

	// Inisialisasi variabel
	fnldta := &sync.Map{}

	// Select database and collection
	tablex := Client.Database(Dbases).Collection("apndix_flhour")
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
		var object mdlApndix.MdlApndixFlhourDtbase
		datarw.Decode(&object)
		fnldta.Store(object.Prmkey, object)
		fnldta.Store(object.Routfl, object)
	}

	// return data
	return fnldta
}
