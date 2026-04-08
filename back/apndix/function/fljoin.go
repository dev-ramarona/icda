package fncApndix

import (
	mdlApndix "back/apndix/model"
	mdlPsglst "back/psglst/model"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get Response Upload database from input
func FncApndixFljoinUpload(c *gin.Context) {

	// Set header untuk file CSV
	filenm := "Response_upload_flight_join.csv"
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

	// Get multiple form
	formtp, err := c.MultipartForm()
	if err != nil {
		writer.Write([]string{"Empty file input"})
		return
	}
	fnlErrrsp := [][]string{}
	getFilesx := formtp.File["file"]
	for idx, filenw := range getFilesx {

		// buka file
		src, err := filenw.Open()
		if err != nil {
			writer.Write([]string{"Empty file input"})
			return
		}
		defer src.Close()

		// read CSV per row
		csvReader := csv.NewReader(src)
		if len(fnlErrrsp) <= idx+1 {
			intCountd := 1
			mgoFljoin := []mongo.WriteModel{}
			mgoPsgsmr := []mongo.WriteModel{}
			mgoPsgdtl := []mongo.WriteModel{}
			fnlErrrsp := [][8]string{}
			nowFljoin := ""
			for {
				slcRowdta, err := csvReader.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					writer.Write([]string{"Empty row csv file input"})
					return
				}

				// Variable scope
				tmpFljoin := mdlApndix.MdlApndixFljoinDtbase{}
				tmpErrrsp := [8]string{strconv.Itoa(intCountd), "", "", "", "", "", "", ""}

				// Get arline
				getAirlfl := strings.TrimSpace(slcRowdta[0])
				if getAirlfl == "" {
					nowFljoin = ""
					continue
				}
				if len(getAirlfl) != 2 {
					tmpErrrsp[1] = "Invalid airline code"
				}
				tmpFljoin.Airlfl = getAirlfl

				// Get flight number
				getFlnbfl := strings.TrimSpace(slcRowdta[1])
				if getFlnbfl == "" {
					nowFljoin = ""
					continue
				}
				intFlnbfl, err := strconv.Atoi(getFlnbfl)
				if err != nil {
					tmpErrrsp[2] = "Invalid flight number"
				}
				tmpFljoin.Flnbfl = strconv.Itoa(intFlnbfl)

				// Get departure
				getDepart := strings.TrimSpace(slcRowdta[2])
				if getDepart == "" {
					nowFljoin = ""
					continue
				}
				if len(getDepart) != 3 {
					tmpErrrsp[3] = "Invalid departure"
				}
				tmpFljoin.Depart = getDepart

				// Get arrival
				getArrivl := strings.TrimSpace(slcRowdta[3])
				if len(getArrivl) != 3 {
					tmpErrrsp[4] = "Invalid arrival"
				}
				tmpFljoin.Arrivl = getArrivl

				// Get status
				getStatus := strings.TrimSpace(slcRowdta[4])
				if getStatus == "" {
					nowFljoin = ""
					continue
				}
				if getStatus == "ORG" {
					nowFljoin = tmpFljoin.Flnbfl
				}

				// Get Total pax
				getTotpax := strings.TrimSpace(slcRowdta[5])
				if getTotpax == "" {
					nowFljoin = ""
					continue
				}
				intTotpax, err := strconv.Atoi(getTotpax)
				if err != nil {
					tmpErrrsp[6] = "Invalid total pax"
				}
				tmpFljoin.Totpax = int32(intTotpax)

				// Get dateflown
				getDatefl := strings.TrimSpace(slcRowdta[6])
				if getDatefl == "" {
					continue
				}
				intDatefl, err := strconv.Atoi(getDatefl)
				if err != nil {
					tmpErrrsp[7] = "Invalid date flown"
				}
				tmpFljoin.Datefl = FncApndixFormatDatexl(intDatefl)

				// Counter and push error respon if exist
				intCountd++
				if len(tmpErrrsp) > 1 {
					fnlErrrsp = append(fnlErrrsp, tmpErrrsp)
				}

				// Push to mongomodel
				if nowFljoin != "" {
					getPrmkey := tmpFljoin.Airlfl + tmpFljoin.Flnbfl + tmpFljoin.Depart + strconv.Itoa(int(tmpFljoin.Datefl))
					tmpFljoin.Flnbjn = nowFljoin
					mgoFljoin = append(mgoFljoin, mongo.NewUpdateOneModel().
						SetFilter(bson.M{"prmkey": getPrmkey}).
						SetUpdate(bson.M{"$set": tmpFljoin}).SetUpsert(true))
					FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
						"apndix_fljoin": &mgoFljoin}, 200)

					// Update psgsmr
					mgoPsgsmr = append(mgoPsgsmr, mongo.NewUpdateOneModel().
						SetFilter(bson.M{"prmkey": getPrmkey}).
						SetUpdate(bson.M{"$set": bson.M{"flnbjn": nowFljoin}}))
					FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
						"psglst_psgsmr": &mgoPsgsmr}, 200)

					// Update psgdtl
					mgoPsgdtl = append(mgoPsgdtl, mongo.NewUpdateOneModel().
						SetFilter(bson.M{
							"airlfl": tmpFljoin.Airlfl,
							"flnbfl": tmpFljoin.Flnbfl,
							"depart": tmpFljoin.Depart,
							"datefl": tmpFljoin.Datefl}).
						SetUpdate(bson.M{"$set": bson.M{"flnbjn": nowFljoin}}))
					FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
						"psglst_psgdtl": &mgoPsgdtl}, 200)
				}
			}

			// Push last mongomodel
			FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
				"apndix_fljoin": &mgoFljoin}, 0)
			FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
				"psglst_psgsmr": &mgoPsgsmr}, 0)
			FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
				"psglst_psgdtl": &mgoPsgdtl}, 0)
		}
	}

	// Cek error respon available or not
	if len(fnlErrrsp) > 0 {
		for _, slcrsp := range fnlErrrsp {
			writer.Write(slcrsp)
		}
	} else {
		writer.Write([]string{"Success"})
	}
}

// Get object slice
func FncApndixFljoinGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlApndix.MdlApdnixParamsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	var totidx = 0
	var slcobj any
	tablex := Client.Database(Dbases).Collection("apndix_fljoin")
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
		var slctmp = []mdlApndix.MdlApndixFljoinFrntnd{}
		for rawDtaset.Next(contxt) {
			slcDtaset := mdlApndix.MdlApndixFljoinFrntnd{}
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
func FncApndixFljoinDownld(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("0601021504")}
	var inputx mdlApndix.MdlApdnixParamsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	tablex := Client.Database(Dbases).Collection("apndix_fljoin")
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
	c.Header("Content-Disposition", "attachment; filename=apndix_fljoin_"+fnlFilenm+".csv")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")

	// Streaming file CSV ke client
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()
	writer.Write([]string{
		"Prmkey",
		"Airlfl",
		"Flnbfl",
		"Flnbjn",
		"Depart",
		"Arrivl",
		"Statjn",
		"Totpax",
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
		var slcDtaset mdlApndix.MdlApndixFljoinDtbase
		rawDtaset.Decode(&slcDtaset)
		strDatefl := FncApndixFormatDateot(int(slcDtaset.Datefl))

		// Write to CSV
		writer.Write([]string{
			slcDtaset.Prmkey,
			slcDtaset.Airlfl,
			slcDtaset.Flnbfl,
			slcDtaset.Flnbjn,
			slcDtaset.Depart,
			slcDtaset.Arrivl,
			slcDtaset.Statjn,
			fmt.Sprintf("%v", slcDtaset.Totpax),
			strDatefl,
		})

		// Flush every 1000row
		countr++
		if countr%mxflus == 0 {
			writer.Flush()
		}
	}
}
