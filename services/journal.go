package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/scmo/apayment-backend/models"
	"github.com/scmo/apayment-backend/services/tvd"
	"strconv"
)

func AddJournalEntry(journalEntry *models.JournalEntry) error {
	o := orm.NewOrm()
	_, err := o.Insert(journalEntry)

	if err != nil {
		beego.Error("Insert JournalEntry. ", err.Error())
		return err
	}

	m2m := o.QueryM2M(journalEntry, "TVDS")
	for _, cow := range journalEntry.TVDS {
		if _, id, err := o.ReadOrCreate(cow, "TVD"); err == nil {
			cow.Id = id
		} else {
			beego.Error("ReadOrCreate by adding TVD to Journal ", err.Error())
			return err
		}
		_, err := m2m.Add(cow)
		if err != nil {
			beego.Error("Many2Many Add ", err.Error())
			return err
		}
	}

	return nil
}

func GetJournal(tvd string) (*models.Journal, error) {

	journal := new(models.Journal)

	return journal, nil
}

func GetMonthlyStats(userTvd int32, month uint8, year uint16) (models.MonthlyStats, error) {
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
		cowStats.SetDayOutside(month, year)
		categoryCode, err := tvd.GetAnimalCategory(cattleLivestockDataV2)
		if err != nil {
			beego.Critical(err.Error())
		}

		categoryStats := addOrGetCategoryStats(&monthlyStats, categoryCode)

		// TODO: total Days outside

		categoryStats.CowsStats = append(categoryStats.CowsStats, &cowStats)
		categoryStats.NumberOfAnimals++
	}
	monthlyStats.SetCategoryMinMaxDays()
	return monthlyStats, nil
}

func addOrGetCategoryStats(monthlyStats *models.MonthlyStats, categoryCode uint8) *models.CategoryStats {
	for _, categoryStats := range monthlyStats.CategoryStats {
		if categoryStats.CategoryCode == categoryCode {
			return categoryStats
		}
	}
	categoryStats := models.CategoryStats{CategoryCode: categoryCode, CategoryName: "A" + strconv.FormatUint(uint64(categoryCode), 16), CategoryDescription: tvd.GetPointGroupName()[categoryCode-1], NumberOfDaysRequired: 27}
	monthlyStats.CategoryStats = append(monthlyStats.CategoryStats, &categoryStats)
	return &categoryStats
}
