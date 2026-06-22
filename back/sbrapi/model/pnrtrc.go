package mdlSbrapi

type MdlSbrapiPnrtrcDtbase struct {
	Actcde string `json:"actcde" bson:"actcde,omitempty"`
	Depart string `json:"depart" bson:"depart,omitempty"`
	Agtnme string `json:"agtnme" bson:"agtnme,omitempty"`
	Routfl string `json:"routfl" bson:"routfl,omitempty"`
	Flnbfl string `json:"flnbfl" bson:"flnbfl,omitempty"`
	Timefl int64  `json:"timefl" bson:"timefl,omitempty"`
	Timenw int64  `json:"timenw" bson:"timenw,omitempty"`
	Pnrcde string `json:"pnrcde" bson:"pnrcde,omitempty"`
	Clssbk string `json:"clssbk" bson:"clssbk,omitempty"`
	Issued string `json:"issued" bson:"issued,omitempty"`
	Totpax int32  `json:"totpax" bson:"totpax,omitempty"`
}
