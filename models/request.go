package models

import "github.com/astaxie/beego/orm"

type Request struct {
	Id            int64 `json:"id"`
	Contributions []*Contribution `orm:"rel(m2m)" json:"contributions"`
	User          *User `orm:"rel(fk)" json:"user"`
	Address       string
}

func init() {
	// Register model
	orm.RegisterModel(new(Request))
}