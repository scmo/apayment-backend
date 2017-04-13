package services

import (
	"github.com/scmo/foodchain-backend/models"
	"github.com/astaxie/beego"
)

func CreateRequest(r *models.Request) error {
	//o := orm.NewOrm()
	//_, err := o.Insert(c)
	//return err
	beego.Debug(r.Id)
	for i := range r.Contributions {
		beego.Debug(r.Contributions[i].Name)
	}
	beego.Debug(r.User.Username)

	return nil
}
