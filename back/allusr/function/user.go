package fncAllusr

import (
	mdlAllusr "back/allusr/model"
	fncApndix "back/apndix/function"
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Get status data process
func FncAllusrStatusPrcess(c *gin.Context) {
	c.JSON(http.StatusOK, fncApndix.Status)
}

// Handle Login function
func FncAllusrLoginxHandle(c *gin.Context) {

	// Variable login
	var loginp mdlAllusr.MdlAllusrParamsLoginx
	var usrdbs mdlAllusr.MdlAllusrUsrlstDtbase

	// Bind JSON body input to var Login parameter
	if err := c.BindJSON(&loginp); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Select database and collection
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("allusr_usrlst")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find user by username in database
	err := tablex.FindOne(contxt, bson.M{"usrnme": loginp.Usrnme}).Decode(&usrdbs)
	if err != nil {
		c.JSON(401, gin.H{"error": "user"})
		return
	}

	// Compare provided password with stored password hash
	err = bcrypt.CompareHashAndPassword([]byte(usrdbs.Psswrd), []byte(loginp.Psswrd))
	if err != nil {
		c.JSON(401, gin.H{"error": "user"})
		return
	}

	// Generate JWT Token
	claimp := &mdlAllusr.MdlAllusrTokensFormat{
		Usrnme: usrdbs.Usrnme,
		Stfnme: usrdbs.Stfnme,
		Stfeml: usrdbs.Stfeml,
		Access: usrdbs.Access,
		Keywrd: usrdbs.Keywrd,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
		},
	}

	// Translate JWT Token to strings
	tknraw := jwt.NewWithClaims(jwt.SigningMethodHS256, claimp)
	tknstr, err := tknraw.SignedString(fncApndix.Jwtkey)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err})
		return
	}

	// Set the JWT token in the cookie
	c.JSON(200, gin.H{"ok": tknstr})
}

// Handle Logout function
func FncAllusrLogoutHandle(c *gin.Context) {

	// Delete Cookie
	c.SetCookie(fncApndix.Tknnme, "", -1, "/", "", fncApndix.Secure, true)
	c.JSON(200, "Logout")
}

// Handle Logout function
func FncAllusrTokenxHandle(c *gin.Context) {

	// Get cookie
	cookie := c.GetHeader("Authorization")
	if cookie == "" {
		c.String(401, "Authorization header missing")
		return
	}

	// Convert JWT to claims
	tokenx, err := jwt.ParseWithClaims(cookie, &mdlAllusr.MdlAllusrTokensFormat{},
		func(token *jwt.Token) (interface{}, error) {
			return fncApndix.Jwtkey, nil
		})

	// Final Result
	if val, ist := tokenx.Claims.(*mdlAllusr.MdlAllusrTokensFormat); ist && err == nil {
		c.JSON(200, val)
		return
	}
	c.JSON(500, "Loggin First")
}

func FncAllusrApplstGetall(c *gin.Context) {

	// Select database and collection
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("allusr_applst")
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
	var slices = []mdlAllusr.MdlAllusrApplstDtbase{}
	for datarw.Next(contxt) {
		var object mdlAllusr.MdlAllusrApplstDtbase
		if err := datarw.Decode(&object); err == nil {
			slices = append(slices, object)
		}
	}

	// Send token to frontend
	c.JSON(200, slices)
}

// Create user
func FncAllusrRegistHandle(c *gin.Context) {

	// Bind JSON input
	var usript mdlAllusr.MdlAllusrUsrlstDtbase
	if err := c.BindJSON(&usript); err != nil {
		c.JSON(400, "Invalid usript")
		return
	}

	// Validasi kosong
	if usript.Usrnme == "" || usript.Psswrd == "" {
		c.JSON(400, "Username or password empty")
		return
	}

	// Ambil collection
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("allusr_usrlst")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(usript.Psswrd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": "Password hashing failed",
		})
		return
	}
	usript.Psswrd = string(hashed)

	// Cek apakah username sudah ada
	if usript.Action == "regist" {
		var usrdbs mdlAllusr.MdlAllusrUsrlstDtbase
		err := tablex.FindOne(contxt, bson.M{
			"usrnme": usript.Usrnme,
		}).Decode(&usrdbs)
		if err == nil || len(usrdbs.Usrnme) > 0 {
			c.JSON(409, gin.H{
				"error": "user",
			})
			return
		}

		// Kalau error bukan ErrNoDocuments → error database
		if err != mongo.ErrNoDocuments {
			c.JSON(500, "Database error")
			return
		}

		// Insert user baru
		_, err = tablex.InsertOne(contxt, usript)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"error": "Failed to create user",
			})
			return
		}

	} else {

		// Update user
		_, err = tablex.UpdateOne(contxt,
			bson.M{"usrnme": usript.Usrnme},
			bson.M{"$set": usript})
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"error": "Failed to update user",
			})
			return
		}
	}

	// Response
	c.JSON(200, gin.H{
		"ok": "user updated",
	})
}

// Get all user
func FncAllusrUsrlstGetall(c *gin.Context) {

	// Bind JSON Body input to variable
	csvFilenm := []string{time.Now().Format("02Jan06/15:04")}
	var inputx mdlAllusr.MdlAllusrInputxSrcprm
	if err := c.BindJSON(&inputx); err != nil {
		panic(err)
	}

	// Select db and context to do
	var totidx = 0
	var slcobj interface{}
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("allusr_usrlst")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Pipeline get the data logic match
	var mtchdt = bson.A{}
	var sortdt = bson.D{{Key: "$sort", Value: bson.D{{Key: "prmkey", Value: 1}}}}
	var wg sync.WaitGroup

	// Check if data Route all is isset
	if inputx.Usrnme != "" {
		csvFilenm = append(csvFilenm, inputx.Usrnme)
		mtchdt = append(mtchdt, bson.D{{Key: "usrnme",
			Value: inputx.Usrnme}})
	}
	if inputx.Stfnme != "" {
		csvFilenm = append(csvFilenm, inputx.Stfnme)
		mtchdt = append(mtchdt, bson.D{{Key: "stfnme",
			Value: bson.D{
				{Key: "$regex", Value: inputx.Stfnme},
				{Key: "$options", Value: "i"},
			}}})
	}
	if inputx.Stfeml != "" {
		csvFilenm = append(csvFilenm, inputx.Stfeml)
		mtchdt = append(mtchdt, bson.D{{Key: "stfeml",
			Value: inputx.Stfeml}})
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
			bson.D{{Key: "$skip", Value: (max(inputx.Pagenw, 1) - 1) * inputx.Limitp}},
			bson.D{{Key: "$limit", Value: inputx.Limitp}},
		}

		// Find user by username in database
		rawDtaset, err := tablex.Aggregate(contxt, pipeln)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer rawDtaset.Close(contxt)

		// Store to slice from raw bson
		var slctmp = []mdlAllusr.MdlAllusrFrntndFormat{}
		for rawDtaset.Next(contxt) {
			slcDtaset := mdlAllusr.MdlAllusrFrntndFormat{}
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

// Delete user
func FncAllusrDeleteHandle(c *gin.Context) {

	// Get username from URL parameter
	usrnme := c.Param("usrnme")

	// Select db and context to do
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("allusr_usrlst")
	contxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Delete user by username in database
	_, err := tablex.DeleteOne(contxt, bson.M{"usrnme": usrnme})
	if err != nil {
		c.JSON(500, "Failed to delete user")
		return
	}

	// Response
	c.JSON(200, "Success")
}
