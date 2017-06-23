package services

import (
	"github.com/scmo/apayment-backend/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
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

func GetAllPointGroupById(_id int64) (*models.PointGroup, error) {
	o := orm.NewOrm()
	pointGroup := models.PointGroup{Id: _id}
	err := o.Read(&pointGroup)
	if err == orm.ErrNoRows {
		beego.Error("No result found.")
		return nil, err
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
		return nil, err
	}

	return &pointGroup, nil
}

func GetAllPointGroupByCode(_code uint16) (*models.PointGroup, error) {
	o := orm.NewOrm()
	pointGroup := models.PointGroup{PointGroupCode: _code}
	err := o.Read(&pointGroup, "PointGroupCode")
	if err == orm.ErrNoRows {
		beego.Error("No result found. ProintGroupCode: ", _code)
		return nil, err
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
		return nil, err
	}
	return &pointGroup, err
}

func CountPointGroups() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(models.PointGroup)).Count() // SELECT COUNT(*) FROM USE
	return cnt, err
}
