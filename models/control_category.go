package models

import "github.com/astaxie/beego/orm"

type ControlCategory struct {
	Id                int64              `json:"id"`
	ControlCategoryId string             `json:"controlCategoryId"`
	Abbreviation      string             `json:"abbreviation"`
	ControlCategory   string             `json:"controlCategory"`
	PointGroups       []*PointGroup      `orm:"reverse(many)" json:"pointGroups"`
	Contribution      *Contribution      `orm:"rel(fk);null" json:"-"`
}

func init() {
	// Register model
	orm.RegisterModel(new(ControlCategory))
}
