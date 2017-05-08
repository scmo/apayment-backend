package models

import "github.com/astaxie/beego/orm"

type PointGroup struct {
	Id              int64             `json:"id"`
	PointGroupId    string            `json:"pointGroupId"`
	Abbreviation    string            `json:"abbreviation"`
	PointGroup      string            `json:"pointGroup"`
	//Condition
	//Type
	ControlCategory *ControlCategory  `orm:"rel(fk);null" json:"-"`
	ControlPoints   []*ControlPoint   `orm:"reverse(many)" json:"controlPoints"`
}

func init() {
	// Register model
	orm.RegisterModel(new(PointGroup))
}
