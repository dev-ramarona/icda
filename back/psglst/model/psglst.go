package mdlPsglst

type MdlPsglstErrlogDtbase struct {
	Prmkey string  `json:"prmkey" bson:"prmkey,omitempty"`
	Erstat string  `json:"erstat" bson:"erstat,omitempty"`
	Erpart string  `json:"erpart" bson:"erpart,omitempty"`
	Ersrce string  `json:"ersrce" bson:"ersrce,omitempty"`
	Erdtil string  `json:"erdtil" bson:"erdtil,omitempty"`
	Erdvsn string  `json:"erdvsn" bson:"erdvsn,omitempty"`
	Erignr string  `json:"erignr" bson:"erignr,omitempty"`
	Dateup int32   `json:"dateup" bson:"dateup,omitempty"`
	Timeup int64   `json:"timeup" bson:"timeup,omitempty"`
	Datefl int32   `json:"datefl" bson:"datefl,omitempty"`
	Airlfl string  `json:"airlfl" bson:"airlfl,omitempty"`
	Depart string  `json:"depart" bson:"depart,omitempty"`
	Flnbfl string  `json:"flnbfl" bson:"flnbfl,omitempty"`
	Paxdif string  `json:"paxdif" bson:"paxdif,omitempty"`
	Flstat string  `json:"flstat" bson:"flstat,omitempty"`
	Flhour float64 `json:"flhour" bson:"flhour,omitempty"`
	Routfl string  `json:"routfl" bson:"routfl,omitempty"`
	Updtby string  `json:"updtby" bson:"updtby,omitempty"`
	Worker int32   `json:"worker" bson:"worker,omitempty"`
}

type MdlPsglstFrcalcFrbase struct {
	Routfl string `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Cpnbfl int32  `json:"cpnbfl,omitempty" bson:"cpnbfl,omitempty"`
	Depart string `json:"depart,omitempty" bson:"depart,omitempty"`
	Arrivl string `json:"arrivl,omitempty" bson:"arrivl,omitempty"`
	Airlfl string `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Curncy string `json:"curncy,omitempty" bson:"curncy,omitempty"`
	Frbase string `json:"frbase,omitempty" bson:"frbase,omitempty"`
	Frbcnv string `json:"frbcnv,omitempty" bson:"frbcnv,omitempty"`
	Qsrcrw string `json:"qsrcrw,omitempty" bson:"qsrcrw,omitempty"`
	Qsrcnv string `json:"qsrcnv,omitempty" bson:"qsrcnv,omitempty"`
	Crrate string `json:"crrate,omitempty" bson:"crrate,omitempty"`
	Isitpr string `json:"isitpr,omitempty" bson:"isitpr,omitempty"`
	Isitit string `json:"isitit,omitempty" bson:"isitit,omitempty"`
}

type MdlPsglstFllistDtbase struct {
	Prmkey string  `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Airlfl string  `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Flnbfl string  `json:"flnbfl,omitempty" bson:"flnbfl,omitempty"`
	Timeup int64   `json:"timeup,omitempty" bson:"timeup,omitempty"`
	Timefl int64   `json:"timefl,omitempty" bson:"timefl,omitempty"`
	Timerv int64   `json:"timerv,omitempty" bson:"timerv,omitempty"`
	Datefl int32   `json:"datefl,omitempty" bson:"datefl,omitempty"`
	Mnthfl int32   `json:"mnthfl,omitempty" bson:"mnthfl,omitempty"`
	Ndayfl string  `json:"ndayfl,omitempty" bson:"ndayfl,omitempty"`
	Flstat string  `json:"flstat,omitempty" bson:"flstat,omitempty"`
	Routfl string  `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Routac string  `json:"routac,omitempty" bson:"routac,omitempty"`
	Flsarr string  `json:"flsarr,omitempty" bson:"flsarr,omitempty"`
	Routmx string  `json:"routmx,omitempty" bson:"routmx,omitempty"`
	Flhour float64 `json:"flhour,omitempty" bson:"flhour,omitempty"`
	Flrpdc int32   `json:"flrpdc,omitempty" bson:"flrpdc,omitempty"`
	Flgate string  `json:"flgate,omitempty" bson:"flgate,omitempty"`
	Depart string  `json:"depart,omitempty" bson:"depart,omitempty"`
	Arrivl string  `json:"arrivl,omitempty" bson:"arrivl,omitempty"`
	Airtyp string  `json:"airtyp,omitempty" bson:"airtyp,omitempty"`
	Aircnf string  `json:"aircnf,omitempty" bson:"aircnf,omitempty"`
	Seatcn string  `json:"seatcn,omitempty" bson:"seatcn,omitempty"`
	Autrzc int32   `json:"autrzc,omitempty" bson:"autrzc,omitempty"`
	Autrzy int32   `json:"autrzy,omitempty" bson:"autrzy,omitempty"`
	Bookdc int32   `json:"bookdc,omitempty" bson:"bookdc,omitempty"`
	Bookdy int32   `json:"bookdy,omitempty" bson:"bookdy,omitempty"`
}

type MdlPsglstActlogDtbase struct {
	Dateup int32  `json:"dateup,omitempty" bson:"dateup,omitempty"`
	Datefl int32  `json:"datefl,omitempty" bson:"datefl,omitempty"`
	Timeup int64  `json:"timeup,omitempty" bson:"timeup,omitempty"`
	Statdt string `json:"statdt,omitempty" bson:"statdt,omitempty"`
}
