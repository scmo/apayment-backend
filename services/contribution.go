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
		o.LoadRelated(contribution, "ControlCategories")
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
	o.LoadRelated(&contribution, "ControlCategories")
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
	o.LoadRelated(&contribution, "ControlCategories")
	for _, controlCategory := range contribution.ControlCategories {
		o.LoadRelated(controlCategory, "PointGroups")
		for _, pointgroup := range controlCategory.PointGroups {
			o.LoadRelated(pointgroup, "ControlPoints")
			for _, controlPoint := range pointgroup.ControlPoints {
				o.LoadRelated(controlPoint, "Lacks")
			}
		}
	}
	return &contribution, nil
}

func CountContributions() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(models.Contribution)).Count() // SELECT COUNT(*) FROM USE
	return cnt, err
}

func GetContributionByInspectionLack(iL *models.InspectionLack) *models.Contribution {
	o := orm.NewOrm()
	contribution := models.Contribution{Code:iL.ContributionCode}
	err := o.Read(&contribution, "Code")
	if err != nil {
		beego.Error("Error while loading Contribution. (ContributionCode: ", iL.ContributionCode, ") Error:", err)
		return &contribution
	}
	controlCategory := models.ControlCategory{ControlCategoryId:iL.ControlCategoryId, Contribution: &contribution}
	err = o.Read(&controlCategory, "ControlCategoryId")
	if err != nil {
		beego.Error("Error while loading ControlCategory. (ControlCategoryId: ", iL.ControlCategoryId, ") Error:", err)
		return &contribution
	}
	contribution.ControlCategories = append(contribution.ControlCategories, &controlCategory)

	pointGroup := models.PointGroup{PointGroupId:iL.PointGroupId, ControlCategory: &controlCategory}
	err = o.Read(&pointGroup, "PointGroupId")
	if err != nil {
		beego.Error("Error while loading PointGroup. (PointGroupId: ", iL.PointGroupId, ") Error:", err)
		return &contribution
	}
	controlCategory.PointGroups = append(controlCategory.PointGroups, &pointGroup)

	controlPoint := models.ControlPoint{ControlPointId: iL.ControlPointId, PointGroup: &pointGroup}
	err = o.Read(&controlPoint, "ControlPointId")
	if err != nil {
		beego.Error("Error while loading ControlPoint. (ControlPointId: ", iL.ControlPointId, ") Error:", err)
		return &contribution
	}
	pointGroup.ControlPoints = append(pointGroup.ControlPoints, &controlPoint)

	lack := models.Lack{Id:iL.LackId, ControlPoint:&controlPoint}
	err = o.Read(&lack)
	if err != nil {
		beego.Error("Error while loading Lack. (Lackid:", iL.LackId, ") Error:", err)
		return &contribution
	}
	controlPoint.Lacks = append(controlPoint.Lacks, &lack)

	return &contribution
}