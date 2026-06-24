package mdlJeddah

type MdlJeddahPramsInputx struct {
	Airlfl_jeddah string `json:"airlfl_jeddah,omitempty" bson:"airlfl_jeddah,omitempty"`
	Flnbfl_jeddah string `json:"flnbfl_jeddah,omitempty" bson:"flnbfl_jeddah,omitempty"`
	Depart_jeddah string `json:"depart_jeddah,omitempty" bson:"depart_jeddah,omitempty"`
	Routfl_jeddah string `json:"routfl_jeddah,omitempty" bson:"routfl_jeddah,omitempty"`
	Pnrcde_jeddah string `json:"pnrcde_jeddah,omitempty" bson:"pnrcde_jeddah,omitempty"`
	Pagenw_jeddah int    `json:"pagenw_jeddah,omitempty" bson:"pagenw_jeddah,omitempty"`
	Limitp_jeddah int    `json:"limitp_jeddah,omitempty" bson:"limitp_jeddah,omitempty"`
	Worker_jeddah int    `json:"worker_jeddah,omitempty" bson:"worker_jeddah,omitempty"`
}

type MdlJeddahFlnblsDtbase struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Airlfl string `json:"airlfl" bson:"airlfl,omitempty"`
	Flnbfl string `json:"flnbfl" bson:"flnbfl,omitempty"`
	Datefl int32  `json:"datefl" bson:"datefl,omitempty"`
	Depart string `json:"depart" bson:"depart,omitempty"`
}

type MdlJeddahPnrsmrDtbase struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Pnrcde string `json:"pnrcde" bson:"pnrcde,omitempty"`
	Pnrsrc string `json:"pnrsrc" bson:"pnrsrc,omitempty"`
	Agtnme string `json:"agtnme" bson:"agtnme,omitempty"`
	Flnbsg string `json:"flnbsg" bson:"flnbsg,omitempty"`
	Routsg string `json:"routsg" bson:"routsg,omitempty"`
	Clssbk string `json:"clssbk" bson:"clssbk,omitempty"`
	Timefl int64  `json:"timefl" bson:"timefl,omitempty"`
	Timerv int64  `json:"timerv" bson:"timerv,omitempty"`
	Timecx int64  `json:"timecx" bson:"timecx,omitempty"`
	Routpv int64  `json:"routpv" bson:"routpv,omitempty"`
	Drtion int32  `json:"drtion" bson:"drtion,omitempty"`
	Spltfr string `json:"spltfr" bson:"spltfr,omitempty"`
	Spltto string `json:"spltto" bson:"spltto,omitempty"`
	Totpax int32  `json:"totpax" bson:"totpax"`
	Totbok int32  `json:"totbok" bson:"totbok"`
	Totcxl int32  `json:"totcxl" bson:"totcxl"`
	Totisd int32  `json:"totisd" bson:"totisd"`
	Totori int32  `json:"totori" bson:"totori,omitempty"`
	Frbase int32  `json:"frbase" bson:"frbase"`
}

type MdlJeddahPnrsmrCmpare struct {
	Totisd int32 `json:"totisd" bson:"totisd"`
	Totbok int32 `json:"totbok" bson:"totbok"`
	Totcxl int32 `json:"totcxl" bson:"totcxl"`
}
