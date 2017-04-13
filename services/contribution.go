package services

import (
	"github.com/scmo/foodchain-backend/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func CreateContribution(c *models.Contribution) error {
	o := orm.NewOrm()
	_, err := o.Insert(c)
	return err
}

func GetAllContributions() ([]*models.Contribution) {
	o := orm.NewOrm()
	var contributions []*models.Contribution
	o.QueryTable(new(models.Contribution)).All(&contributions)
	for _, contribution := range contributions {
		o.LoadRelated(contribution, "InspectionCriteria")
	}
	return contributions
}

func GetContributionById(_id int64) (*models.Contribution, error) {
	o := orm.NewOrm()
	contribution := models.Contribution{Id: _id}
	err := o.Read(&contribution)
	if err == orm.ErrNoRows {
		beego.Error("No result found.")
		return nil, err
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
		return nil, err
	}
	o.LoadRelated(&contribution, "InspectionCriteria")
	return &contribution, nil
}

func GetContributionByCode(_code uint16) (*models.Contribution, error) {
	o := orm.NewOrm()
	contribution := models.Contribution{Code: _code}
	err := o.Read(&contribution, "Code")
	if err == orm.ErrNoRows {
		beego.Error("No result found.")
		return nil, err
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
		return nil, err
	}
	o.LoadRelated(&contribution, "InspectionCriteria")
	return &contribution, nil
}

func CountContributions() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(models.Contribution)).Count() // SELECT COUNT(*) FROM USE
	return cnt, err
}
