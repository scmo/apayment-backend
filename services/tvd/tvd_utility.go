package tvd

import (
	"errors"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

func GetPointGroupCodes() []uint16 {
	return []uint16{1110, 1150, 1128, 1141, 1142, 1124, 1129, 1143, 1144}
}

func GetPointGroupName() []string {
	return []string{
		"Milchkühe",
		"andere Kühe",
		"weibliche Tiere über 365 - 730 Tage alt, ohne Abkalbung",
		"weibliche Tiere über 160 - 365 Tage alt",
		"weibliche Tiere bis 160 Tage alt",
		"männliche Tiere, über 730 Tage alt",
		"männliche Tiere, über 365 bis 730 Tage alt",
		"männliche Tiere, über 160 bis 365 Tage alt",
		"männliche Tiere, bis 160 Tage alt",
	}
}

func GetUserCattleLivestock(userTvd int32, begin time.Time, end time.Time) (*GetCattleLivestockV2Response, error) {
	animalTracingPortType := NewAnimalTracingPortType("https://ws-in.wbf.admin.ch/Livestock/AnimalTracing/1", true, getAuth())

	workingFocus := []WorkingFocus{}
	workingFocus = append(workingFocus, WorkingFocus{
		WorkingFocusType: 3,
		TVDNumber:        userTvd,
		MandateGiver:     beego.AppConfig.String("agate_username"),
	})

	workingFocusArray := WorkingFocusArray{
		WorkingFocusItem: &workingFocus,
	}

	Request := GetCattleLivestockV2{
		Action:           "http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1/AnimalTracingPortType/GetCattleLivestockV2",
		PManufacturerKey: beego.AppConfig.String("tvd_manufacturerKey"),
		PTVDNumber:       userTvd,
		PLCID:            2055,
		PSearchDateFrom:  begin,
		PSearchDateTo:    end,
		PWorkingFocus:    &workingFocusArray,
	}

	cattleLivestockV2Response, err := animalTracingPortType.GetCattleLivestockV2(&Request)
	if err != nil {
		return nil, err
	}
	return cattleLivestockV2Response, nil
}

func getFieldTestGVE() (map[uint16]uint32, error) {
	return map[uint16]uint32{
		1110: 11,
		1150: 12,
		1128: 13,
		1141: 14,
		1142: 15,
		1124: 16,
		1129: 17,
		1143: 18,
		1144: 19,
	}, nil
}

/*
	Calculates the GVE from previous year.
*/
func GetNumberOfGVELastYear(userTvd int32) (map[uint16]uint32, error) {
	// FOR FIELD TEST
	if userTvd == 0 {
		return getFieldTestGVE()
	}
	a1 := float32(0) // a1 1110    Milchkühe
	a2 := float32(0) // a2 1150   andere Kühe
	a3 := float32(0) // a3 1128    weibliche Tiere, über 365 Tage alt, bis zur ersten Abkalbung,
	a4 := float32(0) // a4 1141    weibliche Tiere über 160 - 365 Tage alt
	a5 := float32(0) // a5 1142   weibliche Tiere bis 160 Tage alt (nur RAUS)
	a6 := float32(0) // a6 1124   männliche Tiere, über 730 Tage alt
	a7 := float32(0) // a7 1129   männliche Tiere, über 365 bis 730 Tage alt
	a8 := float32(0) // a8 1143   männliche Tiere, über 160 bis 365 Tage alt
	a9 := float32(0) // a9 1144   männliche Tiere, bis 160 Tage alt (nur RAUS)

	oc, _ := time.LoadLocation("Europe/Zurich")
	begin := time.Date(time.Now().Year()-1, time.Month(1), 0, 0, 0, 0, 0, oc)
	end := time.Date(time.Now().Year()-1, time.Month(12), 30, 23, 59, 0, 0, oc)

	cattleLivestockV2Response, err := GetUserCattleLivestock(userTvd, begin, end)
	if err != nil {
		beego.Error("Error while fetching CattleLiveStockV2: ", err)
		// When TVD doesnt work, just return dummy values
		return getFieldTestGVE()
	}
	for _, cattleLiveStockDataItem := range cattleLivestockV2Response.GetCattleLivestockV2Result.Resultdetails.CattleLivestockDataItem {
		cat, err := GetAnimalCategory(cattleLiveStockDataItem)
		if err != nil {
			return nil, err
		}
		days, err := cattleLiveStockDataItem.getStayLengthInDays()

		gve := float32(days) / 365
		switch cat {
		case 1:
			a1 += gve
		case 2:
			a2 += gve
		case 3:
			a3 += gve
		case 4:
			a4 += gve
		case 5:
			a5 += gve
		case 6:
			a6 += gve
		case 7:
			a7 += gve
		case 8:
			a8 += gve
		case 9:
			a9 += gve
		default:
			beego.Error("No category defined")
		}
	}
	// multiply by 10'000 to allow 4 decimals
	return map[uint16]uint32{
		1110: uint32(a1 * 10000),
		1150: uint32(a2 * 10000),
		1128: uint32(a3 * 10000),
		1141: uint32(a4 * 10000),
		1142: uint32(a5 * 10000),
		1124: uint32(a6 * 10000),
		1129: uint32(a7 * 10000),
		1143: uint32(a8 * 10000),
		1144: uint32(a9 * 10000),
	}, nil
}


func GetAnimalCategory(cattleLiveStockDataItem *CattleLivestockDataV2) (uint8, error) {
	ageInDays, err := getAgeInDays(cattleLiveStockDataItem)
	if err != nil {
		return 0, err
	}

	genderCode, err := strconv.Atoi(cattleLiveStockDataItem.Gender)
	if err != nil {
		beego.Error("Error while parsing Gender: ", err)
	}
	// female = 1; male = 2
	if genderCode == 1 {
		// female
		if len(cattleLiveStockDataItem.FirstCalvingDate) > 0 {
			// Ohne abkalberung
			_, err := time.Parse(beego.AppConfig.String("tvd_timeformat"), cattleLiveStockDataItem.FirstCalvingDate)
			if err != nil {
				beego.Error("Error while parsing lastCalvingDate: ", err)
				return 0, err
			}
			return 1, nil // A3
		}

		if ageInDays <= 160 {
			return 5, nil // a5
		}
		if 160 < ageInDays && ageInDays <= 365 {
			return 4, nil // a4
		}
		if 365 < ageInDays {
			return 3, nil // a3
		}
	} else if genderCode == 2 {
		// male
		if ageInDays <= 160 {
			return 9, nil // a5
		}
		if 160 < ageInDays && ageInDays <= 365 {
			return 8, nil // a5
		}
		if 365 < ageInDays && ageInDays <= 730 {
			return 7, nil // a5
		}
		if 730 < ageInDays {
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

func (cattleLiveStockDataItem *CattleLivestockDataV2) getStayLengthInDays() (uint16, error) {
	layout := "2006-01-02T00:00:00"
	arrival, err := time.Parse(layout, cattleLiveStockDataItem.ArrivalDate)
	if err != nil {
		beego.Error("Error while parsing date.", err)
		return 0, err
	}
	oc, _ := time.LoadLocation("Europe/Zurich")
	leaving := time.Date(time.Now().Year()-1, time.Month(12), 30, 23, 59, 0, 0, oc)
	if cattleLiveStockDataItem.LeavingDate != "" {
		leaving, err = time.Parse(layout, cattleLiveStockDataItem.LeavingDate)
		if err != nil {
			beego.Error("Error while parsing date.", err)
			return 0, err
		}
	}

	duration := leaving.Sub(arrival).Hours()
	return uint16(duration / 24), nil
}

func GetPersonAddressFromTVD() (*PersonAddressResult, error) {
	animalTracingPortType := NewAnimalTracingPortType(beego.AppConfig.String("animalTracingURL"), true, getAuth())

	request := GetPersonAddress{
		Action:           "http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1/AnimalTracingPortType/GetPersonAddress",
		PManufacturerKey: beego.AppConfig.String("tvd_manufacturerKey"),
		PLCID:            2055,
		PAgateNumber:     beego.AppConfig.String("agate_username"),
	}
	response, err := animalTracingPortType.GetPersonAddress(&request)
	if err != nil {
		return nil, err
	}
	return response.GetPersonAddressResult, err
}

func GetAnimalHusbandryDetailFromTVD(userTVD int32) (*GetAnimalHusbandryDetailResult, error) {
	animalTracingPortType := NewAnimalTracingPortType(beego.AppConfig.String("animalTracingURL"), true, getAuth())
	getAnimalHusbandryDetailRequest := new(GetAnimalHusbandryDetailRequest)
	getAnimalHusbandryDetailRequest.BaseRequest = getBaseRequest()
	getAnimalHusbandryDetailRequest.TVDNumber = userTVD

	request := GetAnimalHusbandryDetail{
		PGetAnimalHusbandryDetailRequest: getAnimalHusbandryDetailRequest,
	}
	response, err := animalTracingPortType.GetAnimalHusbandryDetail(&request)
	if err != nil {
		return nil, err
	}
	return response.GetAnimalHusbandryDetailResult, err
}

func getAuth() *BasicAuth {
	return &BasicAuth{
		Login:    beego.AppConfig.String("agate_username"),
		Password: beego.AppConfig.String("agate_password"),
	}
}
func getBaseRequest() *BaseRequest {
	return &BaseRequest{
		LCID:            2055,
		ManufacturerKey: beego.AppConfig.String("tvd_manufacturerKey"),
	}
}
