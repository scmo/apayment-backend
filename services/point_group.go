package services

import (
	"github.com/scmo/foodchain-backend/models"
	"github.com/astaxie/beego/orm"
)

func CreatePointGroup(pc *models.PointGroup) error {
	o := orm.NewOrm()
	_, err := o.Insert(pc)
	return err
}

func GetAllPointGroups() ([]*models.PointGroup) {
	o := orm.NewOrm()
	var pointGroups []*models.PointGroup
	o.QueryTable(new(models.PointGroup)).All(&pointGroups)
	for _, pointGroup := range pointGroups {
		o.LoadRelated(pointGroup, "ControlPoints")
	}

	return pointGroups
}

func CountPointGroups() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(models.PointGroup)).Count() // SELECT COUNT(*) FROM USE
	return cnt, err
}
