package services

import (
	"github.com/scmo/apayment-backend/models"
	"github.com/astaxie/beego/orm"
)

func CreateLegalForm(lf *models.LegalForm) error {
	o := orm.NewOrm()
	_, err := o.Insert(lf)
	return err
}

func GetAllLegalForms() ([]*models.LegalForm) {
	o := orm.NewOrm()
	var legalforms []*models.LegalForm
	o.QueryTable(new(models.LegalForm)).All(&legalforms)
	return legalforms
}

func CountLegalForms() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(models.LegalForm)).Count() // SELECT COUNT(*) FROM USE
	return cnt, err
}