package fncApndix

import (
	mdlApndix "back/apndix/model"
	mdlPsglst "back/psglst/model"
	"encoding/csv"
	"encoding/json"
	"io"
	"strconv"
	"strings"

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
				tmpFljoin := mdlApndix.MdlApndixFljoinUpdate{}
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
