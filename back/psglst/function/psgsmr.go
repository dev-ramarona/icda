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
	mtchdt = append(mtchdt, bson.D{
		{Key: "totpax", Value: bson.D{
			{Key: "$ne", Value: 0}}}})
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
	if inputx.Keywrd_psgsmr != "" && !strings.Contains(inputx.Keywrd_psgsmr, "REG ALL") {
		var slcKeywrd []string
		if err := json.Unmarshal([]byte(inputx.Keywrd_psgsmr), &slcKeywrd); err == nil {
			csvFilenm = append(csvFilenm, inputx.Keywrd_psgsmr)
			mtchdt = append(mtchdt, bson.D{{Key: "provnc",
				Value: bson.D{{Key: "$in", Value: slcKeywrd}}}})
		}
	}

	// Final match pipeline
	var mtchfn bson.D
	if len(mtchdt) != 0 {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: mtchdt}}}}
	} else {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{}}}
	}

	// Final group pipeline
	var addfld = bson.D{{Key: "$addFields", Value: bson.D{
		{Key: "flnb_fix", Value: bson.D{
			{Key: "$ifNull", Value: bson.A{
				bson.D{{Key: "$cond", Value: bson.A{
					bson.D{{Key: "$eq", Value: bson.A{"$flnbjn", ""}}},
					nil, "$flnbjn"}}}, "$flnbfl"}}}}}}}
	var grupfn = bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: bson.D{
			{Key: "flnbfl", Value: "$flnb_fix"},
			{Key: "datefl", Value: "$datefl"},
			{Key: "depart", Value: "$depart"},
		}},
		{Key: "data", Value: bson.D{{Key: "$first", Value: "$$ROOT"}}},
		{Key: "totpax", Value: bson.D{{Key: "$sum", Value: "$totpax"}}},
		{Key: "totnta", Value: bson.D{{Key: "$sum", Value: "$totnta"}}},
		{Key: "tottyq", Value: bson.D{{Key: "$sum", Value: "$tottyq"}}},
		{Key: "totfae", Value: bson.D{{Key: "$sum", Value: "$totfae"}}},
		{Key: "totrph", Value: bson.D{{Key: "$sum", Value: "$totrph"}}},
	}}}
	var rplrot = bson.D{{Key: "$replaceRoot", Value: bson.D{
		{Key: "newRoot", Value: bson.D{
			{Key: "$mergeObjects", Value: bson.A{
				"$data",
				bson.D{
					{Key: "flnbfl", Value: "$_id.flnbfl"}, // ⬅️ penting!
					{Key: "totpax", Value: "$totpax"},
					{Key: "totnta", Value: "$totnta"},
					{Key: "tottyq", Value: "$tottyq"},
					{Key: "totfae", Value: "$totfae"},
					{Key: "totrph", Value: "$totrph"},
				},
			}},
		}},
	}}}

	// Get Total Count Data
	wg.Add(1)
	go func() {
		defer wg.Done()
		nowPillne := mongo.Pipeline{mtchfn}
		if inputx.Isitjn_psgsmr == "Combined" {
			nowPillne = mongo.Pipeline{mtchfn, addfld, grupfn, rplrot}
		}
		nowPillne = append(nowPillne, bson.D{{Key: "$count", Value: "totalCount"}})

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
		pipeln := mongo.Pipeline{mtchfn}
		if inputx.Isitjn_psgsmr == "Combined" {
			pipeln = mongo.Pipeline{mtchfn,
				mtchfn, addfld, grupfn, rplrot}
		}
		pipeln = append(pipeln, sortdt)
		pipeln = append(pipeln, bson.D{{Key: "$skip", Value: (max(inputx.Pagenw_psgsmr, 1) - 1) * inputx.Limitp_psgsmr}})
		pipeln = append(pipeln, bson.D{{Key: "$limit", Value: inputx.Limitp_psgsmr}})

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
	c.Header("Content-Disposition", "attachment; filename=Psglst_Detail_"+fnlFilenm+".csv")
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
		"prmkey",
		"airlfl",
		"provnc",
		"depart",
		"flnbfl",
		"flnbjn",
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
		"totrph",
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
		strDatefl := fncApndix.FncApndixFormatDateot(int(slcDtaset.Datefl))
		strMnthfl := fncApndix.FncApndixFormatMnthot(int(slcDtaset.Mnthfl))

		// Write to CSV
		writer.Write([]string{
			slcDtaset.Prmkey,
			slcDtaset.Airlfl,
			slcDtaset.Provnc,
			slcDtaset.Depart,
			slcDtaset.Flnbfl,
			slcDtaset.Flnbjn,
			slcDtaset.Routfl,
			slcDtaset.Ndayfl,
			strDatefl,
			strMnthfl,
			slcDtaset.Flstat,
			slcDtaset.Seatcn,
			slcDtaset.Airtyp,
			fmt.Sprintf("%v", slcDtaset.Flhour),
			fmt.Sprintf("%v", slcDtaset.Totnta),
			fmt.Sprintf("%v", slcDtaset.Tottyq),
			fmt.Sprintf("%v", slcDtaset.Totpax),
			fmt.Sprintf("%v", slcDtaset.Totfae),
			fmt.Sprintf("%v", slcDtaset.Totqfr),
			fmt.Sprintf("%v", slcDtaset.Totrph),
		})

		// Flush every 1000row
		countr++
		if countr%mxflus == 0 {
			writer.Flush()
		}
	}
}
