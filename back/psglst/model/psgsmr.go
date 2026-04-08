package mdlPsglst

type MdlPsglstPsgsmrDtbase struct {
	Prmkey string  `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Airlfl string  `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Provnc string  `json:"provnc,omitempty" bson:"provnc,omitempty"`
	Depart string  `json:"depart,omitempty" bson:"depart,omitempty"`
	Flnbfl string  `json:"flnbfl,omitempty" bson:"flnbfl,omitempty"`
	Flnbjn string  `json:"flnbjn,omitempty" bson:"flnbjn,omitempty"`
	Routfl string  `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Ndayfl string  `json:"ndayfl,omitempty" bson:"ndayfl,omitempty"`
	Datefl int32   `json:"datefl,omitempty" bson:"datefl,omitempty"`
	Mnthfl int32   `json:"mnthfl,omitempty" bson:"mnthfl,omitempty"`
	Flstat string  `json:"flstat,omitempty" bson:"flstat,omitempty"`
	Seatcn string  `json:"seatcn,omitempty" bson:"seatcn,omitempty"`
	Airtyp string  `json:"airtyp,omitempty" bson:"airtyp,omitempty"`
	Flhour float64 `json:"flhour,omitempty" bson:"flhour,omitempty"`
	Totnta float64 `json:"totnta,omitempty" bson:"totnta,omitempty"`
	Tottyq float64 `json:"tottyq,omitempty" bson:"tottyq,omitempty"`
	Totpax int64   `json:"totpax,omitempty" bson:"totpax,omitempty"`
	Totfae float64 `json:"totfae,omitempty" bson:"totfae,omitempty"`
	Totqfr float64 `json:"totqfr,omitempty" bson:"totqfr,omitempty"`
	Totrph float64 `json:"totrph,omitempty" bson:"totrph,omitempty"`
}
type MdlPsglstPsgsmrFrtend struct {
	Prmkey string  `json:"prmkey" bson:"prmkey"`
	Airlfl string  `json:"airlfl" bson:"airlfl"`
	Provnc string  `json:"provnc" bson:"provnc"`
	Depart string  `json:"depart" bson:"depart"`
	Flnbfl string  `json:"flnbfl" bson:"flnbfl"`
	Flnbjn string  `json:"flnbjn" bson:"flnbjn"`
	Routfl string  `json:"routfl" bson:"routfl"`
	Ndayfl string  `json:"ndayfl" bson:"ndayfl"`
	Datefl int32   `json:"datefl" bson:"datefl"`
	Mnthfl int32   `json:"mnthfl" bson:"mnthfl"`
	Flstat string  `json:"flstat" bson:"flstat"`
	Seatcn string  `json:"seatcn" bson:"seatcn"`
	Airtyp string  `json:"airtyp" bson:"airtyp"`
	Flhour float64 `json:"flhour" bson:"flhour"`
	Totnta float64 `json:"totnta" bson:"totnta"`
	Tottyq float64 `json:"tottyq" bson:"tottyq"`
	Totpax int64   `json:"totpax" bson:"totpax"`
	Totfae float64 `json:"totfae" bson:"totfae"`
	Totqfr float64 `json:"totqfr" bson:"totqfr"`
	Totrph float64 `json:"totrph" bson:"totrph"`
}
