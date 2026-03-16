package fncApndix

import (
	mdlApndix "back/apndix/model"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FncApndixApplstGetall(c *gin.Context) {

	// Select database and collection
	tablex := Client.Database(Dbases).Collection("apndix_applst")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	ptipns := options.Find().SetSort(bson.D{{Key: "pagenb", Value: 1}})
	datarw, err := tablex.Find(contxt, bson.M{}, ptipns)
	if err != nil {
		panic("fail")
	}
	defer datarw.Close(contxt)

	// Append to slice
	var slices = []string{}
	for datarw.Next(contxt) {
		var object mdlApndix.MdlApndixApplstDtbase
		if err := datarw.Decode(&object); err == nil {
			slices = append(slices, object.Apndix)
		}
	}

	// Send token to frontend
	c.JSON(200, slices)
}
