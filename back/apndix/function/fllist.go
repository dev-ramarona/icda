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
func FncApndixFllistGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlApndix.MdlApdnixParamsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	var totidx = 0
	var slcobj any
	tablex := Client.Database(Dbases).Collection("apndix_fllist")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "timeup", Value: -1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Datefl_apndix != "" {
		intDatefl := FncApndixFormatDatein(inputx.Datefl_apndix)
		csvFilenm = append(csvFilenm, inputx.Datefl_apndix)
		mtchdt = append(mtchdt, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
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
		var slctmp = []mdlApndix.MdlApndixFllistFrntnd{}
		for rawDtaset.Next(contxt) {
			slcDtaset := mdlApndix.MdlApndixFllistFrntnd{}
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

// Download
func FncApndixFllistDownld(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("0601021504")}
	var inputx mdlApndix.MdlApdnixParamsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	tablex := Client.Database(Dbases).Collection("apndix_fllist")
	contxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}

	// Check if data Route all is isset
	if inputx.Datefl_apndix != "" {
		intDatefl := FncApndixFormatDatein(inputx.Datefl_apndix)
		csvFilenm = append(csvFilenm, inputx.Datefl_apndix)
		mtchdt = append(mtchdt, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
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
	c.Header("Content-Disposition", "attachment; filename=apndix_fllist_"+fnlFilenm+".csv")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")

	// Streaming file CSV ke client
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()
	writer.Write([]string{
		"Prmkey",
		"Airlfl",
		"Flnbfl",
		"Timeup",
		"Timefl",
		"Timerv",
		"Datefl",
		"Mnthfl",
		"Ndayfl",
		"Flstat",
		"Routfl",
		"Routac",
		"Flsarr",
		"Routmx",
		"Flhour",
		"Flrpdc",
		"Flgate",
		"Depart",
		"Arrivl",
		"Airtyp",
		"Aircnf",
		"Seatcn",
		"Autrzc",
		"Autrzy",
		"Bookdc",
		"Bookdy",
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
		var slcDtaset mdlApndix.MdlApndixFllistDtbase
		rawDtaset.Decode(&slcDtaset)
		strMnthfl := FncApndixFormatMnthot(int(slcDtaset.Mnthfl))
		strDatefl := FncApndixFormatDateot(int(slcDtaset.Datefl))
		strTimefl := FncApndixFormatTimeot(int(slcDtaset.Timefl))
		strTimerv := FncApndixFormatTimeot(int(slcDtaset.Timerv))
		strTimeup := FncApndixFormatTimeot(int(slcDtaset.Timeup))

		// Write to CSV
		writer.Write([]string{
			slcDtaset.Prmkey,
			slcDtaset.Airlfl,
			slcDtaset.Flnbfl,
			strTimeup,
			strTimefl,
			strTimerv,
			strDatefl,
			strMnthfl,
			slcDtaset.Ndayfl,
			slcDtaset.Flstat,
			slcDtaset.Routfl,
			slcDtaset.Routac,
			slcDtaset.Flsarr,
			slcDtaset.Routmx,
			fmt.Sprintf("%v", slcDtaset.Flhour),
			fmt.Sprintf("%v", slcDtaset.Flrpdc),
			slcDtaset.Flgate,
			slcDtaset.Depart,
			slcDtaset.Arrivl,
			slcDtaset.Airtyp,
			slcDtaset.Aircnf,
			slcDtaset.Seatcn,
			fmt.Sprintf("%v", slcDtaset.Autrzc),
			fmt.Sprintf("%v", slcDtaset.Autrzy),
			fmt.Sprintf("%v", slcDtaset.Bookdc),
			fmt.Sprintf("%v", slcDtaset.Bookdy),
		})

		// Flush every 1000row
		countr++
		if countr%mxflus == 0 {
			writer.Flush()
		}
	}
}
