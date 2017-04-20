package db

import (
	"github.com/scmo/foodchain-backend/services"
	"github.com/scmo/foodchain-backend/models"
)

func Seed_Contributions() {
	//Only seed if table is empty
	if cnt, _ := services.CountContributions(); cnt > 0 {
		return
	}
	services.CreateContribution(&models.Contribution{Code:5416, Name:"Beitrag für besonders Tierfreundliche Stallhaltungssysteme"})
	services.CreateContribution(&models.Contribution{Code:5417, Name:"Beitrag für regelmässigen Auslauf im Freien"})
}

func Seed_ControlPoints() {

	//Only seed if table is empty
	if cnt, _ := services.CountControlCategories(); cnt > 0 {
		return
	}
	if cnt, _ := services.CountPointGroups(); cnt > 0 {
		return
	}
	if cnt, _ := services.CountControlPoints(); cnt > 0 {
		return
	}
	Seed_BTS_Rindergattung_Wasserbueffel()
}

func Seed_BTS_Rindergattung_Wasserbueffel() {
	cc := models.ControlCategory{CategoryId:"12.01_2017", Abbreviation:"BTS - Rindergattung und Wasserbüffel"}
	services.CreateControlCategory(&cc)

	pg1 := models.PointGroup{PointGroupId:"A1", Abbreviation:"Rinder - Milchkühe", ControlCategory: &cc}
	services.CreatePointGroup(&pg1)

	cp1 := models.ControlPoint{ControlPointId:"01", Abbreviation:"Alle Tiere frei in Gruppen gehalten", ControlPoint:"Alle Tiere der Kategorie in Gruppen gehalten oder zulässige Abweichungen gemäss DZV Anhang 6, A, 1.4", PointGroup: &pg1}
	services.CreateControlPoint(&cp1)
	cp2 := models.ControlPoint{ControlPointId:"02", Abbreviation:"Mind. 15 Lux Tageslicht im Stall", ControlPoint:"Alle Ställe, in denen sich die Tiere überwiegend aufhalten, verfügen über Tageslicht von mindestens 15 Lux Stärke (Kunstlicht zur Beurteilung ausschalten!). In Ruhe- und Rückugsbereichen ist eine geringere Beleuchtung zulässig.", PointGroup: &pg1}
	services.CreateControlPoint(&cp2)

}

func Seed_Users() {
	if cnt, _ := services.CountRoles(); cnt > 0 {
		return
	}
	admin := models.Role{Name:"Admin"}
	farmer := models.Role{Name:"Farmer"}
	canton := models.Role{Name:"Canton"}
	inspector := models.Role{Name:"Inspector"}
	bund := models.Role{Name:"bund"}

	services.CreateRole(&admin);
	services.CreateRole(&farmer);
	services.CreateRole(&canton);
	services.CreateRole(&inspector);
	services.CreateRole(&bund);

	if cnt, _ := services.CountUsers(); cnt > 0 {
		return
	}

	roles := make([]*models.Role, 1)
	roles[0] = &farmer
	services.CreateUser(&models.User{Username:"farmer1", Password:"farmer1", Firstname: "Florian", Lastname:"Meisterhans", Roles: roles})
	services.CreateUser(&models.User{Username:"farmer2", Password:"farmer2", Firstname: "Max", Lastname:"Keller", Roles: roles})

	roles = make([]*models.Role, 1)
	roles[0] = &inspector
	services.CreateUser(&models.User{Username:"inspect", Password:"inspect", Firstname: "Inspector", Lastname: "Gadget", Roles: roles})

	roles = make([]*models.Role, 1)
	roles[0] = &admin
	services.CreateUser(&models.User{Username:"admin", Password:"admin", Firstname: "Admin", Lastname: "Admin", Roles: roles})

}

func Seed_LegalForm() {
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

func Seed_PlantType() {
	if cnt, _ := services.CountPlantTypes(); cnt > 0 {
		return
	}
	services.CreatePlantType(&models.PlantType{Code: 1, Name:"Ganzjahresbetrieb"})
	services.CreatePlantType(&models.PlantType{Code: 2, Name:"Produktsstätte"})
	services.CreatePlantType(&models.PlantType{Code: 3, Name:"Betriebsgemeinschaft"})
}
