package services

import (
	"github.com/scmo/foodchain-backend/models"
	"github.com/astaxie/beego/orm"
)

func CreateControlPoint(cp *models.ControlPoint) error  {
	o := orm.NewOrm()
	_, err := o.Insert(cp)
	return err
}

func GetAllControlPoints() ([]*models.ControlPoint) {
	o := orm.NewOrm()
	var controlPoints []*models.ControlPoint
	o.QueryTable(new(models.ControlPoint)).All(&controlPoints)
	//for _, contribution := range contributions {
	//	o.LoadRelated(contribution, "InspectionCriteria")
	//}
	return controlPoints
}

func CountControlPoints() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(models.ControlPoint)).Count() // SELECT COUNT(*) FROM USE
	return cnt, err
}
