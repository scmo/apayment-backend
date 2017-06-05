package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/scmo/apayment-backend/models"
)

func CreateRole(r *models.Role) error  {
	o := orm.NewOrm()
	_, err := o.Insert(r)
	return err
}

func CountRoles() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(models.Role)).Count() // SELECT COUNT(*) FROM USE
	return cnt, err
}

