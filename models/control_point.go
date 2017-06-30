package models

import "github.com/astaxie/beego/orm"

type ControlPoint struct {
	Id             int64       `json:"id"`
	ControlPointId string      `json:"controlPointId"`
	Abbreviation   string      `json:"abbreviation"`
	ControlPoint   string      `json:"controlPoint"`
	Lacks          []*Lack     `orm:"reverse(many)" json:"lacks"`
	PointGroup     *PointGroup `orm:"rel(fk);null" json:"-"`
}

func init() {
	// Register model
	orm.RegisterModel(new(ControlPoint))
}
