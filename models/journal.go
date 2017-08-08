package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Journal struct {
	JournalEntries []*JournalEntry
}

type JournalEntry struct {
	Id             int64     `json:"-"`
	TVDS           []*Cow    `orm:"rel(m2m)" json:"tvds"`
	Year           uint16    `orm:"-" json:"year"`
	Month          uint8     `orm:"-" json:"month"`
	Day            uint8     `orm:"-" json:"day"`
	Date           time.Time `orm:"auto_now_add;type(date)" json:"date"`
	MinutesOutside int       `json:"minutes_outside"`
	TypeOfLairage  string    `json:"type_of_field_lairage"`
	Created        time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
	Updated        time.Time `orm:"auto_now;type(datetime)"`
}

type MonthlyStats struct {
	Month         uint8            `json:"month"`
	Year          uint16           `json:"year"`
	CategoryStats []*CategoryStats `json:"category_stats"`
}

type CategoryStats struct {
	NumberOfDaysRequired int         `json:"number_of_days_required"`
	NumberOfAnimals      int         `json:"number_of_animals"`
	LowestDays           int16       `json:"lowest_days"`
	MaximumDays          int16       `json:"maximum_days"`
	CategoryName         string      `json:"category_name"`
	CategoryDescription  string      `json:"category_description"`
	CategoryCode         uint8       `json:"-"`
	CowsStats            []*CowStats `json:"cows"`
}

type CowStats struct {
	EarTagNumber        string `json:"ear_tag_number"`
	NumberOfDaysOutside int16  `json:"number_of_days_outside"`
	TotalMinutesOutside int    `json:"total_minutes_outside"`
	Name                string `json:"name"`
}

type RausJournalRequest struct {
	TVDs []string `json:"tvds"`
	Year int16 `json:"year"`
}

type RausJournalResponse struct {
	Valid bool `json:"valid"`
	MissedDays int16 `json:"missed_days"`
	NumberOfInvalidCows int16 `json:"num_of_invalid_cows"`
	year int16 `json:"year"`
}

func init() {
	// Register model
	orm.RegisterModel(new(JournalEntry))
}

func (je *JournalEntry) TableName() string {
	return "journal"
}

/*	FUNCTIONS	 */

func (journalEntry *JournalEntry) SetDate() {
	oc, _ := time.LoadLocation("Europe/Zurich")
	journalEntry.Date = time.Date(int(journalEntry.Year), time.Month(journalEntry.Month), int(journalEntry.Day), 0, 0, 0, 0, oc)
}

func (cowStats *CowStats) SetDayOutside(month uint8, year uint16) {
	o := orm.NewOrm()
	var journalEntries []*JournalEntry

	oc, _ := time.LoadLocation("Europe/Zurich")
	startDate := time.Date(int(year), time.Month(month), 1, 0, 0, 0, 0, oc)
	endDate := time.Date(int(year), time.Month(month), 31, 0, 0, 0, 0, oc)
	num, err := o.QueryTable(new(JournalEntry)).Filter("TVDS__Cow__TVD", cowStats.EarTagNumber).Filter("date__gte", startDate).Filter("date__lte", endDate).OrderBy("-date").All(&journalEntries)
	if err != nil {
		beego.Error("Could not load JournalEntries.", err.Error())
	}
	for _, entry := range journalEntries {
		cowStats.TotalMinutesOutside += entry.MinutesOutside
	}
	cowStats.NumberOfDaysOutside = int16(num)
}

func (monthlyStats *MonthlyStats) SetCategoryMinMaxDays() {
	for _, categoryStats := range monthlyStats.CategoryStats {
		categoryStats.LowestDays = 27
		for _, cowStats := range categoryStats.CowsStats {
			if cowStats.NumberOfDaysOutside < categoryStats.LowestDays {
				categoryStats.LowestDays = cowStats.NumberOfDaysOutside
			}
			if cowStats.NumberOfDaysOutside > categoryStats.MaximumDays {
				categoryStats.MaximumDays = cowStats.NumberOfDaysOutside
			}
		}
	}
}
