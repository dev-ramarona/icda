package fncApndix

import (
	mdlApndix "back/apndix/model"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get map object
func FncApndixProvncMapobj() map[string]string {

	// Inisialisasi variabel
	fnldta := map[string]string{}

	// Select database and collection
	tablex := Client.Database(Dbases).Collection("apndix_provnc")
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
		var object mdlApndix.MdlApndixProvncDtbase
		datarw.Decode(&object)
		fnldta[object.Routfl] = object.Provnc
	}

	// return data
	return fnldta
}

// Get object slice
func FncApndixProvncGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlApndix.MdlApdnixParamsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Treatment date number
	intDatefl := 0
	if inputx.Datefl != "" {
		strDatefl, _ := time.Parse("2006-01-02", inputx.Datefl)
		intDatefl, _ = strconv.Atoi(strDatefl.Format("060102"))
	}

	// Select db and context to do
	var totidx = 0
	var slcobj any
	tablex := Client.Database(Dbases).Collection("apndix_provnc")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Datefl != "" {
		csvFilenm = append(csvFilenm, strconv.Itoa(intDatefl))
		mtchdt = append(mtchdt, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
	if inputx.Airlfl != "" {
		csvFilenm = append(csvFilenm, inputx.Airlfl)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: inputx.Airlfl}})
	}
	if inputx.Depart != "" {
		csvFilenm = append(csvFilenm, inputx.Depart)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: bson.D{{Key: "$regex",
				Value: "^" + inputx.Depart}}}})
	}
	if inputx.Flnbfl != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl)
		mtchdt = append(mtchdt, bson.D{{Key: "flnbfl",
			Value: inputx.Flnbfl}})
	}
	if inputx.Routfl != "" {
		csvFilenm = append(csvFilenm, inputx.Routfl)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: inputx.Routfl}})
	}
	if inputx.Clssfl != "" {
		csvFilenm = append(csvFilenm, inputx.Clssfl)
		mtchdt = append(mtchdt, bson.D{{Key: "clssfl",
			Value: inputx.Clssfl}})
	}

	// Final match pipeline
	var mtchfn bson.D
	if len(mtchdt) != 0 {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: mtchdt}}}}
	} else {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Get Total Count Data
	wg.Add(1)
	go func() {
		defer wg.Done()
		nowPillne := mongo.Pipeline{
			mtchfn,
			bson.D{{Key: "$count", Value: "totalCount"}},
		}

		// Find user by username in database
		fmt.Println(nowPillne)
		rawDtaset, err := tablex.Aggregate(contxt, nowPillne)
		if err != nil {
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		var slcDtaset []bson.M
		if err = rawDtaset.All(contxt, &slcDtaset); err != nil {
			panic(err)
		}

		// Mengambil jumlah dokumen dari hasil
		if len(slcDtaset) > 0 {
			if count, ok := slcDtaset[0]["totalCount"].(int32); ok {
				totidx = int(count)
			}
		}
	}()

	// Get All Match Data
	wg.Add(1)
	go func() {
		defer wg.Done()
		pipeln := mongo.Pipeline{
			mtchfn,
			sortdt,
			bson.D{{Key: "$skip", Value: (max(inputx.Pagenw, 1) - 1) * inputx.Limitp}},
			bson.D{{Key: "$limit", Value: inputx.Limitp}},
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, pipeln)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		var slctmp = []mdlApndix.MdlApndixProvncFrntnd{}
		for rawDtaset.Next(contxt) {
			slcDtaset := mdlApndix.MdlApndixProvncDtbase{}
			rawDtaset.Decode(&slcDtaset)
			slctmp = append(slctmp, mdlApndix.MdlApndixProvncFrntnd{
				Prmkey: slcDtaset.Routfl,
				Routfl: slcDtaset.Routfl,
				Provnc: slcDtaset.Provnc,
				Updtby: slcDtaset.Updtby,
			})
		}
		slcobj = slctmp
	}()

	// Waiting until all go done
	wg.Wait()

	// Return final output
	c.JSON(200, gin.H{"totdta": totidx, "arrdta": slcobj})
}

// Get Response Update database from input
func FncApndixProvincUpdate(c *gin.Context) {

	// Bind JSON Body input to variable
	var inputx mdlApndix.MdlApndixProvncDtbase
	if err := c.BindJSON(&inputx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input" + err.Error()})
		fmt.Println(err.Error())
		return
	}

	// Get from input
	if inputx.Provnc == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Provnc"})
		return
	}
	if inputx.Routfl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Routfl"})
		return
	}
	if inputx.Updtby == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Routfl"})
		return
	}

	// Push updated data
	rsupdt := FncApndixBulkdbSingle([]mongo.WriteModel{
		mongo.NewUpdateOneModel().
			SetFilter(bson.M{"routfl": inputx.Routfl}).
			SetUpdate(bson.M{"$set": inputx}).
			SetUpsert(true)}, "apndix_provnc")
	if rsupdt != nil {
		panic("Error Insert/Update to DB:" + rsupdt.Error())
	}

	// Send token to frontend
	c.JSON(200, "success")
}
