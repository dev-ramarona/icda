package fncApndix

import (
	mdlAllusr "back/allusr/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Get Sync map data keyword
func FncApndixKeywrdMapobj(sction string) map[string]string {

	// Inisialisasi variabel
	fnldta := make(map[string]string)

	// Select database and collection
	tablex := Client.Database(Dbases).Collection("apndix_keywrd")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	matchd := bson.M{}
	if sction != "" {
		matchd = bson.M{"sction": sction}
	}
	datarw, err := tablex.Find(contxt, matchd)
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var object mdlAllusr.MdlAllusrKeywrdDtbase
		datarw.Decode(&object)
		fnldta[object.Keywrd] = object.Detail
	}

	// return data
	return fnldta
}
