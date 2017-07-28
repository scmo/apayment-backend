package models

import "github.com/astaxie/beego/orm"

type Lack struct {
	Id           int64         `json:"id"`
	Name         string        `json:"name"`
	Points       int8          `json:"points"`
	Francs       int16         `json:"francs"`
	ControlPoint *ControlPoint `orm:"rel(fk);null" json:"-"`
}

func init() {
	// Register model
	orm.RegisterModel(new(Lack))
}
