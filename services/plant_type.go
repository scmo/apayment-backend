package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/scmo/apayment-backend/models"
)

func CreatePlantType(pt *models.PlantType) error {
	o := orm.NewOrm()
	_, err := o.Insert(pt)
	return err
}

func GetAllPlantTypes() []*models.PlantType {
	o := orm.NewOrm()
	var plantTypes []*models.PlantType
	o.QueryTable(new(models.PlantType)).All(&plantTypes)
	return plantTypes
}

func CountPlantTypes() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(models.PlantType)).Count() // SELECT COUNT(*) FROM USE
	return cnt, err
}
