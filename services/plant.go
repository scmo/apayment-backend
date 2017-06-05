package services

import (
	"github.com/scmo/apayment-backend/models"
	"github.com/astaxie/beego/orm"
)

func CreatePlant(r *models.Plant) error  {
	o := orm.NewOrm()
	_, err := o.Insert(r)
	return err
}
