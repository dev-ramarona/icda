package fncJeddah

import (
	fncApndix "back/apndix/function"
	mdlJeddah "back/jeddah/model"
	fncSbrapi "back/sbrapi/function"
	mdlSbrapi "back/sbrapi/model"
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Running process hit passanggerlist daily
func FncSilvnaPrcessMainpg(c *gin.Context) {

	// protect single run
	if fncApndix.Status.Sbrapi != 0.0 {
		return
	}

	// Bind JSON Body input to variable
	inpErrlog := mdlJeddah.MdlJeddahPramsInputx{} //save
	if err := c.BindJSON(&inpErrlog); err != nil {
		panic(err)
	}
	fncApndix.Status.Sbrapi = 0.01
	var mapFlnbls = FncSilvnaFlnblsMapobj()
	var sycFlnbls = &sync.Map{}
	for airlfl, slices := range mapFlnbls {
		totWorker := 8 // totWorker := inpErrlog.Worker_jeddah
		var sycPnrcde = &sync.Map{}
		var sycWgroup sync.WaitGroup
		slcRspssn, err := fncSbrapi.FncSbrapiCrtssnMultpl(airlfl, int(totWorker))
		if err != nil {
			fmt.Println("airlfl" + airlfl + "failed")
			continue
		}

		jobPnrtrc := make(chan MdlSilvnaFlnblsDtbase, 1500)
		for i := 0; i < int(totWorker); i++ {
			if len(slcRspssn) >= i+1 {
				if slcRspssn[i].Bsttkn != "" {
					sycWgroup.Add(1)
					fmt.Println("Success Token-", i)
					go FncSilvnaPrcessWorker(&sycWgroup, jobPnrtrc, sycPnrcde, sycFlnbls, slcRspssn[i], slices)
					continue
				}
				fmt.Println("Failed Token-", i)
			}
		}

		for _, object := range slices {
			if inpErrlog.Flnbfl_jeddah == "" || inpErrlog.Flnbfl_jeddah == object.Flnbfl {
				jobPnrtrc <- object
			}
		}

		close(jobPnrtrc)
		sycWgroup.Wait()
		fncSbrapi.FncSbrapiClsssnMultpl(slcRspssn)
	}

	// Done
	fncApndix.Status.Sbrapi = 0
	c.JSON(200, gin.H{"status": "Done"})
}

func FncSilvnaPrcessWorker(sycWgroup *sync.WaitGroup,
	jobPnrtrc <-chan MdlSilvnaFlnblsDtbase, sycPnrcde, sycFlnbls *sync.Map,
	nowObjtkn mdlSbrapi.MdlSbrapiMsghdrParams,
	slices map[string]MdlSilvnaFlnblsDtbase) {
	defer sycWgroup.Done()
	var mgoTcktok []mongo.WriteModel
	for slcPnrtrc := range jobPnrtrc {
		for dateky, datefl := range map[string]string{"260701": "01JUL", "260702": "02JUL"} {
			routfl := slcPnrtrc.Depart + slcPnrtrc.Arrivl
			cmdsbr := fmt.Sprintf("VCR*OK%v/%v%v/STATUS-OK", slcPnrtrc.Flnbfl, datefl, routfl)
			strOutput, err := fncSbrapi.FncSbrapiCmdscrMainob(nowObjtkn, cmdsbr)
			outlne := strings.Split(strOutput, "\n")
			reg, _ := regexp.Compile(`[^a-zA-Z0-9/ ]`)
			if err == nil {
				for _, outrow := range outlne {
					if len(outrow) >= 64 {
						_, ist := strconv.Atoi(strings.TrimSpace(outrow[:5]))
						if ist == nil {
							objParams := MdlSilvnaTktnflDtbase{
								Prmkey: outrow[35:48] + slcPnrtrc.Flnbfl + slcPnrtrc.Depart + dateky,
								Airlvc: slcPnrtrc.Airlfl,
								Flnbvc: slcPnrtrc.Flnbfl,
								Depart: slcPnrtrc.Depart,
								Arrivl: slcPnrtrc.Arrivl,
								Nmepax: reg.ReplaceAllString(outrow[5:18], ""),
								Clssvc: reg.ReplaceAllString(outrow[21:22], ""),
								Pnrcde: reg.ReplaceAllString(outrow[27:33], ""),
								Tktnvc: reg.ReplaceAllString(outrow[35:48], ""),
							}
							mgoTcktok = append(mgoTcktok, mongo.NewUpdateOneModel().
								SetFilter(bson.M{"prmkey": objParams.Prmkey}).
								SetUpdate(bson.M{"$set": objParams}).SetUpsert(true))
							fncApndix.FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
								"silvna_tktnls": &mgoTcktok,
							}, 200)
						}
					}
				}
			} else {
				fmt.Println(err)
			}
			fmt.Println("done:", cmdsbr)
		}
	}

	fncApndix.FncApndixBulkdbBatchs(map[string]*[]mongo.WriteModel{
		"silvna_tktnls": &mgoTcktok,
	}, 0)
}

func FncSilvnaFlnblsMapobj() map[string]map[string]MdlSilvnaFlnblsDtbase {

	// Inisialisasi variabel
	fnldta := make(map[string]map[string]MdlSilvnaFlnblsDtbase)

	// Select database and collection
	tablex := fncApndix.Client.Database(fncApndix.Dbases).Collection("silvna_fllist")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{})
	if err != nil {
		panic(err)
	}
	defer datarw.Close(contxt)

	// Append to slice
	for datarw.Next(contxt) {
		var object MdlSilvnaFlnblsDtbase
		datarw.Decode(&object)
		if _, ist := fnldta[object.Airlfl]; !ist {
			fnldta[object.Airlfl] = map[string]MdlSilvnaFlnblsDtbase{}
		}
		prmkey := fmt.Sprintf("%v%v%v", object.Airlfl, object.Flnbfl, object.Depart)
		fnldta[object.Airlfl][prmkey] = object
	}

	// return data
	return fnldta
}

type MdlSilvnaFlnblsDtbase struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Airlfl string `json:"airlfl" bson:"airlfl,omitempty"`
	Flnbfl string `json:"flnbfl" bson:"flnbfl,omitempty"`
	Depart string `json:"depart" bson:"depart,omitempty"`
	Arrivl string `json:"arrivl" bson:"arrivl,omitempty"`
}

type MdlSilvnaTktnflDtbase struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Airlvc string `json:"airlvc" bson:"airlvc,omitempty"`
	Flnbvc string `json:"flnbvc" bson:"flnbvc,omitempty"`
	Depart string `json:"depart" bson:"depart,omitempty"`
	Arrivl string `json:"arrivl" bson:"arrivl,omitempty"`
	Nmepax string `json:"nmepax" bson:"nmepax,omitempty"`
	Clssvc string `json:"clssvc" bson:"clssvc,omitempty"`
	Pnrcde string `json:"pnrcde" bson:"pnrcde,omitempty"`
	Tktnvc string `json:"tktnvc" bson:"tktnvc,omitempty"`
}
