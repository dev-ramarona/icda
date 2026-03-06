package fncApndix

import (
	mdlApndix "back/apndix/model"
	"bufio"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

// Load Environment Variables from .env file
func FncGlobalMainprLoadnv(filenm string) {
	filenw, err := os.Open(filenm)
	if err != nil {
		panic("Error opening .env file:" + err.Error())
	}
	defer filenw.Close()

	// Scan the file line by line
	scnner := bufio.NewScanner(filenw)
	for scnner.Scan() {
		linenw := strings.TrimSpace(scnner.Text())
		if linenw == "" || strings.HasPrefix(linenw, "#") {
			continue
		}
		partnw := strings.SplitN(linenw, "=", 2)
		if len(partnw) == 2 {
			prmkey := strings.TrimSpace(partnw[0])
			valuex := strings.TrimSpace(partnw[1])
			os.Setenv(prmkey, valuex)
		}
	}
}

// manual load
var Status = mdlApndix.MdlApndixStatusPrcess{Sbrapi: 0, Action: 0}
var Client *mongo.Client
var Jwtkey []byte
var Ipalow []string
var Secure = false
var Dbases, Urlmgo, Pcckey, Usrnme,
	Psswrd, Ptgolg, Ipadrs, Usrcok, Tknnme string

// Initial load environment variables
func init() {
	FncGlobalMainprLoadnv("../front/.env")
	Jwtkey = []byte(os.Getenv("NEXT_PUBLIC_JWT_SECRET"))
	Dbases = os.Getenv("NEXT_PUBLIC_VAR_DTBASE")
	Urlmgo = os.Getenv("NEXT_PUBLIC_URI_MONGOS")
	Pcckey = os.Getenv("NEXT_PUBLIC_SBR_PCCKEY")
	Usrnme = os.Getenv("NEXT_PUBLIC_SBR_USRNME")
	Psswrd = os.Getenv("NEXT_PUBLIC_SBR_PSSWRD")
	Ptgolg = os.Getenv("NEXT_PUBLIC_PRT_GOLANG")
	Ipadrs = os.Getenv("NEXT_PUBLIC_IPV_ADRESS")
	Tknnme = os.Getenv("NEXT_PUBLIC_TKN_COOKIE")
	Ipalow = strings.Split(os.Getenv("NEXT_PUBLIC_IPV_ALLOWD"), "|")
	Tmpscr, err := strconv.ParseBool(os.Getenv("NEXT_PUBLIC_IPV_SECURE"))
	if err == nil {
		Secure = Tmpscr
	}
}
