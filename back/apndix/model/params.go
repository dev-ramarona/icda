package mdlApndix

// applist data apndix
type MdlApdnixParamsInputx struct {
	Pagedb string `json:"pagedb,omitempty" bson:"pagedb,omitempty"`
	Datefl string `json:"datefl,omitempty" bson:"datefl,omitempty"`
	Airlfl string `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Depart string `json:"depart,omitempty" bson:"depart,omitempty"`
	Flnbfl string `json:"flnbfl,omitempty" bson:"flnbfl,omitempty"`
	Routfl string `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Clssfl string `json:"clssfl,omitempty" bson:"clssfl,omitempty"`
	Pagenw int    `json:"pagenw,omitempty" bson:"pagenw,omitempty"`
	Limitp int    `json:"limitp,omitempty" bson:"limitp,omitempty"`
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
	Action float64 `json:"action" bson:"action"`
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
	Datend int32   `json:"datend,omitempty" bson:"datest,omitempty"`
	Airtyp string  `json:"airtyp,omitempty" bson:"airtyp,omitempty"`
	Airmls int32   `json:"airmls,omitempty" bson:"airmls,omitempty"`
	Hstory string  `json:"hstory,omitempty" bson:"hstory,omitempty"`
	Updtby string  `json:"updtby,omitempty" bson:"updtby,omitempty"`
}
type MdlApndixFlhourInputx struct {
	Prmkey string `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Airlfl string `json:"airlfl,omitempty" bson:"airlfl,omitempty"`
	Routfl string `json:"routfl,omitempty" bson:"routfl,omitempty"`
	Flnbfl string `json:"flnbfl,omitempty" bson:"flnbfl,omitempty"`
	Flhour string `json:"flhour,omitempty" bson:"flhour,omitempty"`
	Timefl string `json:"timefl,omitempty" bson:"timefl,omitempty"`
	Timerv string `json:"timerv,omitempty" bson:"timerv,omitempty"`
	Timeup string `json:"timeup,omitempty" bson:"timeup,omitempty"`
	Dateup string `json:"dateup,omitempty" bson:"dateup,omitempty"`
	Datend string `json:"datend,omitempty" bson:"datest,omitempty"`
	Airtyp string `json:"airtyp,omitempty" bson:"airtyp,omitempty"`
	Airmls string `json:"airmls,omitempty" bson:"airmls,omitempty"`
	Hstory string `json:"hstory,omitempty" bson:"hstory,omitempty"`
	Updtby string `json:"updtby,omitempty" bson:"updtby,omitempty"`
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
	Datend int32   `json:"datend" bson:"datest"`
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
