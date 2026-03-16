package mdlAllusr

import "github.com/golang-jwt/jwt/v5"

type MdlAllusrParamsLoginx struct {
	Usrnme string `json:"usrnme,omitempty" bson:"usrnme,omitempty"`
	Psswrd string `json:"psswrd,omitempty" bson:"psswrd,omitempty"`
}
type MdlAllusrInputxSrcprm struct {
	Usrnme string `json:"usrnme,omitempty" bson:"usrnme,omitempty"`
	Stfnme string `json:"stfnme,omitempty" bson:"stfnme,omitempty"`
	Stfeml string `json:"stfeml,omitempty" bson:"stfeml,omitempty"`
	Pagenw int    `json:"pagenw,omitempty" bson:"pagenw,omitempty"`
	Limitp int    `json:"limitp,omitempty" bson:"limitp,omitempty"`
}
type MdlAllusrUsrlstDtbase struct {
	Action string   `json:"action,omitempty" bson:"action,omitempty"`
	Usrnme string   `json:"usrnme,omitempty" bson:"usrnme,omitempty"`
	Psswrd string   `json:"psswrd,omitempty" bson:"psswrd,omitempty"`
	Stfnme string   `json:"stfnme,omitempty" bson:"stfnme,omitempty"`
	Stfeml string   `json:"stfeml,omitempty" bson:"stfeml,omitempty"`
	Access []string `json:"access,omitempty" bson:"access,omitempty"`
	Keywrd []string `json:"keywrd,omitempty" bson:"keywrd,omitempty"`
}
type MdlAllusrFrntndFormat struct {
	Usrnme string   `json:"usrnme" bson:"usrnme"`
	Stfnme string   `json:"stfnme" bson:"stfnme"`
	Stfeml string   `json:"stfeml" bson:"stfeml"`
	Access []string `json:"access" bson:"access"`
	Keywrd []string `json:"keywrd" bson:"keywrd"`
}
type MdlAllusrTokensFormat struct {
	Stfnme string   `json:"stfnme"`
	Usrnme string   `json:"usrnme"`
	Stfeml string   `json:"stfeml"`
	Access []string `json:"access"`
	Keywrd []string `json:"keywrd"`
	jwt.RegisteredClaims
}
type MdlAllusrApplstDtbase struct {
	Pagenb int32  `json:"pagenb" bson:"pagenb"`
	Prmkey string `json:"prmkey" bson:"prmkey"`
	Detail string `json:"detail" bson:"detail"`
}
