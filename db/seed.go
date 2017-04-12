package db

import (
	"github.com/scmo/foodchain-backend/services"
	"github.com/scmo/foodchain-backend/models"
	"github.com/astaxie/beego"
)

func SeedContributions() {
	//Only seed if table is empty
	if cnt, _ := services.CountContributions(); cnt > 0 {
		return
	}
	services.CreateContribution(&models.Contribution{Code:5416, Name:"Beitrag für besonders Tierfreundliche Stallhaltungssysteme"})
	services.CreateContribution(&models.Contribution{Code:5417, Name:"Beitrag für regelmässigen Auslauf im Freien"})
}

func SeedInspectionCriterion() {
	//Only seed if table is empty
	if cnt, _ := services.CountInspectionCriteria(); cnt > 0 {
		return
	}
	contribution, err := services.GetContributionByCode(5416)
	if err != nil {
		beego.Error("Error fetching contribution.", err.Error())
	}
	services.CreateInspectionCriterion(&models.InspectionCriterion{Name:"Happy Cow", Description:"Is the a happy cow?", Contribution:contribution})
}
