package services

import (
	"github.com/scmo/foodchain-backend/models"
	"github.com/astaxie/beego/orm"
)

func CreateInspectionCriterion(ic *models.InspectionCriterion) error {
	o := orm.NewOrm()
	_, err := o.Insert(ic)
	return err
}

func CountInspectionCriteria() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(models.InspectionCriterion)).Count() // SELECT COUNT(*) FROM USE
	return cnt, err
}
