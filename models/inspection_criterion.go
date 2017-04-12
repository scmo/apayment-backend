package models

import "github.com/astaxie/beego/orm"

type InspectionCriterion struct {
	Id           uint32 `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Contribution *Contribution `orm:"rel(fk)"`
}

func init() {
	// Register model
	orm.RegisterModel(new(InspectionCriterion))
}