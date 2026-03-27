package fncApndix

import (
	mdlApndix "back/apndix/model"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get Sync map data Class level for farebase
func FncApndixFrtaxsSycmap() *sync.Map {

	// Inisialisasi variabel
	fnldta := &sync.Map{}

	// Select database and collection
	tablex := Client.Database(Dbases).Collection("apndix_frtaxs")
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
		var object mdlApndix.MdlApndixFrtaxsDtbase
		datarw.Decode(&object)
		fnldta.Store(object.Prmkey, object)
	}

	// return data
	return fnldta
}

// Get object slice
func FncApndixFrtaxsGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlApndix.MdlApdnixParamsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	var totidx = 0
	var slcobj any
	tablex := Client.Database(Dbases).Collection("apndix_frtaxs")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "frtaxs", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Routfl_apndix != "" {
		csvFilenm = append(csvFilenm, inputx.Routfl_apndix)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: bson.D{{Key: "$regex",
				Value: "^" + inputx.Routfl_apndix}}}})
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
			bson.D{{Key: "$skip", Value: (max(inputx.Pagenw_apndix, 1) - 1) * inputx.Limitp_apndix}},
			bson.D{{Key: "$limit", Value: inputx.Limitp_apndix}},
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, pipeln)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		var slctmp = []mdlApndix.MdlApndixFrtaxsFrntnd{}
		for rawDtaset.Next(contxt) {
			slcDtaset := mdlApndix.MdlApndixFrtaxsFrntnd{}
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
func FncApndixFrtaxsUpdate(c *gin.Context) {

	// Bind JSON Body input to variable
	var inputx mdlApndix.MdlApndixFrtaxsDtbase
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
	if inputx.Cbinfl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Cbinfl"})
		return
	}
	if inputx.Routfl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Routfl"})
		return
	}
	if inputx.Ftppnx == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Ftppnx"})
		return
	}
	if inputx.Ftaptx == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Ftaptx"})
		return
	}
	if inputx.Ftfuel == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Ftfuel"})
		return
	}
	if inputx.Ftiwjr == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Ftiwjr"})
		return
	}
	if inputx.Ftaxyr == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Ftaxyr"})
		return
	}
	if inputx.Datend == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Datend"})
		return
	}
	if inputx.Ftothr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Ftothr"})
		return
	}

	// Final push data
	if inputx.Prmkey == "" || inputx.Prmkey == "add" {
		inputx.Prmkey = inputx.Airlfl + inputx.Routfl + inputx.Cbinfl
	}

	// Push updated data
	rsupdt := FncApndixBulkdbSingle([]mongo.WriteModel{
		mongo.NewUpdateOneModel().
			SetFilter(bson.M{"routfl": inputx.Prmkey}).
			SetUpdate(bson.M{"$set": inputx}).
			SetUpsert(true)}, "apndix_frtaxs")
	if rsupdt != nil {
		panic("Error Insert/Update to DB:" + rsupdt.Error())
	}

	// Send token to frontend
	c.JSON(200, "success")
}

// Download
func FncApndixFrtaxsDownld(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("0601021504")}
	rawipt := c.PostForm("data")
	if rawipt == "" {
		c.String(400, "missing data")
		return
	}
	var inputx mdlApndix.MdlApdnixParamsInputx
	if err := json.Unmarshal([]byte(rawipt), &inputx); err != nil {
		c.String(400, "invalid data")
		return
	}

	// Select db and context to do
	tablex := Client.Database(Dbases).Collection("apndix_frtaxs")
	contxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}

	// Check if data Route all is isset
	if inputx.Airlfl_apndix != "" {
		csvFilenm = append(csvFilenm, inputx.Airlfl_apndix)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: inputx.Airlfl_apndix}})
	}
	if inputx.Clssfl_apndix != "" {
		csvFilenm = append(csvFilenm, inputx.Clssfl_apndix)
		mtchdt = append(mtchdt, bson.D{{Key: "clssfl",
			Value: inputx.Clssfl_apndix}})
	}
	if inputx.Routfl_apndix != "" {
		csvFilenm = append(csvFilenm, inputx.Routfl_apndix)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: bson.D{{Key: "$regex",
				Value: "^" + inputx.Routfl_apndix}}}})
	}

	// Final match pipeline
	var mtchfn bson.D
	if len(mtchdt) != 0 {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: mtchdt}}}}
	} else {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Set header untuk file CSV
	fnlFilenm := strings.Join(csvFilenm, "_")
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=apndix_Frtaxs_"+fnlFilenm+".csv")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")

	// Streaming file CSV ke client
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()
	writer.Write([]string{
		"Prmkey",
		"Airlfl",
		"Cbinfl",
		"Routfl",
		"Ftppnx",
		"Ftaptx",
		"Ftfuel",
		"Ftiwjr",
		"Ftaxyr",
		"Datend",
		"Ftothr",
		"Hstory",
		"Updtby",
	})
	writer.Flush()

	// Get All Match Data
	pipeln := mongo.Pipeline{
		mtchfn,
		sortdt,
	}

	// Find user by username in database
	rawDtaset, err := tablex.Aggregate(contxt, pipeln)
	if err != nil {
		panic(err)
	}
	defer rawDtaset.Close(contxt)

	// Store to slice from raw bson
	mxflus := 5000
	countr := 0
	for rawDtaset.Next(contxt) {
		var slcDtaset mdlApndix.MdlApndixFrtaxsDtbase
		rawDtaset.Decode(&slcDtaset)

		// Write to CSV
		writer.Write([]string{
			slcDtaset.Prmkey,
			slcDtaset.Airlfl,
			slcDtaset.Cbinfl,
			slcDtaset.Routfl,
			fmt.Sprintf("%v", slcDtaset.Ftppnx),
			fmt.Sprintf("%v", slcDtaset.Ftaptx),
			fmt.Sprintf("%v", slcDtaset.Ftfuel),
			fmt.Sprintf("%v", slcDtaset.Ftiwjr),
			fmt.Sprintf("%v", slcDtaset.Ftaxyr),
			fmt.Sprintf("%v", slcDtaset.Datend),
			slcDtaset.Ftothr,
			slcDtaset.Hstory,
			slcDtaset.Updtby,
		})

		// Flush every 1000row
		countr++
		if countr%mxflus == 0 {
			writer.Flush()
		}
	}
}
