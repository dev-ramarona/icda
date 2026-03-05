package fncPsglst

import (
	fncApndix "back/apndix/function"
	mdlPsglst "back/psglst/model"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get Detail PNR from database
func FncPsglstChartsProvnc(c *gin.Context) {

	// Bind JSON Body input to variable
	var inputx mdlPsglst.MdlPsglstParamsInputx
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	var slcobj any
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("psglst_psgsmr")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Airlfl_psgsmr != "" {
		mtchdt = append(mtchdt, bson.D{{Key: "airlfl",
			Value: inputx.Airlfl_psgsmr}})
	}
	if inputx.Flnbfl_psgsmr != "" {
		mtchdt = append(mtchdt, bson.D{{Key: "flnbfl",
			Value: inputx.Flnbfl_psgsmr}})
	}
	if inputx.Depart_psgsmr != "" {
		mtchdt = append(mtchdt, bson.D{{Key: "depart",
			Value: inputx.Depart_psgsmr}})
	}
	if inputx.Routfl_psgsmr != "" {
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

	// Get All Match Data
	wg.Add(1)
	go func() {
		defer wg.Done()
		pipeln := mongo.Pipeline{
			mtchfn,
			sortdt,
			bson.D{{Key: "$group", Value: bson.D{
				{Key: "provnc", Value: "$provnc"},
				{Key: "totnta", Value: bson.D{{Key: "$sum", Value: "$totnta"}}},
			}}},
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
	c.JSON(200, gin.H{"provnc": slcobj})
}
