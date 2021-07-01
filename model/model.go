package model

type Vote struct {
	Uuid   string `bson:"_id,omitempty" json:"uuid,omitempty"`
	User   string `bson:"user,omitempty" json:"user,omitempty"`
	CoinId string `bson:"CoinId,omitempty" json:"CoinId,omitempty"`
	Yes    bool   `bson:"yes,omitempty" json:"yes,omitempty"`
}

type User struct {
	Uuid   string `bson:"_id,omitempty" json:"uuid,omitempty"`
	UserId string `bson:"userId,omitempty" json:"userId,omitempty"`
	Name   string `bson:"userName,omitempty" json:"userName,omitempty"`
	Email  string `bson:"email,omitempty" json:"email,omitempty"`
}
