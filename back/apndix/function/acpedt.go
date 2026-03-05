package fncApndix

import (
	mdlApndix "back/apndix/model"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Get data can accepted edit
func FncApndixAcpedtGetall(c *gin.Context) {

	// Select database and collection
	dvsion := c.Param("dvsion")
	tablex := Client.Database(Dbases).Collection("apndix_acpedt")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{"dvsion": dvsion})
	if err != nil {
		panic("fail")
	}
	defer datarw.Close(contxt)

	// Append to slice
	var slices = []mdlApndix.MdlApndixAcpedtDtbase{}
	for datarw.Next(contxt) {
		var object mdlApndix.MdlApndixAcpedtDtbase
		if err := datarw.Decode(&object); err == nil {
			slices = append(slices, object)
		}
	}

	// Send token to frontend
	c.JSON(200, slices)
}
