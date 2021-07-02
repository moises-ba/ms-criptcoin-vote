package model

type Vote struct {
	Uuid     string `bson:"_id,omitempty" json:"uuid,omitempty"`
	UserId   string `bson:"userId,omitempty" json:"userId,omitempty"`
	CoinId   string `bson:"CoinId,omitempty" json:"CoinId,omitempty"`
	Approved bool   `bson:"approved,omitempty" json:"approved,omitempty"`
}

type User struct {
	Uuid   string `bson:"_id,omitempty" json:"uuid,omitempty"`
	UserId string `bson:"userId,omitempty" json:"userId,omitempty"`
	Name   string `bson:"userName,omitempty" json:"userName,omitempty"`
	Email  string `bson:"email,omitempty" json:"email,omitempty"`
}

type Coin struct {
	Id                    string `bson:"_id,omitempty" json:"uuid,omitempty"` //sigla da moeda
	Name                  string `bson:"name,omitempty" json:"name,omitempty"`
	Description           string `bson:"description,omitempty" json:"description,omitempty"`
	TotalApprovedVotes    int    //totalizadores ignorados na persistencia
	TotalDisapprovedVotes int    //totalizadores ignorados na persistencia
}
