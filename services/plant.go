package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/scmo/apayment-backend/models"
)

func CreatePlant(r *models.Plant) error {
	o := orm.NewOrm()
	_, err := o.Insert(r)
	return err
}
