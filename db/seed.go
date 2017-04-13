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

func SeedLegalForm() {
	if cnt, _ := services.CountLegalForms(); cnt > 0 {
		return
	}
	services.CreateLegalForm(&models.LegalForm{Code: 1, Name:"Natürtliche Person"})
	services.CreateLegalForm(&models.LegalForm{Code: 2, Name:"Einfache Gesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 3, Name:"Kollektivgesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 4, Name:"Kommanditgesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 5, Name:"Kommanditaktiengesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 6, Name:"Aktiengesellschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 7, Name:"Gesellschaft mit beschränkter Haftung"})
	services.CreateLegalForm(&models.LegalForm{Code: 8, Name:"Genossenschaft"})
	services.CreateLegalForm(&models.LegalForm{Code: 9, Name:"Verein, Vereinigung"})
	services.CreateLegalForm(&models.LegalForm{Code: 10, Name:"Stiftung"})

	services.CreateLegalForm(&models.LegalForm{Code: 24, Name:"Öffentlich-rechtliche Körperschaft(Verwaltung)"})
	services.CreateLegalForm(&models.LegalForm{Code: 25, Name:"Staatlich anerkannte Landeskirche"})
	services.CreateLegalForm(&models.LegalForm{Code: 30, Name:"Bund (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 31, Name:"Kanton (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 32, Name:"Bezirk (Betrieb)"})

	services.CreateLegalForm(&models.LegalForm{Code: 33, Name:"Gemeinde (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 34, Name:"Öffentlich-rechtliche Körperschaft (Betrieb)"})
	services.CreateLegalForm(&models.LegalForm{Code: 99, Name:"Nicht zugeteilt"})
}

func SeedPlantType(){
	if cnt, _ := services.CountPlantTypes(); cnt > 0 {
		return
	}
	services.CreatePlantType(&models.PlantType{Code: 1, Name:"Ganzjahresbetrieb"})
	services.CreatePlantType(&models.PlantType{Code: 2, Name:"Produktsstätte"})
	services.CreatePlantType(&models.PlantType{Code: 3, Name:"Betriebsgemeinschaft"})
}
