package fncPsglst

import (
	fncApndix "back/apndix/function"
	mdlPsglst "back/psglst/model"
	"io"
	"math"
	"net/http"
	"reflect"
	"slices"

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
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "tktnfl", Value: inputx.Tktnfl_psgdtl}},
			bson.D{{Key: "tktnvc", Value: inputx.Tktnfl_psgdtl}}}}})
	}
	if inputx.Isitfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Isitfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "isitfl",
			Value: inputx.Isitfl_psgdtl}})
	}
	if inputx.Isitir_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Isitir_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "isitir",
			Value: inputx.Isitir_psgdtl}})
	}
	if inputx.Isittx_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Isittx_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "isittx",
			Value: inputx.Isittx_psgdtl}})
	}
	if inputx.Keywrd_psgdtl != "" && !strings.Contains(inputx.Keywrd_psgdtl, "REG ALL") {
		var slcKeywrd []string
		if err := json.Unmarshal([]byte(inputx.Keywrd_psgdtl), &slcKeywrd); err == nil {
			csvFilenm = append(csvFilenm, inputx.Keywrd_psgdtl)
			mtchdt = append(mtchdt, bson.D{{Key: "provnc",
				Value: bson.D{{Key: "$in", Value: slcKeywrd}}}})
		}
	}
	if inputx.Nclear_psgdtl != "ALL" {
		var mtchor = bson.A{}
		if inputx.Nclear_psgdtl == "MNFERR" || inputx.Nclear_psgdtl == "" {
			mtchor = append(mtchor, bson.D{{Key: "mnfest", Value: "NOT CLEAR"}})
		}
		if inputx.Nclear_psgdtl == "SLSERR" || inputx.Nclear_psgdtl == "" {
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
		switch inputx.Format_psgdtl {
		case "MNFERR":
			var slctmp = []mdlPsglst.MdlPsglstPsgdtlMnferr{}
			for rawDtaset.Next(contxt) {
				slcDtaset := mdlPsglst.MdlPsglstPsgdtlMnferr{}
				rawDtaset.Decode(&slcDtaset)
				slctmp = append(slctmp, slcDtaset)
			}
			slcobj = slctmp
		case "SLSERR":
			var slctmp = []mdlPsglst.MdlPsglstPsgdtlSlserr{}
			for rawDtaset.Next(contxt) {
				slcDtaset := mdlPsglst.MdlPsglstPsgdtlSlserr{}
				rawDtaset.Decode(&slcDtaset)
				slctmp = append(slctmp, slcDtaset)
			}
			slcobj = slctmp
		case "EBTFMT":
			var slctmp = []mdlPsglst.MdlPsglstPsgdtlEbtfmt{}
			for rawDtaset.Next(contxt) {
				slcDtaset := mdlPsglst.MdlPsglstPsgdtlEbtfmt{}
				rawDtaset.Decode(&slcDtaset)
				slctmp = append(slctmp, slcDtaset)
			}
			slcobj = slctmp
		case "TKTFMT":
			var slctmp = []mdlPsglst.MdlPsglstPsgdtlTktfmt{}
			for rawDtaset.Next(contxt) {
				slcDtaset := mdlPsglst.MdlPsglstPsgdtlTktfmt{}
				rawDtaset.Decode(&slcDtaset)
				slctmp = append(slctmp, slcDtaset)
			}
			slcobj = slctmp
		default:
			var slctmp = []mdlPsglst.MdlPsglstPsgdtlDfault{}
			for rawDtaset.Next(contxt) {
				slcDtaset := mdlPsglst.MdlPsglstPsgdtlDfault{}
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
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgdtl")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}

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
		csvFilenm = append(csvFilenm, inputx.Isitfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "isitfl",
			Value: inputx.Isitfl_psgdtl}})
	}
	if inputx.Isitir_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Isitir_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "isitir",
			Value: inputx.Isitir_psgdtl}})
	}
	if inputx.Isittx_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Isittx_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "isittx",
			Value: inputx.Isittx_psgdtl}})
	}
	if inputx.Nclear_psgdtl != "ALL" {
		var mtchor = bson.A{}
		if inputx.Nclear_psgdtl == "MNFERR" || inputx.Nclear_psgdtl == "" {
			mtchor = append(mtchor, bson.D{{Key: "mnfest", Value: "NOT CLEAR"}})
		}
		if inputx.Nclear_psgdtl == "SLSERR" || inputx.Nclear_psgdtl == "" {
			mtchor = append(mtchor, bson.D{{Key: "slsrpt", Value: "NOT CLEAR"}})
		}
		if len(mtchor) > 0 {
			csvFilenm = append(csvFilenm, inputx.Nclear_psgdtl)
			mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: mtchor}})
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
	switch inputx.Format_psgdtl {
	case "MNFERR":
		writer.Write([]string{
			"Noterr",
			"Prmkey",
			"Pnrcde",
			"Seatpx",
			"Nmelst",
			"Nmefst",
			"Groupc",
			"Airlfl",
			"Flnbfl",
			"Routfl",
			"Datefl",
			"Tktnvc_target",
			"Airlvc_target",
			"Flnbvc_target",
			"Routvc_target",
			"Cpnbvc_target",
			"Statvc_target",
			"Timeis_target",
			"Remark_target",
		})
	case "SLSERR":
		writer.Write([]string{
			"Noterr",
			"Prmkey",
			"Pnrcde",
			"Airlfl",
			"Flnbfl",
			"Routfl",
			"Provnc",
			"Datefl",
			"Timecr",
			"Timeis",
			"Tktnvc",
			"Isitnr",
			"Frcalc",
			"Curncy_target",
			"Ntafvc_target",
			"Yqtxvc_target",
			"Qsrcvc_target",
		})
	case "EBTFMT":
		writer.Write([]string{
			"CEK GROUP",
			"Prmkey",
			"Mnthfl",
			"Airlfl",
			"Flnbfl",
			"Datefl",
			"Depart",
			"Arrivl",
			"Nmelst",
			"Nmefst",
			"Groupc",
			"Totpax",
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
			"Cmnstf",
			"Agtdcr",
			"Descae",
			"Currae",
			"Total",
			"domint",
			"Timeis",
			"Pnrcde",
			"Pnritl",
			"Agtdie",
			"Isitct",
			"Airtyp",
			"folllow",
			"categr",
			"Isitfl",
			"Isittx",
		})
	case "TKTFMT":
		writer.Write([]string{
			"Nmefst",
			"Nmelst",
			"Airlfl",
			"Flnbfl",
			"Datefl",
			"Depart",
			"Groupc",
			"Arrivl",
			"Seatpx",
			"Tktnvc",
			"Cpnbvc",
			"Datevc",
			"Clssvc",
			"Routvc",
			"Statvc",
			"Isittx",
			"OTHER CPN",
			"OTHER CLS",
			"OTHER ROUTE",
			"OTHER STATUS",
			"Remark",
			"Gender",
			"Routfl",
			"Isitir",
		})
	default:
		writer.Write([]string{
			"Mnfest",
			"Slsrpt",
			"Noterr",
			"Source",
			"Tktnfl",
			"Tktnvc",
			"Pnrcde",
			"Pnritl",
			"Curncy",
			"Ntaffl",
			"Ntafvc",
			"Yqtxfl",
			"Yqtxvc",
			"Frrate",
			"Frbcde",
			"Qsrcrw",
			"Qsrcvc",
			"Frcalc",
			"Ndayfl",
			"Datefl",
			"Datevc",
			"Daterv",
			"Mnthfl",
			"Timefl",
			"Timerv",
			"Timeis",
			"Timecr",
			"Airlfl",
			"Airlvc",
			"Airtyp",
			"Seatcn",
			"Flhour",
			"Flnbfl",
			"Flnbvc",
			"Flgate",
			"Bookdc",
			"Bookdy",
			"Depart",
			"Arrivl",
			"Provnc",
			"Routfl",
			"Routvc",
			"Routvf",
			"Routac",
			"Routmx",
			"Routfr",
			"Routfx",
			"Routsg",
			"Linenb",
			"Ckinnb",
			"Gender",
			"Typepx",
			"Seatpx",
			"Groupc",
			"Totpax",
			"Segpnr",
			"Segtkt",
			"Psgrid",
			"Tourcd",
			"Staloc",
			"Stanbr",
			"Wrkloc",
			"Hmeloc",
			"Lniata",
			"Emplid",
			"Nmefst",
			"Nmelst",
			"Cpnbfl",
			"Cpnbvc",
			"Clssfl",
			"Clssvc",
			"Statvc",
			"Cbinfl",
			"Cbinvc",
			"Agtdie",
			"Agtdcr",
			"Codels",
			"Isitfl",
			"Isittx",
			"Isitir",
			"Isitct",
			"Isitnr",
			"Srcfrb",
			"Srcyqf",
			"Noteup",
			"Updtby",
			"Prmkey",

			// Ancillary
			"Gpcdae",
			"Sbcdae",
			"Descae",
			"Wgbgae",
			"Qtbgae",
			"Routae",
			"Fareae",
			"Currae",
			"Emdnae",

			// Bagtag
			"Nmbrbt",
			"Qntybt",
			"Wghtbt",
			"Paidbt",
			"Fbavbt",
			"Hfbabt",
			"Qtotbt",
			"Wtotbt",
			"Ptotbt",
			"Ftotbt",
			"Excsbt",
			"Typebt",
			"Coment",

			// Outbound
			"Airlob",
			"Flnbob",
			"Clssob",
			"Routob",
			"Dateob",
			"Timeob",

			// Inbound
			"Airlib",
			"Flnbib",
			"Clssib",
			"Dstrib",
			"Dateib",
			"Timeib",

			// Ireg
			"Codeir",
			"Airlir",
			"Flnbir",
			"Dateir",

			// Infant
			"Tktnif",
			"Cpnbif",
			"Dateif",
			"Clssif",
			"Routif",
			"Statif",
			"Paxsif",

			// Cancel bagtag
			"Airlxt",
			"Dstrxt",
			"Nmbrxt",
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
	mxflus := 1000
	countr := 0
	for rawDtaset.Next(contxt) {
		var slcDtaset mdlPsglst.MdlPsglstPsgdtlDtbase
		rawDtaset.Decode(&slcDtaset)
		strTimeis := fncApndix.FncApndixFormatTimeot(int(slcDtaset.Timeis))
		strTimecr := fncApndix.FncApndixFormatTimeot(int(slcDtaset.Timecr))
		strMnthfl := fncApndix.FncApndixFormatMnthot(int(slcDtaset.Timeis))
		strDatefl := fncApndix.FncApndixFormatDateot(int(slcDtaset.Datefl))
		strDatevc := fncApndix.FncApndixFormatDateot(int(slcDtaset.Datevc))

		// Write to CSV
		switch inputx.Format_psgdtl {
		case "MNFERR":
			writer.Write([]string{
				slcDtaset.Noterr,
				slcDtaset.Prmkey,
				slcDtaset.Pnrcde,
				slcDtaset.Seatpx,
				slcDtaset.Nmelst,
				slcDtaset.Nmefst,
				slcDtaset.Groupc,
				slcDtaset.Airlfl,
				slcDtaset.Flnbfl,
				slcDtaset.Routfl,
				fmt.Sprintf("%v", slcDtaset.Datefl),
				slcDtaset.Tktnvc,
				slcDtaset.Airlvc,
				slcDtaset.Flnbvc,
				slcDtaset.Routvc,
				fmt.Sprintf("%v", slcDtaset.Cpnbvc),
				slcDtaset.Statvc,
				fmt.Sprintf("%v", slcDtaset.Timeis),
				slcDtaset.Remark,
			})
		case "SLSERR":
			writer.Write([]string{
				slcDtaset.Noterr,
				slcDtaset.Prmkey,
				slcDtaset.Pnrcde,
				slcDtaset.Airlfl,
				slcDtaset.Flnbfl,
				slcDtaset.Routfl,
				slcDtaset.Provnc,
				strDatefl,
				strTimeis,
				strTimecr,
				slcDtaset.Tktnvc,
				slcDtaset.Isitnr,
				slcDtaset.Frcalc,
				slcDtaset.Curncy,
				fmt.Sprintf("%v", slcDtaset.Ntafvc),
				fmt.Sprintf("%v", slcDtaset.Yqtxvc),
				fmt.Sprintf("%v", slcDtaset.Qsrcvc),
			})
		case "EBTFMT":
			writer.Write([]string{
				"",
				slcDtaset.Prmkey,
				strMnthfl,
				slcDtaset.Airlfl,
				slcDtaset.Flnbfl,
				strDatefl,
				slcDtaset.Depart,
				slcDtaset.Arrivl,
				slcDtaset.Nmelst,
				slcDtaset.Nmefst,
				slcDtaset.Groupc,
				fmt.Sprintf("%v", slcDtaset.Totpax),
				slcDtaset.Seatpx,
				slcDtaset.Tktnvc,
				fmt.Sprintf("C0%v", slcDtaset.Cpnbvc),
				slcDtaset.Clssvc,
				fmt.Sprintf("%v", slcDtaset.Qtotbt),
				fmt.Sprintf("%v", slcDtaset.Wtotbt),
				fmt.Sprintf("%v", slcDtaset.Ftotbt),
				fmt.Sprintf("%v", slcDtaset.Ftotbt-slcDtaset.Wtotbt),
				fmt.Sprintf("%v", slcDtaset.Ptotbt),
				fmt.Sprintf("%v", slcDtaset.Ptotbt+(slcDtaset.Ftotbt-slcDtaset.Wtotbt)),
				slcDtaset.Coment,
				"",
				"",
				slcDtaset.Descae,
				slcDtaset.Currae,
				"",
				"",
				strTimeis,
				slcDtaset.Pnrcde,
				slcDtaset.Pnritl,
				slcDtaset.Hmeloc + slcDtaset.Agtdie,
				slcDtaset.Isitct,
				slcDtaset.Airtyp,
				"",
				"",
				slcDtaset.Isitfl,
				slcDtaset.Isittx,
			})
		case "TKTFMT":
			writer.Write([]string{
				slcDtaset.Nmefst,
				slcDtaset.Nmelst,
				slcDtaset.Airlfl,
				slcDtaset.Flnbfl,
				strDatefl,
				slcDtaset.Depart,
				slcDtaset.Groupc,
				slcDtaset.Arrivl,
				slcDtaset.Seatpx,
				slcDtaset.Tktnvc,
				fmt.Sprintf("%v", slcDtaset.Cpnbvc),
				strDatevc,
				slcDtaset.Clssvc,
				slcDtaset.Routvc,
				slcDtaset.Statvc,
				slcDtaset.Isittx,
				"",
				"",
				"",
				"",
				slcDtaset.Remark,
				slcDtaset.Gender,
				slcDtaset.Routfl,
				slcDtaset.Isitir})
		default:
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
				fmt.Sprintf("%v", slcDtaset.Flhour),
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
				fmt.Sprintf("%v", slcDtaset.Cpnbfl),
				fmt.Sprintf("%v", slcDtaset.Cpnbvc),
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
				slcDtaset.Isitnr,
				slcDtaset.Srcfrb,
				slcDtaset.Srcyqf,
				slcDtaset.Noteup,
				slcDtaset.Updtby,
				slcDtaset.Prmkey,
				slcDtaset.Gpcdae,
				slcDtaset.Sbcdae,
				slcDtaset.Descae,
				fmt.Sprintf("%v", slcDtaset.Wgbgae),
				fmt.Sprintf("%v", slcDtaset.Qtbgae),
				slcDtaset.Routae,
				fmt.Sprintf("%v", slcDtaset.Fareae),
				slcDtaset.Currae,
				slcDtaset.Emdnae,
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
				slcDtaset.Airlob,
				slcDtaset.Flnbob,
				slcDtaset.Clssob,
				slcDtaset.Routob,
				fmt.Sprintf("%v", slcDtaset.Dateob),
				fmt.Sprintf("%v", slcDtaset.Timeob),
				slcDtaset.Airlib,
				slcDtaset.Flnbib,
				slcDtaset.Clssib,
				slcDtaset.Dstrib,
				fmt.Sprintf("%v", slcDtaset.Dateib),
				fmt.Sprintf("%v", slcDtaset.Timeib),
				slcDtaset.Codeir,
				slcDtaset.Airlir,
				slcDtaset.Flnbir,
				fmt.Sprintf("%v", slcDtaset.Dateir),
				slcDtaset.Tktnif,
				fmt.Sprintf("%v", slcDtaset.Cpnbif),
				fmt.Sprintf("%v", slcDtaset.Dateif),
				slcDtaset.Clssif,
				slcDtaset.Routif,
				slcDtaset.Statif,
				slcDtaset.Paxsif,
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
	var dvsion = c.Param("dvsion")
	var inputx mdlPsglst.MdlPsglstPsgdtlDtbase
	var findne mdlPsglst.MdlPsglstPsgdtlDtbase
	if err := c.BindJSON(&inputx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
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

	// Get from input manifest
	if dvsion == "mnfest" {
		if inputx.Tktnvc == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Tktnvc"})
			return
		} else if inputx.Tktnvc != "" {
			findne.Tktnvc = inputx.Tktnvc
			if findne.Tktnfl == "" {
				findne.Tktnfl = inputx.Tktnvc
			}
			fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "TKT MANUAL")
		}
		if inputx.Airlvc == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Airlvc"})
			return
		} else if inputx.Airlvc != "" {
			findne.Airlvc = inputx.Airlvc
			fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "AIRLINE MANUAL")
		}
		if inputx.Flnbvc == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Flnbvc"})
			return
		} else if inputx.Flnbvc != "" {
			findne.Flnbvc = inputx.Flnbvc
			fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "FLNUMBER MANUAL")
		}
		if inputx.Cpnbvc == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Cpnbvc"})
			return
		} else if inputx.Cpnbvc != 0 {
			findne.Cpnbvc = inputx.Cpnbvc
			fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "CPN MANUAL")
		}
		if inputx.Routvc == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Routvc"})
			return
		} else if inputx.Routvc != "" {
			findne.Routvc = inputx.Routvc
			fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "ROUTE MANUAL")
		}
		if inputx.Statvc == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Statvc"})
			return
		} else if inputx.Statvc != "" {
			findne.Statvc = inputx.Statvc
			fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "STAT MANUAL")
		}
		if inputx.Timeis == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Timeis"})
			return
		} else if inputx.Timeis != 0 {
			findne.Statvc = inputx.Statvc
			fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "STAT MANUAL")
		}
	}

	// Get from input sales report
	if dvsion == "slsrpt" {
		if inputx.Ntafvc == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Ntafvc"})
			return
		} else if inputx.Ntafvc != 0 {
			findne.Ntafvc = inputx.Ntafvc
			fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "NTA MANUAL")
		}
		if inputx.Curncy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Curncy"})
			return
		} else if inputx.Curncy != "" {
			findne.Curncy = inputx.Curncy
			fncApndix.FncApndixUpdateSlcstr(&findne.Noteup, "CURR MANUAL")
		}
	}

	// Additional
	if inputx.Updtby == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Updtby"})
		return
	} else if inputx.Updtby != "" {
		fncApndix.FncApndixUpdateSlcstr(&findne.Updtby, inputx.Updtby)
	}
	if dvsion == "mnfest" {
		findne.Mnfest = "CLEAR"
	}
	if dvsion == "slsrpt" {
		findne.Slsrpt = "CLEAR"
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

// Get Response Upload database from input
func FncPsglstPsgdtlUpload(c *gin.Context) {

	// Set header untuk file CSV
	filenm := "Response_upload_inputx.Format_psgdtl.csv"
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename="+filenm)
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// Bind JSON Body input to variable
	rawipt := c.PostForm("data")
	if rawipt == "" {
		writer.Write([]string{"Empty format input parameter"})
		return
	}
	var inputx mdlPsglst.MdlPsglstParamsInputx
	if err := json.Unmarshal([]byte(rawipt), &inputx); err != nil {
		writer.Write([]string{"Wrong type data input parameter"})
		return
	}

	// Get default header
	var slcDefhdr []string
	var anyDefhdr any
	var nowColerr = ""
	switch inputx.Format_psgdtl {
	case "MNFERR":
		anyDefhdr = mdlPsglst.MdlPsglstPsgdtlMnferr{}
		nowColerr = "mnfest"
	case "SLSERR":
		anyDefhdr = mdlPsglst.MdlPsglstPsgdtlSlserr{}
		nowColerr = "slsrpt"
	default:
		writer.Write([]string{"Wrong format input parameter"})
		return
	}
	var slcRawhdr = reflect.TypeOf(anyDefhdr)
	var mapFmthdr = map[string]reflect.Type{}
	for i := 0; i < slcRawhdr.NumField(); i++ {
		tag := slcRawhdr.Field(i).Tag.Get("json")
		slcDefhdr = append(slcDefhdr, strings.ToLower(tag))
		mapFmthdr[tag] = slcRawhdr.Field(i).Type
	}

	// Get multiple form
	formtp, err := c.MultipartForm()
	if err != nil {
		writer.Write([]string{"Empty file input"})
		return
	}
	fnlErrrsp := [][]string{}
	getFilesx := formtp.File["file"]
	mgoUpdate := []mongo.WriteModel{}
	for idx, filenw := range getFilesx {

		// buka file
		src, err := filenw.Open()
		if err != nil {
			writer.Write([]string{"Empty file input"})
			return
		}
		defer src.Close()

		// Compare header
		csvReader := csv.NewReader(src)
		slcNowhdr, err := csvReader.Read()
		for idx, val := range slcNowhdr {
			slcNowhdr[idx] = strings.ToLower(val[:int(math.Min(float64(6), float64(len(val))))])
		}
		slcNowhdr = append(slcNowhdr, "note")
		fnlErrrsp = append(fnlErrrsp, append([]string{"baris"}, slcNowhdr...))
		slcTmpnte := []string{}
		slcTmphdr := []string{"header"}
		for _, defhdr := range slcDefhdr {
			mfound := false
			for _, nowhdr := range slcNowhdr {
				mfound = strings.Contains(nowhdr, defhdr)
				if mfound {
					slcTmphdr = append(slcTmphdr, "-")
					break
				}
			}
			if !mfound {
				slcTmpnte = append(slcTmpnte, "False-Empty header "+defhdr)
			}
		}
		if len(slcTmpnte) != 0 {
			slcTmphdr = append(slcTmphdr, strings.Join(slcTmpnte, "|"))
		}
		if strings.Contains(strings.Join(slcTmphdr, "-"), "False") {
			fnlErrrsp = append(fnlErrrsp, slcTmphdr)
		}

		// read CSV per row
		if len(fnlErrrsp) <= idx+1 {
			intCountd := 1
			for {
				slcRowdta, err := csvReader.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					writer.Write([]string{"Empty row csv file input"})
					return
				}

				// Cek header and data
				if len(slcDefhdr) != len(slcRowdta) {
					writer.Write([]string{"Header and data mismatch"})
					return
				}

				// read CSV per col
				objUpdate := make(map[string]any)
				objUpdate[nowColerr] = "CLEAR"
				slcTmprsp := []string{strconv.Itoa(intCountd)}
				getPrmkey := ""
				for col, colval := range slcRowdta {
					getDefhdr := slcDefhdr[col]
					getFmtnow := mapFmthdr[slcDefhdr[col]]
					slcTmprsp = append(slcTmprsp, "-")

					// Parse and cek format data
					switch getFmtnow.Kind() {
					case reflect.String:
						objUpdate[getDefhdr] = string(colval)
						if slices.Contains([]string{"tktnvc", "tktnfl"}, getDefhdr) {
							if strings.Contains(colval, "+") {
								slcTmprsp[col+1] = "False-format ticket"
							}
						}
					case reflect.Float64:
						intColval, err := strconv.ParseFloat(colval, 64)
						if err != nil {
							slcTmprsp[col+1] = "False-format number only"
						}
						objUpdate[getDefhdr] = intColval
					case reflect.Int32, reflect.Int64, reflect.Int:
						intColval, err := strconv.Atoi(colval)
						if err != nil {
							slcTmprsp[col+1] = "False-format number only"
						}
						refColval := reflect.ValueOf(intColval).Convert(getFmtnow)
						objUpdate[getDefhdr] = refColval.Interface()
					}

					// Get primary key
					if getDefhdr == "prmkey" {
						getPrmkey = colval
					}
				}
				if strings.Contains(strings.Join(slcTmprsp, "-"), "False") {
					fnlErrrsp = append(fnlErrrsp, slcTmprsp)
				}
				intCountd++

				// Push to mongomodel
				mgoUpdate = append(mgoUpdate, mongo.NewUpdateOneModel().
					SetFilter(bson.M{"prmkey": getPrmkey}).
					SetUpdate(bson.M{"$set": objUpdate}).SetUpsert(true))
			}
		}
	}

	// Cek error respon available or not
	if len(fnlErrrsp) > len(getFilesx) {
		for _, slcrsp := range fnlErrrsp {
			writer.Write(slcrsp)
		}
	} else {
		// Push last mongomodel
		fncApndix.FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
			"psglst_psgdtl": &mgoUpdate,
		}, 0)
		writer.Write([]string{"Success"})
	}
}
