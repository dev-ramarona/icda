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
	var slcobj any
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgdtl")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var keywrd = fncApndix.FncApndixKeywrdMapobj("provnc")
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Datefl_psgdtl != "" {
		csvFilenm = append(csvFilenm, strconv.Itoa(intDatefl))
		mtchdt = append(mtchdt, bson.D{{Key: "datefl",
			Value: intDatefl}})
	}
	if inputx.Airlfl_psgdtl != "" {
		slcAirlfl := strings.Split(inputx.Airlfl_psgdtl, "-")
		csvFilenm = append(csvFilenm, inputx.Airlfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: bson.M{"$in": slcAirlfl}}})
	}
	if inputx.Flnbfl_psgdtl != "" {
		csvFilenm = append(csvFilenm, inputx.Flnbfl_psgdtl)
		mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "flnbfl", Value: inputx.Flnbfl_psgdtl}},
			bson.D{{Key: "flnbjn", Value: inputx.Flnbfl_psgdtl}}}}})
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
	if inputx.Keywrd_psgdtl != "" {
		var slcKeywrd []string
		if err := json.Unmarshal([]byte(inputx.Keywrd_psgdtl), &slcKeywrd); err == nil {
			for _, key := range slcKeywrd {
				if val, ist := keywrd[key]; ist && val != "ALL" {
					csvFilenm = append(csvFilenm, val)
					mtchdt = append(mtchdt, bson.D{{Key: "provnc",
						Value: val}})
				}
			}
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
		case "FMWCHR":
			var slctmp = []mdlPsglst.MdlPsglstPsgdtlFmwchr{}
			for rawDtaset.Next(contxt) {
				slcDtaset := mdlPsglst.MdlPsglstPsgdtlFmwchr{}
				rawDtaset.Decode(&slcDtaset)
				slctmp = append(slctmp, slcDtaset)
			}
			slcobj = slctmp
		case "FMTINF":
			var slctmp = []mdlPsglst.MdlPsglstPsgdtlFmtinf{}
			for rawDtaset.Next(contxt) {
				slcDtaset := mdlPsglst.MdlPsglstPsgdtlFmtinf{}
				rawDtaset.Decode(&slcDtaset)
				slctmp = append(slctmp, slcDtaset)
			}
			slcobj = slctmp
		case "FMTHAI":
			var slctmp = []mdlPsglst.MdlPsglstPsgdtlFmthai{}
			for rawDtaset.Next(contxt) {
				slcDtaset := mdlPsglst.MdlPsglstPsgdtlFmthai{}
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
	if strings.Contains(inputx.Keywrd_psgdtl, "dwlsls") {
		tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgdtl")
		contxt, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		// Pipeline get the data logic match
		var mtchdt = bson.A{}
		var keywrd = fncApndix.FncApndixKeywrdMapobj("provnc")
		var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}

		// Check if data Route all is isset
		if inputx.Datefl_psgdtl != "" {
			csvFilenm = append(csvFilenm, strconv.Itoa(intDatefl))
			mtchdt = append(mtchdt, bson.D{{Key: "datefl",
				Value: intDatefl}})
		}
		if inputx.Airlfl_psgdtl != "" {
			slcAirlfl := strings.Split(inputx.Airlfl_psgdtl, "-")
			csvFilenm = append(csvFilenm, inputx.Airlfl_psgdtl)
			mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
				Value: bson.M{"$in": slcAirlfl}}})
		}
		if inputx.Flnbfl_psgdtl != "" {
			csvFilenm = append(csvFilenm, inputx.Flnbfl_psgdtl)
			mtchdt = append(mtchdt, bson.D{{Key: "$or", Value: bson.A{
				bson.D{{Key: "flnbfl", Value: inputx.Flnbfl_psgdtl}},
				bson.D{{Key: "flnbjn", Value: inputx.Flnbfl_psgdtl}}}}})
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
		if inputx.Keywrd_psgdtl != "" {
			var slcKeywrd []string
			if err := json.Unmarshal([]byte(inputx.Keywrd_psgdtl), &slcKeywrd); err == nil {
				for _, key := range slcKeywrd {
					if val, ist := keywrd[key]; ist && val != "ALL" {
						csvFilenm = append(csvFilenm, val)
						mtchdt = append(mtchdt, bson.D{{Key: "provnc",
							Value: val}})
					}
				}
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
				"Rmkusr_target",
			})
		case "SLSERR":
			writer.Write([]string{
				"Noterr",
				"Prmkey",
				"Pnrcde",
				"Airlfl",
				"Flnbfl",
				"Routfl",
				"Routvc",
				"Routac",
				"Provnc",
				"Datefl",
				"Timecr",
				"Timeis",
				"Tktnvc",
				"Clssfl",
				"Clssvc",
				"Isitnr",
				"Isitfl",
				"Isitct",
				"Frcalc",
				"Ntacrr_target",
				"Ntacrt_target",
				"Ntafvc_target",
				"Taxcrr_target",
				"Taxcrt_target",
				"Yqtxvc_target",
				"Yrtxvc_target",
				"Qsrcvc_target",
			})
		case "FMTHAI":
			writer.Write([]string{
				"Noterr",
				"Prmkey",
				"Pnrcde",
				"Airlfl",
				"Flnbfl",
				"Routfl",
				"Routvc",
				"Routac",
				"Provnc",
				"Datefl",
				"Timecr",
				"Timeis",
				"Tktnvc",
				"Nmefst",
				"Nmelst",
				"Clssfl",
				"Clssvc",
				"Isitnr",
				"Isitfl",
				"Isitct",
				"Frcalc",
				"Ntacrr_target",
				"Ntacrt_target",
				"Ntafvc_target",
				"Taxcrr_target",
				"Taxcrt_target",
				"Yqtxvc_target",
				"Yrtxvc_target",
				"Qsrcvc_target",
				"Aptxvc",
				"Rwtxvc",
			})
		case "FMWCHR":
			writer.Write([]string{
				"Airline Flown",
				"Flight number flown",
				"Date flown",
				"Departure",
				"Arrival",
				"Name first passenger",
				"Name last passenger",
				"Group code",
				"Total pax",
				"Seat number passenger",
				"Ticket number flown",
				"Coupon number VCR",
				"Class RBD VCR",
				"PNR code",
				"PNR interline",
				"Comment",
				"EMD number acnillary",
				"Group code acnillary",
				"Description acnillary",
				"Route acnillary",
				"Currency acnillary",
				"Fare acnillary",
				"Is it transit?",
				"Code list from passenger list",
				"STATUS",
			})
		case "FMTINF":
			writer.Write([]string{
				"Airline Flown",
				"Flight number flown",
				"Date flown",
				"Departure",
				"Arrival",
				"Ticket infant",
				"Coupon infant",
				"Date infant",
				"Class infant",
				"Route infant",
				"Status infant",
				"Name first passenger",
				"Name last passenger",
				"Is it transit?",
				"CPN",
				"CLS",
				"RUTE",
				"STATUS",
				"KETERANGAN",
				"Passenger infant",
				"PNR code",
				"Code list from passenger list",
			})
		case "EBTFMT":
			writer.Write([]string{
				"CEK GROUP",
				"Primary key",
				"Month flown",
				"Airline Flown",
				"Flight number flown",
				"Date flown",
				"Departure",
				"Arrival",
				"Name last passenger",
				"Name first passenger",
				"Group code",
				"Total pax",
				"Seat number passenger",
				"Ticket number VCR",
				"Coupon number VCR",
				"Class RBD VCR",
				"Qunatity total baggage",
				"Weight total baggage",
				"Highest FBA total baggage",
				"Totexc",
				"Paid total baggage",
				"Shopay",
				"Comment",
				"Cmnstf",
				"Agentdie create PNR",
				"Description acnillary",
				"Currency acnillary",
				"Total",
				"domint",
				"Time issued",
				"PNR code",
				"PNR interline",
				"Agentdie ticket VCR",
				"Is it charter?",
				"Aircraft type",
				"folllow",
				"categr",
				"Is it flown?",
				"Is it transit?",
			})
		case "TKTFMT":
			writer.Write([]string{
				"Name first passenger",
				"Name last passenger",
				"Airline Flown",
				"Flight number flown",
				"Date flown",
				"Departure",
				"Group code",
				"Arrival",
				"Seat number passenger",
				"Ticket number VCR",
				"Coupon number VCR",
				"Date VCR",
				"Class RBD VCR",
				"Route VCR",
				"Status VCR",
				"Is it transit?",
				"OTHER CPN",
				"OTHER CLS",
				"OTHER ROUTE",
				"OTHER STATUS",
				"Remark user",
				"Gender passenger",
				"Route flown",
				"Is it iregularity?",
			})
		default:
			writer.Write([]string{
				"Manifest errror indicator",
				"Sales error indicator",
				"Note error",
				"Source Data passenger",
				"Ticket number flown",
				"Ticket number VCR",
				"Ticket number exchange",
				"PNR code",
				"PNR interline",
				"NTA currency",
				"NTA currency rate",
				"NTA flown",
				"NTA  VCR",
				"Tax currency",
				"Tax currency rate",
				"YQ flown",
				"YQ VCR",
				"YR flown",
				"YR VCR",
				"Airport tax VCR",
				"Raw all tax VCR",
				"Fare rate proration",
				"Fare base code",
				"Q fare raw text",
				"Q fare number format",
				"Fare calculation",
				"Day in a week",
				"Date flown",
				"Date VCR",
				"Date arrival",
				"Month flown",
				"Time flown",
				"Time arrival",
				"Time issued",
				"Time PNR Create",
				"Airline Flown",
				"Airline VCR",
				"Airline fare calculation",
				"Aircraft type",
				"Seat config",
				"Flight hour",
				"Flight number flown",
				"Flight number join",
				"Flight number VCR",
				"flight gate",
				"Booked business",
				"Booked economy",
				"Departure",
				"Arrival",
				"Province",
				"Route flown",
				"Route VCR",
				"Route VCR all coupon",
				"Route actual full",
				"Route maximal",
				"Route fare calculation",
				"Route fare calculation not use in coupon",
				"Route segment PNR",
				"Line number passenger list",
				"Checkin number sequence",
				"Gender passenger",
				"Type passenger",
				"Seat number passenger",
				"Group code",
				"Total pax",
				"Segment PNR",
				"Segment ticket VCR",
				"Passenger ID",
				"Tour code",
				"Station location",
				"Station number",
				"Work location",
				"Home location",
				"LNIATA",
				"Employee ID",
				"Name first passenger",
				"Name last passenger",
				"Name full passenger",
				"Coupon number flown",
				"Coupon number VCR",
				"Class RBD flown",
				"Class RBD VCR",
				"Status VCR",
				"Cabin flown",
				"Cabin VCR",
				"Agentdie ticket VCR",
				"Agentdie create PNR",
				"Code list from passenger list",
				"Is it flown?",
				"Is it infant?",
				"Is it WCHR?",
				"Is it transit?",
				"Is it iregularity?",
				"Is it charter?",
				"Is it non revenue?",
				"Is it invol?",
				"Source Fare",
				"Source YQ tax",
				"Note update",
				"Updated by",
				"Primary key",
				"Remark user",
				"Remark VCR",
				// Ancillary
				"Group code acnillary",
				"Sub code acnillary",
				"Description acnillary",
				"Weight baggage acnillary",
				"Quantity baggage acnillary",
				"Route acnillary",
				"Fare acnillary",
				"Currency acnillary",
				"EMD number acnillary",

				// Bagtag
				"Number baggage",
				"Quantity baggage",
				"Weight baggage",
				"Paid baggage",
				"FBA VCR baggage",
				"Highest FBA from rule",
				"Qunatity total baggage",
				"Weight total baggage",
				"Paid total baggage",
				"Highest FBA total baggage",
				"Type weight baggage",
				"Comment",

				// Outbound
				"Airline outbound",
				"Flight number outbound",
				"Class outbound",
				"Route outbound",
				"Date outbound",
				"Time outbound",

				// Inbound
				"Airline inbound",
				"Flight number inbound",
				"Class inbound",
				"District inbound",
				"Date inbound",
				"Time inbound",

				// Ireg
				"Code iregulerity",
				"Airline iregulerity",
				"Flight number iregulerity",
				"Date iregulerity",

				// Infant
				"Ticket infant",
				"Coupon infant",
				"Date infant",
				"Class infant",
				"Route infant",
				"Status infant",
				"Passenger infant",

				// Cancel bagtag
				"Airline cancel bagtag",
				"District cancel bagtag",
				"Number cancel bagtag",
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
			strTimefl := fncApndix.FncApndixFormatTimeot(int(slcDtaset.Timefl))
			strTimerv := fncApndix.FncApndixFormatTimeot(int(slcDtaset.Timerv))
			strMnthfl := fncApndix.FncApndixFormatMnthot(int(slcDtaset.Mnthfl))
			strDatefl := fncApndix.FncApndixFormatDateot(int(slcDtaset.Datefl))
			strDatevc := fncApndix.FncApndixFormatDateot(int(slcDtaset.Datevc))
			strDaterv := fncApndix.FncApndixFormatDateot(int(slcDtaset.Daterv))
			strDateif := fncApndix.FncApndixFormatDateot(int(slcDtaset.Dateif))

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
					slcDtaset.Rmkusr,
				})
			case "SLSERR":
				writer.Write([]string{
					slcDtaset.Noterr,
					slcDtaset.Prmkey,
					slcDtaset.Pnrcde,
					slcDtaset.Airlfl,
					slcDtaset.Flnbfl,
					slcDtaset.Routfl,
					slcDtaset.Routvc,
					slcDtaset.Routac,
					slcDtaset.Provnc,
					fmt.Sprintf("%v", slcDtaset.Datefl),
					fmt.Sprintf("%v", slcDtaset.Timeis),
					fmt.Sprintf("%v", slcDtaset.Timecr),
					slcDtaset.Tktnvc,
					slcDtaset.Clssfl,
					slcDtaset.Clssvc,
					slcDtaset.Isitnr,
					slcDtaset.Isitfl,
					slcDtaset.Isitct,
					slcDtaset.Frcalc,
					slcDtaset.Ntacrr,
					fmt.Sprintf("%v", slcDtaset.Ntacrt),
					fmt.Sprintf("%v", slcDtaset.Ntafvc),
					slcDtaset.Taxcrr,
					fmt.Sprintf("%v", slcDtaset.Taxcrt),
					fmt.Sprintf("%v", slcDtaset.Yqtxvc),
					fmt.Sprintf("%v", slcDtaset.Yrtxvc),
					fmt.Sprintf("%v", slcDtaset.Qsrcvc),
				})
			case "FMTINF":
				writer.Write([]string{
					slcDtaset.Airlfl,
					slcDtaset.Flnbfl,
					strDatefl,
					slcDtaset.Depart,
					slcDtaset.Arrivl,
					slcDtaset.Tktnif,
					fmt.Sprintf("%v", slcDtaset.Cpnbif),
					strDateif,
					slcDtaset.Clssif,
					slcDtaset.Routif,
					slcDtaset.Statif,
					slcDtaset.Nmefst,
					slcDtaset.Nmelst,
					slcDtaset.Isittx,
					fmt.Sprintf("%v", slcDtaset.Cpnbvc),
					slcDtaset.Clssvc,
					slcDtaset.Routvc,
					slcDtaset.Statvc,
					"",
					slcDtaset.Paxsif,
					slcDtaset.Pnrcde,
					slcDtaset.Codels,
				})
			case "FMWCHR":
				writer.Write([]string{
					slcDtaset.Airlfl,
					slcDtaset.Flnbfl,
					fmt.Sprintf("%v", slcDtaset.Datefl),
					slcDtaset.Depart,
					slcDtaset.Arrivl,
					slcDtaset.Nmefst,
					slcDtaset.Nmelst,
					slcDtaset.Groupc,
					fmt.Sprintf("%v", slcDtaset.Totpax),
					slcDtaset.Seatpx,
					slcDtaset.Tktnfl,
					fmt.Sprintf("%v", slcDtaset.Cpnbvc),
					slcDtaset.Clssvc,
					slcDtaset.Pnrcde,
					slcDtaset.Pnritl,
					slcDtaset.Coment,
					slcDtaset.Emdnae,
					slcDtaset.Gpcdae,
					slcDtaset.Descae,
					slcDtaset.Routae,
					slcDtaset.Currae,
					fmt.Sprintf("%v", slcDtaset.Fareae),
					slcDtaset.Isittx,
					slcDtaset.Codels,
					"",
				})
			case "FMTHAI":
				writer.Write([]string{
					slcDtaset.Noterr,
					slcDtaset.Prmkey,
					slcDtaset.Pnrcde,
					slcDtaset.Airlfl,
					slcDtaset.Flnbfl,
					slcDtaset.Routfl,
					slcDtaset.Routvc,
					slcDtaset.Routac,
					slcDtaset.Provnc,
					fmt.Sprintf("%v", slcDtaset.Datefl),
					fmt.Sprintf("%v", slcDtaset.Timeis),
					fmt.Sprintf("%v", slcDtaset.Timecr),
					slcDtaset.Tktnvc,
					slcDtaset.Nmefst,
					slcDtaset.Nmelst,
					slcDtaset.Clssfl,
					slcDtaset.Clssvc,
					slcDtaset.Isitnr,
					slcDtaset.Isitfl,
					slcDtaset.Isitct,
					slcDtaset.Frcalc,
					slcDtaset.Ntacrr,
					fmt.Sprintf("%v", slcDtaset.Ntacrt),
					fmt.Sprintf("%v", slcDtaset.Ntafvc),
					slcDtaset.Taxcrr,
					fmt.Sprintf("%v", slcDtaset.Taxcrt),
					fmt.Sprintf("%v", slcDtaset.Yqtxvc),
					fmt.Sprintf("%v", slcDtaset.Yrtxvc),
					fmt.Sprintf("%v", slcDtaset.Qsrcvc),
					fmt.Sprintf("%v", slcDtaset.Aptxvc),
					slcDtaset.Rwtxvc,
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
					slcDtaset.Rmkusr,
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
					slcDtaset.Tktnxc,
					slcDtaset.Pnrcde,
					slcDtaset.Pnritl,
					slcDtaset.Ntacrr,
					fmt.Sprintf("%v", slcDtaset.Ntacrt),
					fmt.Sprintf("%v", slcDtaset.Ntaffl),
					fmt.Sprintf("%v", slcDtaset.Ntafvc),
					slcDtaset.Taxcrr,
					fmt.Sprintf("%v", slcDtaset.Taxcrt),
					fmt.Sprintf("%v", slcDtaset.Yqtxfl),
					fmt.Sprintf("%v", slcDtaset.Yqtxvc),
					fmt.Sprintf("%v", slcDtaset.Yrtxfl),
					fmt.Sprintf("%v", slcDtaset.Yrtxvc),
					fmt.Sprintf("%v", slcDtaset.Aptxvc),
					slcDtaset.Rwtxvc,
					fmt.Sprintf("%v", slcDtaset.Frrate),
					slcDtaset.Frbcde,
					slcDtaset.Qsrcrw,
					fmt.Sprintf("%v", slcDtaset.Qsrcvc),
					slcDtaset.Frcalc,
					slcDtaset.Ndayfl,
					strDatefl,
					strDatevc,
					strDaterv,
					strMnthfl,
					strTimefl,
					strTimerv,
					strTimeis,
					strTimecr,
					slcDtaset.Airlfl,
					slcDtaset.Airlvc,
					slcDtaset.Airlfr,
					slcDtaset.Airtyp,
					slcDtaset.Seatcn,
					fmt.Sprintf("%v", slcDtaset.Flhour),
					slcDtaset.Flnbfl,
					slcDtaset.Flnbjn,
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
					slcDtaset.Nmemax,
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
					slcDtaset.Isitif,
					slcDtaset.Isitwc,
					slcDtaset.Isittx,
					slcDtaset.Isitir,
					slcDtaset.Isitct,
					slcDtaset.Isitnr,
					slcDtaset.Isitiv,
					slcDtaset.Srcfrb,
					slcDtaset.Srcyqf,
					slcDtaset.Noteup,
					slcDtaset.Updtby,
					slcDtaset.Prmkey,
					slcDtaset.Rmkusr,
					slcDtaset.Rmkvcr,
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
		if inputx.Ntacrr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input Ntacrr"})
			return
		} else if inputx.Ntacrr != "" {
			findne.Ntacrr = inputx.Ntacrr
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
