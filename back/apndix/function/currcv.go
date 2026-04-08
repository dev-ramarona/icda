package fncApndix

import (
	mdlApndix "back/apndix/model"
	"context"
	"encoding/csv"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get object slice
func FncApndixCurrcvGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlApndix.MdlApdnixParamsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	var totidx = 0
	var slcobj any
	tablex := Client.Database(Dbases).Collection("apndix_currcv")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "timeup", Value: -1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Airlfl_apndix != "" {
		csvFilenm = append(csvFilenm, inputx.Airlfl_apndix)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: inputx.Airlfl_apndix}})
	}
	if inputx.Flnbfl_apndix != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl_apndix)
		mtchdt = append(mtchdt, bson.D{{Key: "flnbfl",
			Value: inputx.Flnbfl_apndix}})
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
		var slctmp = []mdlApndix.MdlApndixCurrcvFrntnd{}
		for rawDtaset.Next(contxt) {
			slcDtaset := mdlApndix.MdlApndixCurrcvFrntnd{}
			rawDtaset.Decode(&slcDtaset)
			slcDtaset.Prmkey = slcDtaset.Crcode
			slctmp = append(slctmp, slcDtaset)
		}
		slcobj = slctmp
	}()

	// Waiting until all go done
	wg.Wait()

	// Return final output
	c.JSON(200, gin.H{"totdta": totidx, "arrdta": slcobj})
}

// Download
func FncApndixCurrcvDownld(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("0601021504")}
	var inputx mdlApndix.MdlApdnixParamsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	tablex := Client.Database(Dbases).Collection("apndix_currcv")
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
	if inputx.Flnbfl_apndix != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl_apndix)
		mtchdt = append(mtchdt, bson.D{{Key: "flnbfl",
			Value: inputx.Flnbfl_apndix}})
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
	c.Header("Content-Disposition", "attachment; filename=apndix_currcv_"+fnlFilenm+".csv")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")

	// Streaming file CSV ke client
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()
	writer.Write([]string{
		"crctry",
		"crcode",
		"crname",
		"crrate",
		"datend",
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
		var slcDtaset mdlApndix.MdlApndixCurrcvDtbase
		rawDtaset.Decode(&slcDtaset)
		strDatend := FncApndixFormatDateot(int(slcDtaset.Datend))

		// Write to CSV
		writer.Write([]string{
			slcDtaset.Crctry,
			slcDtaset.Crcode,
			slcDtaset.Crname,
			fmt.Sprintf("%v", slcDtaset.Crrate),
			strDatend,
		})

		// Flush every 1000row
		countr++
		if countr%mxflus == 0 {
			writer.Flush()
		}
	}
}
