package main

import (
	fncAllusr "back/allusr/function"
	fncApndix "back/apndix/function"
	fncPsglst "back/psglst/function"
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Initialize MongoDB connection
	contxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	portdb := options.Client().ApplyURI(fncApndix.Urlmgo)
	fncApndix.Client, _ = mongo.Connect(contxt, portdb)

	// Framework Gin
	r := gin.Default()

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: fncApndix.Ipalow,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type",
			"Authorization", "Cookie", "Content-Disposition"},
		AllowCredentials: true,
	}))

	// Handle global
	r.GET("/global/status", fncApndix.FncApndixStatusPrcess)
	r.GET("/global/acpedt/:dvsion", fncApndix.FncApndixAcpedtGetall)

	// Handle web link API all user
	r.GET("/allusr/applst", fncAllusr.FncAllusrApplstGetall)
	r.POST("/allusr/loginx", fncAllusr.FncAllusrLoginxHandle)
	r.GET("/allusr/tokenx", fncAllusr.FncAllusrTokenxHandle)
	r.GET("/allusr/logout", fncAllusr.FncAllusrLogoutHandle)
	r.POST("/allusr/regist", fncAllusr.FncAllusrRegistHandle)
	r.POST("/allusr/getall", fncAllusr.FncAllusrUsrlstGetall)
	r.GET("/allusr/delete/:usrnme", fncAllusr.FncAllusrDeleteHandle)

	// Handle web link API Passangger list
	r.POST("/psglst/prcess", fncPsglst.FncPsglstPrcessMainpg)
	r.POST("/psglst/psgdtl/getall/downld", fncPsglst.FncPsglstPsgdtlDownld)
	r.POST("/psglst/psgdtl/getall/:dvsion", fncPsglst.FncPsglstPsgdtlGetall)
	r.POST("/psglst/psgsmr/getall", fncPsglst.FncPsglstPsgsmrGetall)
	r.POST("/psglst/errlog/getall", fncPsglst.FncPsglstErrlogGetall)
	r.GET("/psglst/actlog/getall", fncPsglst.FncPsglstActlogGetall)
	r.POST("/psglst/psgdtl/update", fncPsglst.FncPsglstPsgdtlUpdate)

	// Run server
	r.Run("0.0.0.0:" + fncApndix.Ptgolg)
}
