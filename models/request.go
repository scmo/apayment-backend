package models

import (
	"github.com/astaxie/beego/orm"
	"math/big"
)

func init() {
	// Register model
	orm.RegisterModel(new(Request))
}

type Request struct {
	Id                     int64 `json:"id"`
	User                   *User `orm:"rel(fk)" json:"user"`

	Address                string `json:"address"`
	Contributions          []*Contribution `orm:"-" json:"contributions"`
	Remark                 string `orm:"-" json:"remark"`
	Created                *big.Int `orm:"-" json:"created"`
	Modified               *big.Int `orm:"-" json:"modified"`

	Inspector              *User `orm:"rel(fk);null" json:"inspector"`
	ContributionsWithLacks []*Contribution `orm:"-" json:"contributionsWithLacks"`

	GVE                    []*GVE `orm:"-" json:"gve"`
	Payments               []*Payment `orm:"-" json:"payments"`
}

type GVE struct {
	Amount     uint16 `json:"amount"`
	PointGroup *PointGroup `json:"pointGroup"`
}

type Payment struct {
	From      string `orm:"-" json:"from"`
	Amount    *big.Int `json:"amount"`
	Timestamp *big.Int `orm:"-" json:"timestamp"`
}