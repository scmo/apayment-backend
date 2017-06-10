package tvd

import (
	"github.com/astaxie/beego"
	"time"
	"strconv"
	"errors"
)

func GetBTSPointGroupCodes() []uint16 {
	return []uint16{1110, 1150, 1123, 1128, 1141, 1142, 1124, 1129, 1143, 1144};
}

func GetUserCattleLivestock(userTvd int32, agateUsername string, agatePassword string) (*GetCattleLivestockV2Response, error) {
	auth := BasicAuth{
		Login: agateUsername,
		Password: agatePassword,
	}
	animalTracingPortType := NewAnimalTracingPortType("https://ws-in.wbf.admin.ch/Livestock/AnimalTracing/1", true, &auth)

	workingFocus := []WorkingFocus{}
	workingFocus = append(workingFocus, WorkingFocus{
		WorkingFocusType: 3,
		TVDNumber: userTvd,
		MandateGiver: agateUsername,
	})

	workingFocusArray := WorkingFocusArray{
		WorkingFocusItem: &workingFocus,
	}

	Request := GetCattleLivestockV2{
		Action: "http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1/AnimalTracingPortType/GetCattleLivestockV2",
		PManufacturerKey: beego.AppConfig.String("tvd_manufacturerKey"),
		PTVDNumber: userTvd,
		PLCID: 2055,
		PSearchDateFrom: time.Now().AddDate(0, 0, -1),
		PSearchDateTo: time.Now().AddDate(0, 0, 0),
		PWorkingFocus: &workingFocusArray,
	}

	cattleLivestockV2Response, err := animalTracingPortType.GetCattleLivestockV2(&Request)
	if err != nil {
		return nil, err
	}
	return cattleLivestockV2Response, nil
}

func GetNumberOfGVE(userTvd int32, agateUsername string, agatePassword string) (map[uint16]uint16, error) {

	a1 := 0 // a1 1110    Milchkühe
	a2 := 0 // a2 1150   andere Kühe
	a3 := 0 // a3 1128    weibliche Tiere, über 365 Tage alt, bis zur ersten Abkalbung,
	a4 := 0 // a4 1141    weibliche Tiere über 160 - 365 Tage alt
	a5 := 0 // a5 1142   weibliche Tiere bis 160 Tage alt (nur RAUS)
	a6 := 0 // a6 1124   männliche Tiere, über 730 Tage alt
	a7 := 0 // a7 1129   männliche Tiere, über 365 bis 730 Tage alt
	a8 := 0 // a8 1143   männliche Tiere, über 160 bis 365 Tage alt
	a9 := 0 // a9 1144   männliche Tiere, bis 160 Tage alt (nur RAUS)

	cattleLivestockV2Response, err := GetUserCattleLivestock(userTvd, agateUsername, agatePassword)
	if (err != nil) {
		beego.Error("Error while fetching CattleLiveStockV2: ", err)
	}
	for _, cattleLiveStockDataItem := range cattleLivestockV2Response.GetCattleLivestockV2Result.Resultdetails.CattleLivestockDataItem {
		//beego.Debug(cattleLiveStockDataItem)
		cat, err := GetAnimalCategory(cattleLiveStockDataItem)
		if (err != nil) {
			return nil, err
		}
		switch cat {
		case 1:
			a1++
		case 2:
			a2++
		case 3:
			a3++
		case 4:
			a4++
		case 5:
			a5++
		case 6:
			a6++
		case 7:
			a7++
		case 8:
			a8++
		case 9:
			a9++
		default:
			beego.Error("No category defined")
		}
	}
	return map[uint16]uint16{
		1110: uint16(a1),
		1150: uint16(a2),
		1128: uint16(a3),
		1141: uint16(a4),
		1142: uint16(a5),
		1124: uint16(a6),
		1129: uint16(a7),
		1143: uint16(a8),
		1144: uint16(a9),
	}, nil
};

func GetAnimalCategory(cattleLiveStockDataItem *CattleLivestockDataV2) (uint8, error) {
	ageInDays, err := getAgeInDays(cattleLiveStockDataItem)
	if (err != nil ) {
		return 0, err
	}

	genderCode, err := strconv.Atoi(cattleLiveStockDataItem.Gender)
	if (err != nil) {
		beego.Error("Error while parsing Gender: ", err)
	}
	// female = 1; male = 2
	if (genderCode == 1 ) {
		// female
		if ( len(cattleLiveStockDataItem.FirstCalvingDate) > 0) {
			// Ohne abkalberung
			_, err := time.Parse(beego.AppConfig.String("tvd_timeformat"), cattleLiveStockDataItem.FirstCalvingDate)
			if (err != nil ) {
				beego.Error("Error while parsing lastCalvingDate: ", err)
				return 0, err
			}
			return 1, nil // A3
		}

		if (ageInDays <= 160) {
			return 5, nil // a5
		}
		if (160 < ageInDays && ageInDays <= 365) {
			return 4, nil // a4
		}
		if (365 < ageInDays) {
			return 3, nil // a3
		}
	} else if (genderCode == 2 ) {
		// male
		if (ageInDays <= 160) {
			return 9, nil // a5
		}
		if (160 < ageInDays && ageInDays <= 365) {
			return 8, nil // a5
		}
		if (365 < ageInDays && ageInDays <= 730) {
			return 7, nil // a5
		}
		if (730 < ageInDays) {
			return 6, nil // a5
		}
	}
	return 0, errors.New("No Gender specified")
}

func getAgeInDays(cattleLiveStockDataItem *CattleLivestockDataV2) (uint32, error) {
	birthdate, err := time.Parse(beego.AppConfig.String("tvd_timeformat"), cattleLiveStockDataItem.BirthDate)
	if err != nil {
		beego.Error("Error while converting birthdate to days: ", err)
		return 0, err
	}
	return uint32(time.Now().Sub(birthdate).Hours() / 24), nil
}