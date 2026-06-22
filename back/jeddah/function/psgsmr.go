package fncJeddah

import (
	fncApndix "back/apndix/function"
	mdlJeddah "back/jeddah/model"

	"context"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get Detail PNR from database
func FncJeddahPnrsmrGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlJeddah.MdlJeddahPramsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	var totidx = 0
	var slcobj any
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("jeddah_pnrsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	mtchdt = append(mtchdt, bson.D{
		{Key: "totpax", Value: bson.D{{Key: "$ne", Value: 0}}}})
	mtchdt = append(mtchdt, bson.D{
		{Key: "totpax", Value: bson.D{{Key: "$exists", Value: true}}}})
	if inputx.Airlfl_jeddah != "" {
		slcAirlfl := strings.Split(inputx.Airlfl_jeddah, "-")
		csvFilenm = append(csvFilenm, inputx.Airlfl_jeddah)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: bson.M{"$in": slcAirlfl}}})
	}
	if inputx.Flnbfl_jeddah != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl_jeddah)
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "flnbfl", Value: inputx.Flnbfl_jeddah}},
			bson.D{{Key: "flnbjn", Value: inputx.Flnbfl_jeddah}}}}})
	}
	if inputx.Depart_jeddah != "" {
		csvFilenm = append(csvFilenm, inputx.Depart_jeddah)
		mtchdt = append(mtchdt, bson.D{{Key: "depart",
			Value: inputx.Depart_jeddah}})
	}
	if inputx.Routfl_jeddah != "" {
		csvFilenm = append(csvFilenm, inputx.Routfl_jeddah)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: inputx.Routfl_jeddah}})
	}

	// Final match pipeline
	var mtchfn bson.D
	if len(mtchdt) != 0 {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: mtchdt}}}}
	} else {
		fmt.Println("mtchblnk")
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
		nowPillne := mongo.Pipeline{
			mtchfn,
			sortdt,
			bson.D{{Key: "$skip", Value: (max(inputx.Pagenw_jeddah, 1) - 1) * inputx.Limitp_jeddah}},
			bson.D{{Key: "$limit", Value: inputx.Limitp_jeddah}},
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, nowPillne)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		var slctmp = []mdlJeddah.MdlJeddahPnrsmrDtbase{}
		for rawDtaset.Next(contxt) {
			slcDtaset := mdlJeddah.MdlJeddahPnrsmrDtbase{}
			rawDtaset.Decode(&slcDtaset)
			slcDtaset.Prmkey = slcDtaset.Pnrcde
			slctmp = append(slctmp, slcDtaset)
		}
		slcobj = slctmp
	}()

	// Waiting until all go done
	wg.Wait()

	// Return final output
	c.JSON(200, gin.H{"totdta": totidx, "arrdta": slcobj})
}

// Download PNR Detail all
func FncJeddahPnrsmrDownld(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlJeddah.MdlJeddahPramsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("jeddah_pnrsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}

	// Check if data Route all is isset
	mtchdt = append(mtchdt, bson.D{
		{Key: "totpax", Value: bson.D{{Key: "$ne", Value: 0}}}})
	mtchdt = append(mtchdt, bson.D{
		{Key: "totpax", Value: bson.D{{Key: "$exists", Value: true}}}})
	if inputx.Airlfl_jeddah != "" {
		slcAirlfl := strings.Split(inputx.Airlfl_jeddah, "-")
		csvFilenm = append(csvFilenm, inputx.Airlfl_jeddah)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: bson.M{"$in": slcAirlfl}}})
	}
	if inputx.Flnbfl_jeddah != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl_jeddah)
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "flnbfl", Value: inputx.Flnbfl_jeddah}},
			bson.D{{Key: "flnbjn", Value: inputx.Flnbfl_jeddah}}}}})
	}
	if inputx.Depart_jeddah != "" {
		csvFilenm = append(csvFilenm, inputx.Depart_jeddah)
		mtchdt = append(mtchdt, bson.D{{Key: "depart",
			Value: inputx.Depart_jeddah}})
	}
	if inputx.Routfl_jeddah != "" {
		csvFilenm = append(csvFilenm, inputx.Routfl_jeddah)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: inputx.Routfl_jeddah}})
	}

	// Final match pipeline
	var mtchfn bson.D
	var fltrfn bson.D
	if len(mtchdt) != 0 {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: mtchdt}}}}
		fltrfn = bson.D{{Key: "$and", Value: mtchdt}}
	} else {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Set header untuk file CSV
	fnlFilenm := strings.Join(csvFilenm, "_")
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=PNR_Jeddah_Summary"+fnlFilenm+".csv")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// Get total data count
	totrow, err := tablex.CountDocuments(contxt, fltrfn)
	if err == nil {
		writer.Write([]string{"Total data: " + strconv.Itoa(int(totrow))})
		writer.Flush()
	}

	// Streaming file CSV ke client
	writer.Write([]string{
		"pnrcde",
		"pnrsrc",
		"agtnme",
		"flnbsg",
		"routsg",
		"clssbk",
		"timefl",
		"timerv",
		"spltfr",
		"spltto",
		"totpax",
		"totbok",
		"totcxl",
		"totisd",
		"totori",
	})
	writer.Flush()

	// Get All Match Data
	pipeln := mongo.Pipeline{mtchfn, sortdt}

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
		var slcDtaset mdlJeddah.MdlJeddahPnrsmrDtbase
		rawDtaset.Decode(&slcDtaset)
		strTimefl := fncApndix.FncApndixFormatTimeot(int(slcDtaset.Timefl))
		strTimerv := fncApndix.FncApndixFormatTimeot(int(slcDtaset.Timerv))

		// Write to CSV
		writer.Write([]string{
			slcDtaset.Pnrcde,
			slcDtaset.Pnrsrc,
			slcDtaset.Agtnme,
			slcDtaset.Flnbsg,
			slcDtaset.Routsg,
			slcDtaset.Clssbk,
			strTimefl,
			strTimerv,
			slcDtaset.Spltfr,
			slcDtaset.Spltto,
			fmt.Sprintf("%v", slcDtaset.Totpax),
			fmt.Sprintf("%v", slcDtaset.Totbok),
			fmt.Sprintf("%v", slcDtaset.Totcxl),
			fmt.Sprintf("%v", slcDtaset.Totisd),
			fmt.Sprintf("%v", slcDtaset.Totori),
		})

		// Flush every 1000row
		countr++
		if countr%mxflus == 0 {
			writer.Flush()
		}
	}
}
