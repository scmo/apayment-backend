package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/scmo/apayment-backend/models"
)

func CreateControlCategory(cc *models.ControlCategory) error {
	o := orm.NewOrm()
	_, err := o.Insert(cc)
	return err
}

func GetAllControlCategories() []*models.ControlCategory {
	o := orm.NewOrm()
	var controlCategories []*models.ControlCategory
	o.QueryTable(new(models.ControlCategory)).All(&controlCategories)
	for _, controlCategory := range controlCategories {
		o.LoadRelated(controlCategory, "PointGroups")
		for _, pointGroup := range controlCategory.PointGroups {
			beego.Debug(pointGroup.Abbreviation)
			o.LoadRelated(pointGroup, "ControlPoints")
		}
	}

	return controlCategories
}

func CountControlCategories() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(models.ControlCategory)).Count() // SELECT COUNT(*) FROM USE
	return cnt, err
}
