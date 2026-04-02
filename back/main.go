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

	// Handle apndix
	r.GET("/apndix/acpedt/:dvsion", fncApndix.FncApndixAcpedtGetall)
	r.GET("/apndix/applst/getall", fncApndix.FncApndixApplstGetall)
	r.POST("/apndix/provnc/getall", fncApndix.FncApndixProvncGetall)
	r.POST("/apndix/provnc/update", fncApndix.FncApndixProvncUpdate)
	r.POST("/apndix/provnc/downld", fncApndix.FncApndixProvncDownld)
	r.POST("/apndix/flhour/getall", fncApndix.FncApndixFlhourGetall)
	r.POST("/apndix/flhour/update", fncApndix.FncApndixFlhourUpdate)
	r.POST("/apndix/flhour/downld", fncApndix.FncApndixFlhourDownld)
	r.POST("/apndix/frbase/getall", fncApndix.FncApndixFrbaseGetall)
	r.POST("/apndix/frbase/update", fncApndix.FncApndixFrbaseUpdate)
	r.POST("/apndix/frbase/downld", fncApndix.FncApndixFrbaseDownld)
	r.POST("/apndix/frtaxs/getall", fncApndix.FncApndixFrtaxsGetall)
	r.POST("/apndix/frtaxs/update", fncApndix.FncApndixFrtaxsUpdate)
	r.POST("/apndix/frtaxs/downld", fncApndix.FncApndixFrtaxsDownld)
	r.POST("/apndix/milege/getall", fncApndix.FncApndixMilegeGetall)
	r.POST("/apndix/milege/update", fncApndix.FncApndixMilegeUpdate)
	r.POST("/apndix/milege/downld", fncApndix.FncApndixMilegeDownld)
	r.POST("/apndix/fllist/getall", fncApndix.FncApndixFllistGetall)
	r.POST("/apndix/fllist/downld", fncApndix.FncApndixFllistDownld)

	// Handle web link API all user
	r.GET("/allusr/status", fncAllusr.FncAllusrStatusPrcess)
	r.GET("/allusr/applst", fncAllusr.FncAllusrApplstGetall)
	r.POST("/allusr/loginx", fncAllusr.FncAllusrLoginxHandle)
	r.GET("/allusr/tokenx", fncAllusr.FncAllusrTokenxHandle)
	r.GET("/allusr/logout", fncAllusr.FncAllusrLogoutHandle)
	r.POST("/allusr/regist", fncAllusr.FncAllusrRegistHandle)
	r.POST("/allusr/getall", fncAllusr.FncAllusrUsrlstGetall)
	r.GET("/allusr/delete/:usrnme", fncAllusr.FncAllusrDeleteHandle)

	// Handle web link API Passangger list
	r.POST("/psglst/prcess", fncPsglst.FncPsglstPrcessMainpg)
	r.POST("/psglst/psgdtl/downld", fncPsglst.FncPsglstPsgdtlDownld)
	r.POST("/psglst/psgdtl/getall", fncPsglst.FncPsglstPsgdtlGetall)
	r.POST("/psglst/psgsmr/getall", fncPsglst.FncPsglstPsgsmrGetall)
	r.POST("/psglst/errlog/getall", fncPsglst.FncPsglstErrlogGetall)
	r.GET("/psglst/actlog/getall", fncPsglst.FncPsglstActlogGetall)
	r.POST("/psglst/psgdtl/upload", fncPsglst.FncPsglstPsgdtlUpload)
	r.POST("/psglst/psgdtl/update/:dvsion", fncPsglst.FncPsglstPsgdtlUpdate)

	// Run server
	r.Run("0.0.0.0:" + fncApndix.Ptgolg)
}
