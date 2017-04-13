package models

import "github.com/astaxie/beego/orm"

type InspectionCriterion struct {
	Id           int64 `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Contribution *Contribution `orm:"rel(fk)"`
}

func init() {
	// Register model
	orm.RegisterModel(new(InspectionCriterion))
}