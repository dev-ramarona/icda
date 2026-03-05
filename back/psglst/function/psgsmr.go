package fncPsglst

import (
	fncApndix "back/apndix/function"
	mdlPsglst "back/psglst/model"

	"context"
	"encoding/csv"
	"encoding/json"
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
func FncPsglstPsgsmrGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlPsglst.MdlPsglstParamsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Treatment date number
	intDatefl := 0
	if inputx.Datefl_psgsmr != "" {
		strDatefl, _ := time.Parse("2006-01-02", inputx.Datefl_psgsmr)
		intDatefl, _ = strconv.Atoi(strDatefl.Format("060102"))
	}

	// Select db and context to do
	var totidx = 0
	var slcobj any
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Datefl_psgsmr != "" {
		csvFilenm = append(csvFilenm, strconv.Itoa(intDatefl))
		mtchdt = append(mtchdt, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
	if inputx.Airlfl_psgsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Airlfl_psgsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: inputx.Airlfl_psgsmr}})
	}
	if inputx.Flnbfl_psgsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl_psgsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "flnbfl",
			Value: inputx.Flnbfl_psgsmr}})
	}
	if inputx.Depart_psgsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Depart_psgsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "depart",
			Value: inputx.Depart_psgsmr}})
	}
	if inputx.Routfl_psgsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Routfl_psgsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: inputx.Routfl_psgsmr}})
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
			bson.D{{Key: "$skip", Value: (max(inputx.Pagenw_psgsmr, 1) - 1) * inputx.Limitp_psgsmr}},
			bson.D{{Key: "$limit", Value: inputx.Limitp_psgsmr}},
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, pipeln)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		var slctmp = []mdlPsglst.MdlPsglstPsgsmrFrtend{}
		for rawDtaset.Next(contxt) {
			slcDtaset := mdlPsglst.MdlPsglstPsgsmrFrtend{}
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

// Download PNR Detail all
func FncPsglstPsgsmrDownld(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("0601021504")}
	rawipt := c.PostForm("data")
	if rawipt == "" {
		c.String(400, "missing data")
		return
	}
	var inputx mdlPsglst.MdlPsglstParamsInputx
	if err := json.Unmarshal([]byte(rawipt), &inputx); err != nil {
		c.String(400, "invalid data")
		return
	}

	// Treatment date number
	intDatefl := 0
	if inputx.Datefl_psgsmr != "" {
		strDatefl, _ := time.Parse("2006-01-02", inputx.Datefl_psgsmr)
		intDatefl, _ = strconv.Atoi(strDatefl.Format("060102"))
	}

	// Select db and context to do
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}

	// Check if data Route all is isset
	if inputx.Datefl_psgsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Datefl_psgsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
	if inputx.Airlfl_psgsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Airlfl_psgsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: inputx.Airlfl_psgsmr}})
	}
	if inputx.Flnbfl_psgsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl_psgsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "flnbfl",
			Value: inputx.Flnbfl_psgsmr}})
	}
	if inputx.Depart_psgsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Depart_psgsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "depart",
			Value: inputx.Depart_psgsmr}})
	}
	if inputx.Routfl_psgsmr != "" {
		csvFilenm = append(csvFilenm, inputx.Routfl_psgsmr)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: inputx.Routfl_psgsmr}})
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
	c.Header("Content-Disposition", "attachment; filename=Psglst_Detail_"+fnlFilenm+".csv")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")

	// Streaming file CSV ke client
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()
	writer.Write([]string{
		"prmkey",
		"airlfl",
		"depart",
		"flnbfl",
		"routfl",
		"ndayfl",
		"datefl",
		"mnthfl",
		"flstat",
		"seatcn",
		"airtyp",
		"flhour",
		"totnta",
		"tottyq",
		"totpax",
		"totfae",
		"totqfr",
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
		var slcDtaset mdlPsglst.MdlPsglstPsgsmrDtbase
		rawDtaset.Decode(&slcDtaset)

		// Write to CSV
		writer.Write([]string{
			slcDtaset.Prmkey,
			slcDtaset.Airlfl,
			slcDtaset.Depart,
			slcDtaset.Flnbfl,
			slcDtaset.Routfl,
			slcDtaset.Ndayfl,
			fmt.Sprintf("%v", slcDtaset.Datefl),
			fmt.Sprintf("%v", slcDtaset.Mnthfl),
			slcDtaset.Flstat,
			slcDtaset.Seatcn,
			slcDtaset.Airtyp,
			fmt.Sprintf("%v", slcDtaset.Flhour),
			fmt.Sprintf("%v", slcDtaset.Totnta),
			fmt.Sprintf("%v", slcDtaset.Tottyq),
			fmt.Sprintf("%v", slcDtaset.Totpax),
			fmt.Sprintf("%v", slcDtaset.Totfae),
			fmt.Sprintf("%v", slcDtaset.Totqfr),
		})

		// Flush every 1000row
		countr++
		if countr%mxflus == 0 {
			writer.Flush()
		}
	}
}
