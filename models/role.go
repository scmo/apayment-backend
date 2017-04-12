package models

import "github.com/astaxie/beego/orm"

type Role struct {
	Id    uint32
	Name  string
	Users []*User `orm:"reverse(many)" json:"-"`
}

func init() {
	// Register model
	orm.RegisterModel(new(Role))
}