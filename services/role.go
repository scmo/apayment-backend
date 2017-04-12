package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/scmo/foodchain-backend/models"
)

func CreateRole(r *models.Role) error  {
	o := orm.NewOrm()
	_, err := o.Insert(r)
	return err
}


