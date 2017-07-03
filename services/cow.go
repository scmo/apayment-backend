package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/scmo/apayment-backend/models"
	"github.com/scmo/apayment-backend/services/tvd"
)

func GetCowFromUser(userTvd int32) ([]*models.Cow, error) {
	cows := []*models.Cow{}
	getCattleLivestockV2Response, err := tvd.GetUserCattleLivestock(userTvd)

	if err != nil {
		return cows, err
	}

	for i := 0; i < len(getCattleLivestockV2Response.GetCattleLivestockV2Result.Resultdetails.CattleLivestockDataItem); i++ {
		cow := new(models.Cow)
		cow.Tvd = getCattleLivestockV2Response.GetCattleLivestockV2Result.Resultdetails.CattleLivestockDataItem[i].EarTagNumber
		cow.Name = getCattleLivestockV2Response.GetCattleLivestockV2Result.Resultdetails.CattleLivestockDataItem[i].Name
		cow.Details = getCattleLivestockV2Response.GetCattleLivestockV2Result.Resultdetails.CattleLivestockDataItem[i]

		cows = append(cows, cow)
	}

	return cows, nil
}

func GetCategories(userTvd int32) ([]*models.Category, error) {
	categories := make([]*models.Category, 0)
	cattleLivestockV2Response, err := tvd.GetUserCattleLivestock(userTvd)
	if err != nil {
		return categories, err
	}
	for _, cattleLivestockDataV2 := range cattleLivestockV2Response.GetCattleLivestockV2Result.Resultdetails.CattleLivestockDataItem {
		cow := models.Cow{}
		cow.Name = cattleLivestockDataV2.Name
		cow.Tvd = cattleLivestockDataV2.EarTagNumber
		cow.Details = cattleLivestockDataV2
		// TODO Journal
		cow.GetLatestJournalEntry()

		categoryCode, err := tvd.GetAnimalCategory(cattleLivestockDataV2)
		if err != nil {
			beego.Critical(err.Error())
		}
		var category *models.Category
		categories, category = getCategory(categories, categoryCode)
		category.Cows = append(category.Cows, &cow)

	}
	return categories, err
}

func getCategory(categories []*models.Category, categoryCode uint8) ([]*models.Category, *models.Category) {
	for _, category := range categories {
		if category.CategoryCode == categoryCode {
			return categories, category
		}
	}
	pointGroupNames := tvd.GetPointGroupName()
	category := models.Category{Name: pointGroupNames[categoryCode-1], CategoryCode: categoryCode}
	categories = append(categories, &category)
	return categories, &category
}

func AddCategory(category *models.Category) error {
	o := orm.NewOrm()
	_, err := o.Insert(category)
	return err
}

func IsCategoryExisting(name string, is_default bool) bool {
	o := orm.NewOrm()
	category := models.Category{Name: name, IsDefault: is_default}
	err := o.Read(&category, "Name", "is_default")
	if err == orm.ErrNoRows {
		return false
	}
	return true
}
