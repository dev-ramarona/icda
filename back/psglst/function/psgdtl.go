package fncPsglst

import (
	fncApndix "back/apndix/function"
	mdlPsglst "back/psglst/model"
	"math"
	"net/http"

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
func FncPsglstPsgdtlGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	strDvsion := c.Param("dvsion")
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlPsglst.MdlPsglstParamsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Treatment date number
	intDatefl := 0
	if inputx.Datefl_psgdtl != "" {
		strDatefl, _ := time.Parse("2006-01-02", inputx.Datefl_psgdtl)
		intDatefl, _ = strconv.Atoi(strDatefl.Format("060102"))
	}

	// Select db and context to do
	var totidx = 0
	var slcobj interface{}
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgdtl")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Datefl_psgdtl != "" {
		csvFilenm = append(csvFilenm, strconv.Itoa(intDatefl))
		mtchdt = append(mtchdt, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
	if inputx.Airlfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Airlfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: inputx.Airlfl_psgdtl}})
	}
	if inputx.Flnbfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "flnbfl",
			Value: inputx.Flnbfl_psgdtl}})
	}
	if inputx.Depart_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Depart_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "depart",
			Value: inputx.Depart_psgdtl}})
	}
	if inputx.Routfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Routfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: inputx.Routfl_psgdtl}})
	}
	if inputx.Pnrcde_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Pnrcde_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "pnrcde",
			Value: inputx.Pnrcde_psgdtl}})
	}
	if inputx.Tktnfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Tktnfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "tktnfl",
			Value: inputx.Tktnfl_psgdtl}})
	}
	if inputx.Tktnfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Tktnfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "tktnfl",
			Value: inputx.Tktnfl_psgdtl}})
	}
	if inputx.Isitfl_psgdtl != "" {
		nowIsitfl := "F"
		if inputx.Isitfl_psgdtl == "Not flown" {
			nowIsitfl = "N"
		}
		csvFilenm = append(csvFilenm, inputx.Isitfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "isitfl",
			Value: nowIsitfl}})
	}
	if inputx.Nclear_psgdtl != "ALL" {
		var mtchor = bson.A{}
		if inputx.Nclear_psgdtl == "MNFEST" || inputx.Nclear_psgdtl == "" {
			mtchor = append(mtchor, bson.D{{Key: "mnfest", Value: "NOT CLEAR"}})
		}
		if inputx.Nclear_psgdtl == "SLSRPT" || inputx.Nclear_psgdtl == "" {
			mtchor = append(mtchor, bson.D{{Key: "slsrpt", Value: "NOT CLEAR"}})
		}
		if len(mtchor) > 0 {
			csvFilenm = append(csvFilenm, inputx.Nclear_psgdtl)
			mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: mtchor}})
		}
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
		pipeln := mongo.Pipeline{
			mtchfn,
			sortdt,
			bson.D{{Key: "$skip", Value: (max(inputx.Pagenw_psgdtl, 1) - 1) * inputx.Limitp_psgdtl}},
			bson.D{{Key: "$limit", Value: inputx.Limitp_psgdtl}},
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, pipeln)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		if strDvsion == "mnfest" {
			var slctmp = []mdlPsglst.MdlPsglstPsgdtlMnfest{}
			for rawDtaset.Next(contxt) {
				slcDtaset := mdlPsglst.MdlPsglstPsgdtlMnfest{}
				rawDtaset.Decode(&slcDtaset)
				slctmp = append(slctmp, slcDtaset)
			}
			slcobj = slctmp
		} else {
			var slctmp = []mdlPsglst.MdlPsglstPsgdtlSlsflw{}
			for rawDtaset.Next(contxt) {
				slcDtaset := mdlPsglst.MdlPsglstPsgdtlSlsflw{}
				rawDtaset.Decode(&slcDtaset)
				slctmp = append(slctmp, slcDtaset)
			}
			slcobj = slctmp
		}

	}()

	// Waiting until all go done
	wg.Wait()

	// Return final output
	c.JSON(200, gin.H{"totdta": totidx, "arrdta": slcobj})
}

// Download PNR Detail all
func FncPsglstPsgdtlDownld(c *gin.Context) {

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
	if inputx.Datefl_psgdtl != "" {
		strDatefl, _ := time.Parse("2006-01-02", inputx.Datefl_psgdtl)
		intDatefl, _ = strconv.Atoi(strDatefl.Format("060102"))
	}

	// Select db and context to do
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgdtl")
	contxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}

	// Check if data Route all is isset
	if inputx.Datefl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Datefl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
	if inputx.Airlfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Airlfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: inputx.Airlfl_psgdtl}})
	}
	if inputx.Flnbfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "flnbfl",
			Value: inputx.Flnbfl_psgdtl}})
	}
	if inputx.Depart_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Depart_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "depart",
			Value: inputx.Depart_psgdtl}})
	}
	if inputx.Routfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Routfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "routfl",
			Value: inputx.Routfl_psgdtl}})
	}
	if inputx.Pnrcde_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Pnrcde_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "pnrcde",
			Value: inputx.Pnrcde_psgdtl}})
	}
	if inputx.Tktnfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Tktnfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "tktnfl",
			Value: inputx.Tktnfl_psgdtl}})
	}
	if inputx.Tktnfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Tktnfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "tktnfl",
			Value: inputx.Tktnfl_psgdtl}})
	}
	if inputx.Isitfl_psgdtl != "" {
		nowIsitfl := "F"
		if inputx.Isitfl_psgdtl == "Not flown" {
			nowIsitfl = "N"
		}
		csvFilenm = append(csvFilenm, inputx.Isitfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "isitfl",
			Value: nowIsitfl}})
	}
	if inputx.Nclear_psgdtl != "ALL" {
		var mtchor = bson.A{}
		if inputx.Nclear_psgdtl == "MNFEST" || inputx.Nclear_psgdtl == "" {
			mtchor = append(mtchor, bson.D{{Key: "mnfest", Value: "NOT CLEAR"}})
		}
		if inputx.Nclear_psgdtl == "SLSRPT" || inputx.Nclear_psgdtl == "" {
			mtchor = append(mtchor, bson.D{{Key: "slsrpt", Value: "NOT CLEAR"}})
		}
		if inputx.Nclear_psgdtl == "" {
			csvFilenm = append(csvFilenm, "ALL_NOT_CLEAR")
		} else {
			csvFilenm = append(csvFilenm, inputx.Nclear_psgdtl+"_NOT_CLEAR")
		}
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: mtchor}})
	} else {
		csvFilenm = append(csvFilenm, "ALL")
	}

	// Final match pipeline
	var mtchfn bson.D
	if len(mtchdt) != 0 {
		mtchfn = bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: mtchdt}}}}
	} else {
		fmt.Println("mtchblnk")
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
	if inputx.Format_psgdtl == "EBTFMT" {
		writer.Write([]string{
			"CEK GROUP",
			"Isitfl",
			"Isittx",
			"Airlfl",
			"Flnbfl",
			"Datefl",
			"Depart",
			"Nmelst",
			"Nmefst",
			"Groupc",
			"Totpax",
			"Arrivl",
			"Seatpx",
			"Tktnvc",
			"Cpnbvc",
			"Clssvc",
			"Qtotbt",
			"Wtotbt",
			"Ftotbt",
			"Totexc",
			"Ptotbt",
			"Shopay",
			"Coment",
			"Comment Staff",
			"Agtdcr",
			"Name Staff",
			"Currency",
			"Total",
			"Pnrcde",
			"Pnritl",
			"Timeis",
			"Agtdie",
			"",
			"",
			"",
			"",
		})
	} else {
		writer.Write([]string{
			"mnfest",
			"slsrpt",
			"noterr",
			"source",
			"tktnfl",
			"tktnvc",
			"pnrcde",
			"pnritl",
			"curncy",
			"ntaffl",
			"ntafvc",
			"yqtxfl",
			"yqtxvc",
			"frrate",
			"frbcde",
			"qsrcrw",
			"qsrcvc",
			"frcalc",
			"ndayfl",
			"datefl",
			"datevc",
			"daterv",
			"mnthfl",
			"timefl",
			"timerv",
			"timeis",
			"timecr",
			"airlfl",
			"airlvc",
			"airtyp",
			"seatcn",
			"flnbfl",
			"flnbvc",
			"flgate",
			"bookdc",
			"bookdy",
			"depart",
			"arrivl",
			"routfl",
			"routvc",
			"routvf",
			"routac",
			"routmx",
			"routfr",
			"routfx",
			"routsg",
			"linenb",
			"ckinnb",
			"gender",
			"typepx",
			"seatpx",
			"groupc",
			"totpax",
			"segpnr",
			"segtkt",
			"psgrid",
			"tourcd",
			"staloc",
			"stanbr",
			"wrkloc",
			"hmeloc",
			"lniata",
			"emplid",
			"nmefst",
			"nmelst",
			"cpnbfl",
			"cpnbvc",
			"clssfl",
			"clssvc",
			"statvc",
			"cbinfl",
			"cbinvc",
			"agtdie",
			"agtdcr",
			"codels",
			"isitfl",
			"isittx",
			"isitir",
			"isitct",
			"isittf",
			"isitnr",
			"noteup",
			"updtby",
			"prmkey",

			// Ancillary
			"gpcdae",
			"sbcdae",
			"descae",
			"wgbgae",
			"qtbgae",
			"routae",
			"fareae",
			"currae",
			"emdnae",

			// Bagtag
			"nmbrbt",
			"qntybt",
			"wghtbt",
			"paidbt",
			"fbavbt",
			"hfbabt",
			"qtotbt",
			"wtotbt",
			"ptotbt",
			"ftotbt",
			"excsbt",
			"typebt",
			"coment",

			// Outbound
			"airlob",
			"flnbob",
			"clssob",
			"routob",
			"dateob",
			"timeob",

			// Inbound
			"airlib",
			"flnbib",
			"clssib",
			"dstrib",
			"dateib",
			"timeib",

			// Ireg
			"codeir",
			"airlir",
			"flnbir",
			"dateir",

			// Infant
			"tktnif",
			"cpnbif",
			"dateif",
			"clssif",
			"routif",
			"statif",
			"paxsif",

			// Cancel bagtag
			"airlxt",
			"dstrxt",
			"nmbrxt",
		})
	}
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
		var slcDtaset mdlPsglst.MdlPsglstPsgdtlDtbase
		rawDtaset.Decode(&slcDtaset)
		// fmtTimefl, _ := time.Parse("0601021504", strconv.Itoa(int(slcDtaset.Timefl)))
		// strTimefl := fmtTimefl.Format("02-Jan-2006 15:04")
		// fmtTimerv, _ := time.Parse("0601021504", strconv.Itoa(int(slcDtaset.Timerv)))
		// strTimerv := fmtTimerv.Format("02-Jan-2006 15:04")
		// fmtTimecr, _ := time.Parse("0601021504", strconv.Itoa(int(slcDtaset.Timecr)))
		// strTimecr := fmtTimecr.Format("02-Jan-2006 15:04")
		fmtTimeis, _ := time.Parse("0601021504", strconv.Itoa(int(slcDtaset.Timeis)))
		strTimeis := fmtTimeis.Format("02-Jan-2006 15:04")
		fmtDatefl, _ := time.Parse("060102", strconv.Itoa(int(slcDtaset.Datefl)))
		strDatefl := fmtDatefl.Format("02-Jan-2006")
		// fmtDatevc, _ := time.Parse("060102", strconv.Itoa(int(slcDtaset.Datevc)))
		// strDatevc := fmtDatevc.Format("02-Jan-2006")
		// fmtDaterv, _ := time.Parse("060102", strconv.Itoa(int(slcDtaset.Daterv)))
		// strDaterv := fmtDaterv.Format("02-Jan-2006")

		// Write to CSV
		if inputx.Format_psgdtl == "EBTFMT" {
			writer.Write([]string{
				"",
				slcDtaset.Isitfl,
				slcDtaset.Isittx,
				slcDtaset.Airlfl,
				slcDtaset.Flnbfl,
				strDatefl,
				slcDtaset.Depart,
				slcDtaset.Nmelst,
				slcDtaset.Nmefst,
				slcDtaset.Groupc,
				fmt.Sprintf("%v", slcDtaset.Totpax),
				slcDtaset.Arrivl,
				slcDtaset.Seatpx,
				slcDtaset.Tktnvc,
				fmt.Sprintf("C%02d", slcDtaset.Cpnbvc),
				slcDtaset.Clssvc,
				fmt.Sprintf("%v", slcDtaset.Qtotbt),
				fmt.Sprintf("%v", slcDtaset.Wtotbt),
				fmt.Sprintf("%v", slcDtaset.Ftotbt),
				fmt.Sprintf("%v", slcDtaset.Ftotbt-slcDtaset.Wtotbt),
				fmt.Sprintf("%v", slcDtaset.Ptotbt),
				fmt.Sprintf("%v", slcDtaset.Ptotbt+(slcDtaset.Ftotbt-slcDtaset.Wtotbt)),
				fmt.Sprintf("%v", slcDtaset.Coment),
				"",
				slcDtaset.Agtdcr,
				"",
				"",
				"",
				slcDtaset.Pnrcde,
				slcDtaset.Pnritl,
				strTimeis,
				slcDtaset.Hmeloc + slcDtaset.Agtdcr[int(math.Max(float64(len(slcDtaset.Agtdcr)-3), 0)):],
				"",
				"",
				"",
				"",
			})
		} else {
			writer.Write([]string{
				slcDtaset.Mnfest,
				slcDtaset.Slsrpt,
				slcDtaset.Noterr,
				slcDtaset.Source,
				slcDtaset.Tktnfl,
				slcDtaset.Tktnvc,
				slcDtaset.Pnrcde,
				slcDtaset.Pnritl,
				slcDtaset.Curncy,
				fmt.Sprintf("%v", slcDtaset.Ntaffl),
				fmt.Sprintf("%v", slcDtaset.Ntafvc),
				fmt.Sprintf("%v", slcDtaset.Yqtxfl),
				fmt.Sprintf("%v", slcDtaset.Yqtxvc),
				fmt.Sprintf("%v", slcDtaset.Frrate),
				slcDtaset.Frbcde,
				slcDtaset.Qsrcrw,
				fmt.Sprintf("%v", slcDtaset.Qsrcvc),
				slcDtaset.Frcalc,
				slcDtaset.Ndayfl,
				fmt.Sprintf("%v", slcDtaset.Datefl),
				fmt.Sprintf("%v", slcDtaset.Datevc),
				fmt.Sprintf("%v", slcDtaset.Daterv),
				fmt.Sprintf("%v", slcDtaset.Mnthfl),
				fmt.Sprintf("%v", slcDtaset.Timefl),
				fmt.Sprintf("%v", slcDtaset.Timerv),
				fmt.Sprintf("%v", slcDtaset.Timeis),
				fmt.Sprintf("%v", slcDtaset.Timecr),
				slcDtaset.Airlfl,
				slcDtaset.Airlvc,
				slcDtaset.Airtyp,
				slcDtaset.Seatcn,
				slcDtaset.Flnbfl,
				slcDtaset.Flnbvc,
				slcDtaset.Flgate,
				fmt.Sprintf("%v", slcDtaset.Bookdc),
				fmt.Sprintf("%v", slcDtaset.Bookdy),
				slcDtaset.Depart,
				slcDtaset.Arrivl,
				slcDtaset.Provnc,
				slcDtaset.Routfl,
				slcDtaset.Routvc,
				slcDtaset.Routvf,
				slcDtaset.Routac,
				slcDtaset.Routmx,
				slcDtaset.Routfr,
				slcDtaset.Routfx,
				slcDtaset.Routsg,
				fmt.Sprintf("%v", slcDtaset.Linenb),
				fmt.Sprintf("%v", slcDtaset.Ckinnb),
				slcDtaset.Gender,
				slcDtaset.Typepx,
				slcDtaset.Seatpx,
				slcDtaset.Groupc,
				fmt.Sprintf("%v", slcDtaset.Totpax),
				slcDtaset.Segpnr,
				slcDtaset.Segtkt,
				slcDtaset.Psgrid,
				slcDtaset.Tourcd,
				slcDtaset.Staloc,
				slcDtaset.Stanbr,
				slcDtaset.Wrkloc,
				slcDtaset.Hmeloc,
				slcDtaset.Lniata,
				slcDtaset.Emplid,
				slcDtaset.Nmefst,
				slcDtaset.Nmelst,
				fmt.Sprintf("C%02d", slcDtaset.Cpnbfl),
				fmt.Sprintf("C%02d", slcDtaset.Cpnbvc),
				slcDtaset.Clssfl,
				slcDtaset.Clssvc,
				slcDtaset.Statvc,
				slcDtaset.Cbinfl,
				slcDtaset.Cbinvc,
				slcDtaset.Agtdie,
				slcDtaset.Agtdcr,
				slcDtaset.Codels,
				slcDtaset.Isitfl,
				slcDtaset.Isittx,
				slcDtaset.Isitir,
				slcDtaset.Isitct,
				slcDtaset.Isittf,
				slcDtaset.Isitnr,
				slcDtaset.Noteup,
				slcDtaset.Updtby,
				slcDtaset.Prmkey,

				// Ancillary
				slcDtaset.Gpcdae,
				slcDtaset.Sbcdae,
				slcDtaset.Descae,
				fmt.Sprintf("%v", slcDtaset.Wgbgae),
				fmt.Sprintf("%v", slcDtaset.Qtbgae),
				slcDtaset.Routae,
				fmt.Sprintf("%v", slcDtaset.Fareae),
				slcDtaset.Currae,
				slcDtaset.Emdnae,

				// Bagtag
				slcDtaset.Nmbrbt,
				fmt.Sprintf("%v", slcDtaset.Qntybt),
				fmt.Sprintf("%v", slcDtaset.Wghtbt),
				fmt.Sprintf("%v", slcDtaset.Paidbt),
				fmt.Sprintf("%v", slcDtaset.Fbavbt),
				fmt.Sprintf("%v", slcDtaset.Hfbabt),
				fmt.Sprintf("%v", slcDtaset.Qtotbt),
				fmt.Sprintf("%v", slcDtaset.Wtotbt),
				fmt.Sprintf("%v", slcDtaset.Ptotbt),
				fmt.Sprintf("%v", slcDtaset.Ftotbt),
				fmt.Sprintf("%v", slcDtaset.Excsbt),
				slcDtaset.Typebt,
				slcDtaset.Coment,

				// Outbound
				slcDtaset.Airlob,
				slcDtaset.Flnbob,
				slcDtaset.Clssob,
				slcDtaset.Routob,
				fmt.Sprintf("%v", slcDtaset.Dateob),
				fmt.Sprintf("%v", slcDtaset.Timeob),

				// Inbound
				slcDtaset.Airlib,
				slcDtaset.Flnbib,
				slcDtaset.Clssib,
				slcDtaset.Dstrib,
				fmt.Sprintf("%v", slcDtaset.Dateib),
				fmt.Sprintf("%v", slcDtaset.Timeib),

				// Ireg
				slcDtaset.Codeir,
				slcDtaset.Airlir,
				slcDtaset.Flnbir,
				fmt.Sprintf("%v", slcDtaset.Dateir),

				// Infant
				slcDtaset.Tktnif,
				fmt.Sprintf("%v", slcDtaset.Cpnbif),
				fmt.Sprintf("%v", slcDtaset.Dateif),
				slcDtaset.Clssif,
				slcDtaset.Routif,
				slcDtaset.Statif,
				slcDtaset.Paxsif,

				// Cancel bagtag
				slcDtaset.Airlxt,
				slcDtaset.Dstrxt,
				slcDtaset.Nmbrxt,
			})
		}

		// Flush every 1000row
		countr++
		if countr%mxflus == 0 {
			writer.Flush()
		}
	}
}

// Get Response Update database from input
func FncPsglstPsgdtlUpdate(c *gin.Context) {

	// Bind JSON Body input to variable
	var inputx mdlPsglst.MdlPsglstPsgdtlDtbase
	var findne mdlPsglst.MdlPsglstPsgdtlDtbase
	x, _ := json.MarshalIndent(inputx, " ", "")
	fmt.Println(string(x))
	if err := c.BindJSON(&inputx); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
	}

	// Select database and collection
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgdtl")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get data
	err := tablex.FindOne(contxt, bson.M{"prmkey": inputx.Prmkey}).Decode(&findne)
	if err != nil {
		fmt.Println(err)
		panic("fail")
	}

	// Get from input
	if findne.Mnfest == "NOT CLEAR" && inputx.Tktnvc != "" {
		findne.Tktnvc = inputx.Tktnvc
		if findne.Tktnfl == "" {
			findne.Tktnfl = inputx.Tktnvc
		}
		fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "TKT MANUAL")
	}
	if findne.Mnfest == "NOT CLEAR" && inputx.Airlvc != "" {
		findne.Airlvc = inputx.Airlvc
		fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "AIRLINE MANUAL")
	}
	if findne.Mnfest == "NOT CLEAR" && inputx.Flnbvc != "" {
		findne.Flnbvc = inputx.Flnbvc
		fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "FLNUMBER MANUAL")
	}
	if findne.Mnfest == "NOT CLEAR" && inputx.Cpnbvc != 0 {
		findne.Cpnbvc = inputx.Cpnbvc
		fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "CPN MANUAL")
	}
	if findne.Mnfest == "NOT CLEAR" && inputx.Routvc != "" {
		findne.Routvc = inputx.Routvc
		fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "ROUTE MANUAL")
	}
	if findne.Mnfest == "NOT CLEAR" && inputx.Statvc != "" {
		findne.Statvc = inputx.Statvc
		fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "STAT MANUAL")
	}
	if findne.Mnfest == "NOT CLEAR" && inputx.Timeis != 0 {
		findne.Statvc = inputx.Statvc
		fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "STAT MANUAL")
	}
	if findne.Slsrpt == "NOT CLEAR" && inputx.Ntafvc != 0 {
		findne.Ntafvc = inputx.Ntafvc
		fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "NTA MANUAL")
	}
	if findne.Slsrpt == "NOT CLEAR" && inputx.Qsrcvc != 0 {
		findne.Qsrcvc = inputx.Qsrcvc
		fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "QSRC MANUAL")
	}
	if findne.Slsrpt == "NOT CLEAR" && inputx.Curncy != "" {
		findne.Curncy = inputx.Curncy
		fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "CURR MANUAL")
	}
	if inputx.Updtby != "" {
		fncApndix.FncApndixUpdateSlcstr(&findne.Updtby, inputx.Updtby)
	}

	// Cek data to confirm clear
	if findne.Ntafvc != 0 && findne.Curncy != "" && findne.Ntaffl != 0 {
		findne.Slsrpt = "CLEAR"
	}
	if findne.Tktnfl != "" && findne.Tktnvc != "" && findne.Pnrcde != "" &&
		findne.Timeis != 0 && findne.Routvc != "" {
		findne.Mnfest = "CLEAR"
	}

	// Push updated data
	rsupdt := fncApndix.FncApndixBulkdbSingle([]mongo.WriteModel{
		mongo.NewUpdateOneModel().
			SetFilter(bson.M{"prmkey": findne.Prmkey}).
			SetUpdate(bson.M{"$set": findne}).
			SetUpsert(true)}, "psglst_psgdtl")
	if rsupdt != nil {
		panic("Error Insert/Update to DB:" + rsupdt.Error())
	}

	// Send token to frontend
	c.JSON(200, "success")
}
