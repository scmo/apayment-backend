package models

import "github.com/astaxie/beego/orm"

type Contribution struct {
	Id                int64                	`json:"id"`
	Code              uint16                `json:"code"`
	Name              string                `json:"name"`
	ControlCategories []*ControlCategory    `json:"controlCategories" orm:"reverse(many)"`
}

func init() {
	// Register model
	orm.RegisterModel(new(Contribution))
}