package models

import "github.com/astaxie/beego/orm"

type Plant struct {
	Id                int64      `json:"id"`
	CantonalPlantNr   string     `json:"cantonalPlantNr"` // e.g. 1234/70/123
	Community         string     `json:"community"`       // Deutsch: Gemeinde
	CommunityNr       uint16     `json:"communityNr"`
	Name              string     `json:"name"`
	XCoordinate       uint32     `json:"xCoordinate"`
	YCoordinate       uint32     `json:"xCoordinate"`
	PlantType         *PlantType `orm:"rel(fk);null" json:"plantType"`
	MetersAboveTheSea int16      `json:"metersAboveTheSea"`
							      // TODO:
							      // Wenn Betriebsgemeinschaft, Anzahl Mitgliedsbetreibe eintragen
							      // Mitglied einer Betriebszweiggemeinschft? Ja/Nein
							      // Mitglied einer ÖLN Gemeinschaft? Ja/ Nein
							      // Werden mehr als 50% der Arbeiten auf dem Betrieb durch eigene Arbeitskraefte durchgefurht? Ja/ Nein
}

func init() {
	// Register model
	orm.RegisterModel(new(Plant))
}
