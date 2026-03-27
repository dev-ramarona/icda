package mdlApndix

// applist data apndix
type MdlApdnixParamsInputx struct {
	Pagedb_apndix string `json:"pagedb_apndix,omitempty" bson:"pagedb_apndix,omitempty"`
	Datefl_apndix string `json:"datefl_apndix,omitempty" bson:"datefl_apndix,omitempty"`
	Airlfl_apndix string `json:"airlfl_apndix,omitempty" bson:"airlfl_apndix,omitempty"`
	Depart_apndix string `json:"depart_apndix,omitempty" bson:"depart_apndix,omitempty"`
	Flnbfl_apndix string `json:"flnbfl_apndix,omitempty" bson:"flnbfl_apndix,omitempty"`
	Routfl_apndix string `json:"routfl_apndix,omitempty" bson:"routfl_apndix,omitempty"`
	Clssfl_apndix string `json:"clssfl_apndix,omitempty" bson:"clssfl_apndix,omitempty"`
	Pagenw_apndix int    `json:"pagenw_apndix,omitempty" bson:"pagenw_apndix,omitempty"`
	Limitp_apndix int    `json:"limitp_apndix,omitempty" bson:"limitp_apndix,omitempty"`
}

// applist data apndix
type MdlApndixApplstDtbase struct {
	Apndix string `json:"apndix,omitempty" bson:"apndix,omitempty"`
}

// Acepted edit params
type MdlApndixAcpedtDtbase struct {
	Params string `json:"params,omitempty" bson:"params,omitempty"`
	Length int    `json:"length,omitempty" bson:"length,omitempty"`
	Dvsion string `json:"dvsion,omitempty" bson:"dvsion,omitempty"`
}

// Status process api
type MdlApndixStatusPrcess struct {
	Sbrapi float64 `json:"sbrapi" bson:"sbrapi"`
	Action string  `json:"action" bson:"action"`
}

// Province data
type MdlApndixProvncDtbase struct {
	Routfl string `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Provnc string `json:"provnc,omitempty" bson:"provnc,omitempty"`
	Updtby string `json:"updtby,omitempty" bson:"updtby,omitempty"`
}
type MdlApndixProvncFrntnd struct {
	Prmkey string `json:"prmkey" bson:"prmkey"`
	Routfl string `json:"routfl" bson:"routfl"`
	Provnc string `json:"provnc" bson:"provnc"`
	Updtby string `json:"updtby" bson:"updtby"`
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
	Datend int32   `json:"datend,omitempty" bson:"datend,omitempty"`
	Airtyp string  `json:"airtyp,omitempty" bson:"airtyp,omitempty"`
	Airmls int32   `json:"airmls,omitempty" bson:"airmls,omitempty"`
	Hstory string  `json:"hstory,omitempty" bson:"hstory,omitempty"`
	Updtby string  `json:"updtby,omitempty" bson:"updtby,omitempty"`
}
type MdlApndixFlhourFrntnd struct {
	Prmkey string  `json:"prmkey" bson:"prmkey"`
	Airlfl string  `json:"airlfl" bson:"airlfl"`
	Routfl string  `json:"routfl" bson:"routfl"`
	Flnbfl string  `json:"flnbfl" bson:"flnbfl"`
	Flhour float64 `json:"flhour" bson:"flhour"`
	Timefl int64   `json:"timefl" bson:"timefl"`
	Timerv int64   `json:"timerv" bson:"timerv"`
	Timeup int64   `json:"timeup" bson:"timeup"`
	Dateup int32   `json:"dateup" bson:"dateup"`
	Datend int32   `json:"datend" bson:"datend"`
	Airtyp string  `json:"airtyp" bson:"airtyp"`
	Airmls int32   `json:"airmls" bson:"airmls"`
	Hstory string  `json:"hstory" bson:"hstory"`
	Updtby string  `json:"updtby" bson:"updtby"`
}

// Flight number list
type MdlApndixFlnblsDtbase struct {
	Prmkey string `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Airlfl string `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Flnbfl string `json:"flnbfl,omitempty" bson:"flnbfl,omitempty"`
	Routmx string `json:"routmx,omitempty" bson:"routmx,omitempty"`
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
	Datend int32  `json:"datend,omitempty" bson:"datend,omitempty"`
	Hstory string `json:"hstory,omitempty" bson:"hstory,omitempty"`
	Updtby string `json:"updtby,omitempty" bson:"updtby,omitempty"`
}
type MdlApndixFrbaseFrntnd struct {
	Prmkey string `json:"prmkey" bson:"prmkey"`
	Scdkey string `json:"scdkey" bson:"scdkey"`
	Airlfl string `json:"airlfl" bson:"airlfl"`
	Clssfl string `json:"clssfl" bson:"clssfl"`
	Routfl string `json:"routfl" bson:"routfl"`
	Frbcde string `json:"frbcde" bson:"frbcde"`
	Frbnta int32  `json:"frbnta" bson:"frbnta"`
	Frbsbr int32  `json:"frbsbr" bson:"frbsbr"`
	Datend int32  `json:"datend" bson:"datend"`
	Hstory string `json:"hstory" bson:"hstory"`
	Updtby string `json:"updtby" bson:"updtby"`
}

// Fare taxs data
type MdlApndixFrtaxsDtbase struct {
	Prmkey string  `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Airlfl string  `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Cbinfl string  `json:"cbinfl,omitempty" bson:"cbinfl,omitempty"`
	Routfl string  `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Ftppnx float32 `json:"ftppnx" bson:"ftppnx"`
	Ftaptx int32   `json:"ftaptx" bson:"ftaptx"`
	Ftfuel int32   `json:"ftfuel" bson:"ftfuel"`
	Ftiwjr int32   `json:"ftiwjr" bson:"ftiwjr"`
	Ftaxyr int32   `json:"ftaxyr" bson:"ftaxyr"`
	Datend int32   `json:"datend" bson:"datend"`
	Ftothr string  `json:"ftothr,omitempty" bson:"ftothr,omitempty"`
	Hstory string  `json:"hstory,omitempty" bson:"hstory,omitempty"`
	Updtby string  `json:"updtby,omitempty" bson:"updtby,omitempty"`
}
type MdlApndixFrtaxsFrntnd struct {
	Prmkey string  `json:"prmkey" bson:"prmkey"`
	Airlfl string  `json:"airlfl" bson:"airlfl"`
	Cbinfl string  `json:"cbinfl" bson:"cbinfl"`
	Routfl string  `json:"routfl" bson:"routfl"`
	Ftppnx float32 `json:"ftppnx" bson:"ftppnx"`
	Ftaptx int32   `json:"ftaptx" bson:"ftaptx"`
	Ftfuel int32   `json:"ftfuel" bson:"ftfuel"`
	Ftiwjr int32   `json:"ftiwjr" bson:"ftiwjr"`
	Ftaxyr int32   `json:"ftaxyr" bson:"ftaxyr"`
	Datend int32   `json:"datend" bson:"datend"`
	Ftothr string  `json:"ftothr" bson:"ftothr"`
	Hstory string  `json:"hstory" bson:"hstory"`
	Updtby string  `json:"updtby" bson:"updtby"`
}

// Milege data
type MdlApndixMilegeDtbase struct {
	Routfl string `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Milege int64  `json:"milege,omitempty" bson:"milege,omitempty"`
}

// Class level data
type MdlApndixClsslvDtbase struct {
	Rbdcls string  `json:"rbdcls,omitempty" bson:"rbdcls,omitempty"`
	Lvlcls int32   `json:"lvlcls,omitempty" bson:"lvlcls,omitempty"`
	Cbncls string  `json:"cbncls,omitempty" bson:"cbncls,omitempty"`
	Dscont float64 `json:"dscont,omitempty" bson:"dscont,omitempty"`
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

// Flight list database
type MdlApndixFllistDtbase struct {
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
