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
	Id          string  `bson:"_id,omitempty" json:"uuid,omitempty"` //sigla da moeda
	Name        string  `bson:"name,omitempty" json:"name,omitempty"`
	Description string  `bson:"description,omitempty" json:"description,omitempty"`
	Votes       []*Vote //ignorar na persistencia
}

func (c *Coin) TotalApprovedVotes() int {
	return countVotes(c.Votes, true)
}

func (c *Coin) TotalDisapprovedVotes() int {
	return countVotes(c.Votes, false)
}

//conta votos contrarios ou a favor dependendo do paramentro pApproved
func countVotes(votes []*Vote, pApproved bool) int {
	total := 0
	for _, v := range votes {
		if v.Approved == pApproved {
			total++
		}
	}
	return total
}
