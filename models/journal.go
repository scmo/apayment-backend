package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Journal struct {
	JournalEntries []*JournalEntry
}

type JournalEntry struct {
	TVD            int32     `orm:"pk" json:"tvd"`
	Date           time.Time `orm:"auto_now_add;type(date)" json:"date"`
	MinutesOutside int       `json:"minutes_outside"`
	TypeOfLairage  string    `json:"type_of_field_lairage"`
	Created        time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
	Updated        time.Time `orm:"auto_now;type(datetime)"`
}

type MonthlyStats struct {
	Month         int              `json:"month"`
	Year          int              `json:"year"`
	CategoryStats []*CategoryStats `json:"category_stats"`
}

type CategoryStats struct {
	NumberOfDaysRequired int         `json:"number_of_days_required"`
	NumberOfAnimals      int         `json:"number_of_animals"`
	LowestDays           int         `json:"lowest_days"`
	MaximumDays          int         `json:"maximum_days"`
	CategoryName         string      `json:"category_name"`
	CategoryDescription  string      `json:"category_description"`
	CategoryCode         uint8       `json:"-"`
	Cows                 []*CowStats `json:"cows"`
}

type CowStats struct {
	EarTagNumber        string `json:"ear_tag_number"`
	NumberOfDaysOutside int    `json:"number_of_days_outside"`
	TotalMinutesOutside int    `json:"total_minutes_outside"`
	Name                string `json:"name"`
}

func init() {
	// Register model
	orm.RegisterModel(new(JournalEntry))
}

func (je *JournalEntry) TableName() string {
	return "journal"
}
