package fncApndix

import (
	mdlApndix "back/apndix/model"
	mdlSbrapi "back/sbrapi/model"
	"context"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get Sync map data LC and PUN prev day
func FncApndixFlnbflSycmap() *sync.Map {

	// Inisialisasi variabel
	fnldta := &sync.Map{}

	// Select database and collection
	tablex := Client.Database(Dbases).Collection("apndix_flnbls")
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
		var object mdlApndix.MdlApndixFlnblsDtbase
		datarw.Decode(&object)
		fnldta.Store(object.Prmkey, object)
	}

	// return data
	return fnldta
}

// Treatment to push database
func FncApndixFlnbflPrcess(sycFlnbfl *sync.Map, objParams mdlSbrapi.MdlSbrapiMsghdrApndix,
	prmkey, routfl string) []mongo.WriteModel {
	var intDatenw, _ = strconv.Atoi(time.Now().Format("060102"))
	var nowDatend = int32(intDatenw)
	var nowHstory = string("")
	if val, ist := sycFlnbfl.Load(prmkey); ist {
		if get, mtc := val.(mdlApndix.MdlApndixFlnblsDtbase); mtc {
			nowDatend, nowHstory = FncApndixFormatHstory(get.Routfl,
				routfl, get.Hstory, get.Datefl, int32(intDatenw))
		}
	}
	return []mongo.WriteModel{mongo.NewUpdateOneModel().
		SetFilter(bson.M{"prmkey": prmkey}).
		SetUpdate(bson.M{"$set": mdlApndix.MdlApndixFlnblsDtbase{
			Prmkey: prmkey,
			Airlfl: objParams.Airlfl,
			Flnbfl: objParams.Flnbfl,
			Routfl: routfl,
			Datefl: nowDatend,
			Hstory: nowHstory}}).SetUpsert(true)}
}
