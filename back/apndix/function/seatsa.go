package fncApndix

import (
	mdlApndix "back/apndix/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get
func FncApndixSeatsaMapobj() map[string]mdlApndix.MdlApndixSeatsaDtbase {

	// Inisialisasi variabel
	fnldta := map[string]mdlApndix.MdlApndixSeatsaDtbase{}

	// Select database and collection
	tablex := Client.Database(Dbases).Collection("apndix_seatsa")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{},
		options.Find().SetSort(bson.D{{Key: "levelr", Value: 1}}))
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var object mdlApndix.MdlApndixSeatsaDtbase
		datarw.Decode(&object)
		fnldta[object.Airlfl+object.Seatcn+object.Seatpx] = object
	}

	// return data
	return fnldta
}
