package models

import "github.com/astaxie/beego/orm"

type Role struct {
	Id    int64 `json:"id"`
	Name  string
	Users []*User `orm:"reverse(many)" json:"-"`
}

func init() {
	// Register model
	orm.RegisterModel(new(Role))
}