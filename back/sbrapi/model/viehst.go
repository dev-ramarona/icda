package mdlSbrapi

type MdlViehstPrcessInputx struct {
	Datefl int32  `json:"datefl" bson:"datefl,omitempty"`
	Airlfl string `json:"airlfl" bson:"airlfl,omitempty"`
	Depart string `json:"depart" bson:"depart,omitempty"`
	Flnbfl string `json:"flnbfl" bson:"flnbfl,omitempty"`
	Worker int32  `json:"worker" bson:"worker,omitempty"`
}

type MdlViehstRawtmpDtbase struct {
	Prmkey string `json:"prmkey" bson:"prmkey,omitempty"`
	Datefl int32  `json:"datefl" bson:"datefl,omitempty"`
	Airlfl string `json:"airlfl" bson:"airlfl,omitempty"`
	Flnbfl string `json:"flnbfl" bson:"flnbfl,omitempty"`
	Rawdta string `json:"rawdta" bson:"rawdta,omitempty"`
}
