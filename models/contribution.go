package models

import "github.com/astaxie/beego/orm"

type Contribution struct {
	Id             uint32 `json:"id"`
	Code           uint16 `json:"code"`
	Name           string `json:"name"`
	InspectionCriteria []*InspectionCriterion `json:"inspectioncriteria" orm:"reverse(many)"`
}

func init() {
	// Register model
	orm.RegisterModel(new(Contribution))
}