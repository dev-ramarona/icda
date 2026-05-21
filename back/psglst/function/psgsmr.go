package fncPsglst

import (
	fncApndix "back/apndix/function"
	mdlApndix "back/apndix/model"
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
	var joinrd = false
	var slcobj any
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var keywrd = fncApndix.FncApndixKeywrdMapobj("provnc")
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	mtchdt = append(mtchdt, bson.D{
		{Key: "totpax", Value: bson.D{{Key: "$ne", Value: 0}}}})
	mtchdt = append(mtchdt, bson.D{
		{Key: "totpax", Value: bson.D{{Key: "$exists", Value: true}}}})
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
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "flnbfl", Value: inputx.Flnbfl_psgsmr}},
			bson.D{{Key: "flnbjn", Value: inputx.Flnbfl_psgsmr}}}}})
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
	fmt.Println(inputx.Keywrd_psgsmr)
	if inputx.Keywrd_psgsmr != "" {
		var slcKeywrd []string
		if err := json.Unmarshal([]byte(inputx.Keywrd_psgsmr), &slcKeywrd); err == nil {
			for _, key := range slcKeywrd {
				fmt.Println("xxxxx", key)
				if val, ist := keywrd[key]; ist && val != "ALL" {
					fmt.Println("adwaw", key)
					csvFilenm = append(csvFilenm, inputx.Keywrd_psgsmr)
					mtchdt = append(mtchdt, bson.D{{Key: "provnc",
						Value: val}})
				}
			}
		} else {
			fmt.Println(err)
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
		{Key: "flnbfx", Value: bson.D{
			{Key: "$ifNull", Value: bson.A{
				bson.D{{Key: "$cond", Value: bson.A{
					bson.D{{Key: "$eq", Value: bson.A{"$flnbjn", ""}}},
					nil, "$flnbjn"}}}, "$flnbfl"}}}}}}}
	var addpry = bson.D{{Key: "$addFields", Value: bson.D{
		{Key: "ismtch", Value: bson.D{{Key: "$cond", Value: bson.A{
			bson.D{{Key: "$eq", Value: bson.A{"$flnbfl", "$flnbjn"}}}, 1, 0}}}}}}}
	var srtpry = bson.D{{Key: "$sort", Value: bson.D{
		{Key: "ismtch", Value: -1}}}}
	var grupfn = bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: bson.D{
			{Key: "flnbfl", Value: "$flnbfx"},
			{Key: "datefl", Value: "$datefl"},
			{Key: "airlfl", Value: "$airlfl"},
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
				mtchfn, addfld, addpry, srtpry, grupfn, rplrot}
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

			// Get load factor
			totSeatcn, err := strconv.ParseFloat(slcDtaset.Seatcn, 64)
			if err != nil {
				slcSeatcn := strings.Split(slcDtaset.Seatcn, "/")
				for _, v := range slcSeatcn {
					if seatcn, err := strconv.ParseFloat(v, 64); err == nil {
						totSeatcn += seatcn
					}
				}
			}
			slcDtaset.Loadfc = float64(slcDtaset.Totpax) / totSeatcn
			slcDtaset.Totrph = (slcDtaset.Totnta + slcDtaset.Tottyq + slcDtaset.Tottyr) / slcDtaset.Flhour
			slcDtaset.Totrev = slcDtaset.Totnta + slcDtaset.Tottyq
			slcDtaset.Totcph = float64(slcDtaset.Costph) * slcDtaset.Flhour
			slctmp = append(slctmp, slcDtaset)
		}
		slcobj = slctmp
	}()

	// Get Join data
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Find sample join data by datefl
		objJoindt := mdlApndix.MdlApndixFljoinFrntnd{}
		tablej := fncApndix.Client.Database(fncApndix.Dbases).Collection("apndix_fljoin")
		err := tablej.FindOne(contxt, bson.M{"datefl": intDatefl}).Decode(&objJoindt)
		if err == nil {
			joinrd = true
		}
	}()

	// Waiting until all go done
	wg.Wait()

	// Return final output
	c.JSON(200, gin.H{"totdta": totidx, "arrdta": slcobj, "joinrd": joinrd})
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
	var keywrd = fncApndix.FncApndixKeywrdMapobj("provnc")
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}

	// Check if data Route all is isset
	mtchdt = append(mtchdt, bson.D{
		{Key: "totpax", Value: bson.D{{Key: "$ne", Value: 0}}}})
	mtchdt = append(mtchdt, bson.D{
		{Key: "totpax", Value: bson.D{{Key: "$exists", Value: true}}}})
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
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "flnbfl", Value: inputx.Flnbfl_psgsmr}},
			bson.D{{Key: "flnbjn", Value: inputx.Flnbfl_psgsmr}}}}})
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
	fmt.Println(inputx.Keywrd_psgsmr)
	if inputx.Keywrd_psgsmr != "" {
		var slcKeywrd []string
		if err := json.Unmarshal([]byte(inputx.Keywrd_psgsmr), &slcKeywrd); err == nil {
			for _, key := range slcKeywrd {
				fmt.Println("xxxxx", key)
				if val, ist := keywrd[key]; ist && val != "ALL" {
					fmt.Println("adwaw", key)
					csvFilenm = append(csvFilenm, inputx.Keywrd_psgsmr)
					mtchdt = append(mtchdt, bson.D{{Key: "provnc",
						Value: val}})
				}
			}
		} else {
			fmt.Println(err)
		}
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
	if inputx.Isitjn_psgsmr == "Combined" {
		csvFilenm = append(csvFilenm, "Combined")
	}

	// Set header untuk file CSV
	fnlFilenm := strings.Join(csvFilenm, "_")
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=Psglst_Summary"+fnlFilenm+".csv")
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
		"timefl",
		"mnthfl",
		"flstat",
		"seatcn",
		"airtyp",
		"flhour",
		"loadfc",
		"totpax",
		"totnta",
		"tottyq",
		"tottyr",
		"totrev",
		"totfae",
		"totqfr",
		"totrph",
		"totcph",
	})
	writer.Flush()

	// Final group pipeline
	var addfld = bson.D{{Key: "$addFields", Value: bson.D{
		{Key: "flnbfx", Value: bson.D{
			{Key: "$ifNull", Value: bson.A{
				bson.D{{Key: "$cond", Value: bson.A{
					bson.D{{Key: "$eq", Value: bson.A{"$flnbjn", ""}}},
					nil, "$flnbjn"}}}, "$flnbfl"}}}}}}}
	var addpry = bson.D{{Key: "$addFields", Value: bson.D{
		{Key: "ismtch", Value: bson.D{{Key: "$cond", Value: bson.A{
			bson.D{{Key: "$eq", Value: bson.A{"$flnbfl", "$flnbjn"}}}, 1, 0}}}}}}}
	var srtpry = bson.D{{Key: "$sort", Value: bson.D{
		{Key: "isMatch", Value: -1}}}}
	var grupfn = bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: bson.D{
			{Key: "flnbfl", Value: "$flnbfx"},
			{Key: "datefl", Value: "$datefl"},
			{Key: "airlfl", Value: "$airlfl"},
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

	// Get All Match Data
	pipeln := mongo.Pipeline{mtchfn, sortdt}
	if inputx.Isitjn_psgsmr == "Combined" {
		pipeln = mongo.Pipeline{mtchfn,
			mtchfn, addfld, addpry, srtpry, grupfn, rplrot}
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
		strFlhour := fncApndix.FncApndixRevrseFlhour(slcDtaset.Flhour)
		strTimefl := fncApndix.FncApndixFormatTimeot(int(slcDtaset.Timefl))
		totSeatcn, err := strconv.ParseFloat(slcDtaset.Seatcn, 64)
		if err != nil {
			slcSeatcn := strings.Split(slcDtaset.Seatcn, "/")
			for _, v := range slcSeatcn {
				if seatcn, err := strconv.ParseFloat(v, 64); err == nil {
					totSeatcn += seatcn
				}
			}
		}

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
			strTimefl,
			strMnthfl,
			slcDtaset.Flstat,
			slcDtaset.Seatcn,
			slcDtaset.Airtyp,
			strFlhour,
			fmt.Sprintf("%v", float64(slcDtaset.Totpax)/totSeatcn),
			fmt.Sprintf("%v", slcDtaset.Totpax),
			fmt.Sprintf("%v", slcDtaset.Totnta),
			fmt.Sprintf("%v", slcDtaset.Tottyq),
			fmt.Sprintf("%v", slcDtaset.Tottyr),
			fmt.Sprintf("%v", slcDtaset.Totnta+slcDtaset.Tottyq),
			fmt.Sprintf("%v", slcDtaset.Totfae),
			fmt.Sprintf("%v", slcDtaset.Totqfr),
			fmt.Sprintf("%v", (slcDtaset.Totnta+slcDtaset.Tottyq+slcDtaset.Tottyr)/slcDtaset.Flhour),
			fmt.Sprintf("%v", slcDtaset.Totcph),
		})

		// Flush every 1000row
		countr++
		if countr%mxflus == 0 {
			writer.Flush()
		}
	}
}
