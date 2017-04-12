package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id             uint32
	Username       string  `orm:"unique" json:"username"`
	Password       string  `json:"password"`
	Firstname      string  `json:"firstname"`
	Lastname       string  `json:"lastname"`
	Roles          []*Role `orm:"rel(m2m)" json:"roles"`
	JwtToken       string  `orm:"-" json:"token"`
	EtherumAddress string
}

func init() {
	// Register model
	orm.RegisterModel(new(User))
}
