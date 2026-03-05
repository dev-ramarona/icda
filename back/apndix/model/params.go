package mdlApndix

// Acepted edit params
type MdlApndixAcpedtDtbase struct {
	Params string `json:"params,omitempty" bson:"params,omitempty"`
	Length int    `json:"length,omitempty" bson:"length,omitempty"`
	Dvsion string `json:"dvsion,omitempty" bson:"dvsion,omitempty"`
}

// Status process api
type MdlApndixStatusPrcess struct {
	Sbrapi float64 `json:"sbrapi" bson:"sbrapi"`
	Action float64 `json:"action" bson:"action"`
}

// Province data
type MdlApndixProvncDtbase struct {
	Routfl string `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Provnc string `json:"provnc,omitempty" bson:"provnc,omitempty"`
}

// District data
type MdlApndixDstrctDtbase struct {
	Depart string `json:"depart,omitempty" bson:"depart,omitempty"`
}

// Flight hour data
type MdlApndixFlhourDtbase struct {
	Prmkey string  `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Airlfl string  `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Routfl string  `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Flnbfl string  `json:"flnbfl,omitempty" bson:"flnbfl,omitempty"`
	Flhour float64 `json:"flhour,omitempty" bson:"flhour,omitempty"`
	Timefl int64   `json:"timefl,omitempty" bson:"timefl,omitempty"`
	Timerv int64   `json:"timerv,omitempty" bson:"timerv,omitempty"`
	Timeup int64   `json:"timeup,omitempty" bson:"timeup,omitempty"`
	Dateup int32   `json:"dateup,omitempty" bson:"dateup,omitempty"`
	Datend int32   `json:"datend,omitempty" bson:"datest,omitempty"`
	Airtyp string  `json:"airtyp,omitempty" bson:"airtyp,omitempty"`
	Airmls int32   `json:"airmls,omitempty" bson:"airmls,omitempty"`
	Hstory string  `json:"hstory,omitempty" bson:"hstory,omitempty"`
}

// Flight number list
type MdlApndixFlnblsDtbase struct {
	Prmkey string `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Airlfl string `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Flnbfl string `json:"flnbfl,omitempty" bson:"flnbfl,omitempty"`
	Routfl string `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Datefl int32  `json:"datefl,omitempty" bson:"datefl,omitempty"`
	Hstory string `json:"hstory,omitempty" bson:"hstory,omitempty"`
}

// Fare base data
type MdlApndixFrbaseDtbase struct {
	Prmkey string `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Scdkey string `json:"scdkey,omitempty" bson:"scdkey,omitempty"`
	Airlfl string `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Clssfl string `json:"clssfl,omitempty" bson:"clssfl,omitempty"`
	Routfl string `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Frbcde string `json:"frbcde,omitempty" bson:"frbcde,omitempty"`
	Frbnta int32  `json:"frbnta,omitempty" bson:"frbnta,omitempty"`
	Frbsbr int32  `json:"frbsbr,omitempty" bson:"frbsbr,omitempty"`
	Datend int32  `json:"datend,omitempty" bson:"datest,omitempty"`
	Hstory string `json:"hstory,omitempty" bson:"hstory,omitempty"`
}

// Fare taxs data
type MdlApndixFrtaxsDtbase struct {
	Prmkey string  `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Airlfl string  `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Cbinfl string  `json:"cbinfl,omitempty" bson:"cbinfl,omitempty"`
	Routfl string  `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Ftppnx float32 `json:"ftppnx,omitempty" bson:"ftppnx,omitempty"`
	Ftothr string  `json:"ftothr,omitempty" bson:"ftothr,omitempty"`
	Ftaptx int32   `json:"ftaptx,omitempty" bson:"ftaptx,omitempty"`
	Ftfuel int32   `json:"ftfuel,omitempty" bson:"ftfuel,omitempty"`
	Ftiwjr int32   `json:"ftiwjr,omitempty" bson:"ftiwjr,omitempty"`
	Ftaxyr int32   `json:"ftaxyr,omitempty" bson:"ftaxyr,omitempty"`
	Datend int32   `json:"datend,omitempty" bson:"datest,omitempty"`
	Hstory string  `json:"hstory,omitempty" bson:"hstory,omitempty"`
}

// Milege data
type MdlApndixMilegeDtbase struct {
	Routfl string `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Milege int64  `json:"milege,omitempty" bson:"milege,omitempty"`
}

// Class level data
type MdlApndixClsslvDtbase struct {
	Clssfl string  `json:"clssfl,omitempty" bson:"clssfl,omitempty"`
	Clsslv int32   `json:"clsslv,omitempty" bson:"clsslv,omitempty"`
	Cbinfl string  `json:"cbinfl,omitempty" bson:"cbinfl,omitempty"`
	Clssdc float64 `json:"clssdc,omitempty" bson:"clssdc,omitempty"`
}

// Highest FBA level data
type MdlApndixHfbalvDtbase struct {
	Airlfl string `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Clssfl string `json:"clssfl,omitempty" bson:"clssfl,omitempty"`
	Routfl string `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Levelr int32  `json:"levelr,omitempty" bson:"levelr,omitempty"`
	Hfbabt int32  `json:"hfbabt,omitempty" bson:"hfbabt,omitempty"`
	Source string `json:"source,omitempty" bson:"source,omitempty"`
}

// Currency database
type MdlApndixCurrcvDtbase struct {
	Crctry string  `json:"crctry,omitempty" bson:"crctry,omitempty"`
	Crcode string  `json:"crcode,omitempty" bson:"crcode,omitempty"`
	Crname string  `json:"crname,omitempty" bson:"crname,omitempty"`
	Crrate float64 `json:"crrate,omitempty" bson:"crrate,omitempty"`
}
