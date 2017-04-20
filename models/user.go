package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id             int64 `json:"id"`
	Username       string  `orm:"unique" json:"username"`
	Password       string  `json:"password"`

	Roles          []*Role `orm:"rel(m2m)" json:"roles"`
	JwtToken       string  `orm:"-" json:"token"`
	EtherumAddress string

	Firstname      string  `json:"firstname"`
	Lastname       string  `json:"lastname"`
	Phone string  `json:"phone"`
	EMail string  `json:"email"`
	Address string  `json:"address"`
	ZIP uint16  `json:"zip"`
	Place string  `json:"place"`
	Plant *Plant `orm:"rel(one);null" json:"plant"`
}

func init() {
	// Register model
	orm.RegisterModel(new(User))
}
