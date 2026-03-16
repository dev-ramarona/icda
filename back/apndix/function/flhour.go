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

// Get flighthour sycmap
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

// Get object slice
func FncApndixFlhourGetall(c *gin.Context) {

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
	tablex := Client.Database(Dbases).Collection("apndix_flhour")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "timeup", Value: -1}}}}
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
		mtchdt = append(mtchdt, bson.D{{Key: "depart",
			Value: inputx.Depart}})
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
		var slctmp = []mdlApndix.MdlApndixFlhourFrntnd{}
		for rawDtaset.Next(contxt) {
			slcDtaset := mdlApndix.MdlApndixFlhourFrntnd{}
			rawDtaset.Decode(&slcDtaset)
			slctmp = append(slctmp, slcDtaset)
		}
		slcobj = slctmp
	}()

	// Waiting until all go done
	wg.Wait()

	// Return final output
	c.JSON(200, gin.H{"totdta": totidx, "arrdta": slcobj})
}

// Get Response Update database from input
func FncApndixFlhourUpdate(c *gin.Context) {

	// Bind JSON Body input to variable
	var inputx mdlApndix.MdlApndixFlhourInputx
	if err := c.BindJSON(&inputx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input" + err.Error()})
		fmt.Println(err.Error())
		return
	}

	// Get from input
	if inputx.Airlfl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Airlfl"})
		return
	}
	if inputx.Routfl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Routfl"})
		return
	}
	if inputx.Flnbfl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Flnbfl"})
		return
	}
	if inputx.Flhour == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Flhour"})
		return
	}
	if inputx.Timefl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Timefl"})
		return
	}
	if inputx.Timerv == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Timerv"})
		return
	}
	if inputx.Timeup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Timeup"})
		return
	}
	if inputx.Dateup == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Dateup"})
		return
	}
	if inputx.Datend == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Datend"})
		return
	}
	if inputx.Airtyp == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Airtyp"})
		return
	}
	if inputx.Airmls == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Airmls"})
		return
	}

	// Final push data
	prmkey := inputx.Airlfl + inputx.Flnbfl + inputx.Routfl
	flhour, _ := strconv.ParseFloat(inputx.Flhour, 64)
	airmls, _ := strconv.Atoi(inputx.Flhour)
	timefl := FncApndixFormatTimein(inputx.Timefl)
	timerv := FncApndixFormatTimein(inputx.Timerv)
	timeup := FncApndixFormatTimein(time.Now().Format("2006-01-02T15:04"))
	dateup := FncApndixFormatDatein(time.Now().Format("2006-01-02"))
	datend := FncApndixFormatDatein(time.Now().Format("2006-01-02"))
	finald := mdlApndix.MdlApndixFlhourDtbase{
		Prmkey: prmkey,
		Airlfl: inputx.Airlfl,
		Routfl: inputx.Routfl,
		Flnbfl: inputx.Flnbfl,
		Flhour: flhour,
		Timefl: timefl,
		Timerv: timerv,
		Timeup: timeup,
		Dateup: dateup,
		Datend: datend,
		Airtyp: inputx.Airtyp,
		Airmls: int32(airmls),
		Hstory: inputx.Hstory,
		Updtby: inputx.Updtby,
	}

	// Push updated data
	rsupdt := FncApndixBulkdbSingle([]mongo.WriteModel{
		mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": finald.Prmkey}).
			SetUpdate(bson.M{"$set": finald}).
			SetUpsert(true)}, "apndix_flhour")
	if rsupdt != nil {
		panic("Error Insert/Update to DB:" + rsupdt.Error())
	}

	// Send token to frontend
	c.JSON(200, "success")
}
