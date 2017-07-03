package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/scmo/apayment-backend/models"
	"github.com/scmo/apayment-backend/services/tvd"
	"strconv"
	"time"
)

func AddJournalEntry(tvds []int32, date time.Time, minutesOutside int, typeOfFieldLairage string) error {
	o := orm.NewOrm()
	for _, tvd := range tvds {
		//defaultCategory, err := getDefaultCategory(tvd)
		journalEntry := models.JournalEntry{TVD: tvd, Date: date, MinutesOutside: minutesOutside}
		_, err := o.Insert(journalEntry)
		if err != nil {
			beego.Error("Error while adding JournalEntry")
			return err
		}
	}
	return nil
}

func GetJournal(tvd string) (*models.Journal, error) {

	journal := new(models.Journal)

	return journal, nil
}

func GetMonthlyStats(userTvd int32, month int, year int) (models.MonthlyStats, error) {
	monthlyStats := models.MonthlyStats{Month: month, Year: year}

	cattleLivestockV2Response, err := tvd.GetUserCattleLivestock(userTvd)
	if err != nil {
		return monthlyStats, err
	}

	for _, cattleLivestockDataV2 := range cattleLivestockV2Response.GetCattleLivestockV2Result.Resultdetails.CattleLivestockDataItem {

		cowStats := models.CowStats{
			EarTagNumber: cattleLivestockDataV2.EarTagNumber,
			Name:         cattleLivestockDataV2.Name,
		}

		categoryCode, err := tvd.GetAnimalCategory(cattleLivestockDataV2)
		if err != nil {
			beego.Critical(err.Error())
		}

		categoryStats := addOrGetCategoryStats(&monthlyStats, categoryCode)

		// TODO: total Days outside

		categoryStats.Cows = append(categoryStats.Cows, &cowStats)

		//monthlyStats.CategoryStats = append(monthlyStats.CategoryStats, categoryStats)
	}

	beego.Debug("done")
	return monthlyStats, nil
}

func addOrGetCategoryStats(monthlyStats *models.MonthlyStats, categoryCode uint8) *models.CategoryStats {
	for _, categoryStats := range monthlyStats.CategoryStats {
		if categoryStats.CategoryCode == categoryCode {
			return categoryStats
		}
	}
	categoryStats := models.CategoryStats{CategoryCode: categoryCode, CategoryName: "A" + strconv.FormatUint(uint64(categoryCode), 8)}
	monthlyStats.CategoryStats = append(monthlyStats.CategoryStats, &categoryStats)
	return &categoryStats
}
