package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/scmo/apayment-backend/services/tvd"
)

type Cow struct {
	Id                 int64                      `json:"id"`
	Name               string                     `json:"name"`
	TVD                string                     ` json:"tvd"`
	Details            *tvd.CattleLivestockDataV2 `orm:"-" json:"details"`
	Journal            []*JournalEntry            `orm:"-;reverse(many)" json:"journal"` // WHY IS THIS NEEDED?
	Farm_id            string                     `orm:"-" json:"farm_id,omitempty"`
	Categories         []*Category                `orm:"-" json:"categories,omitempty"`
	Added              string                     `orm:"-" json:"added,omitempty"`
	MoreDetails        *tvd.CattleDetailData      `orm:"-" json:"moreDetails,omitempty"`
	LatestJournalEntry *JournalEntry              `orm:"-" json:"latest_journal_entry,omitempty"` // WHY IS THIS NEEDED?
}

type Category struct {
	CategoryCode uint8  `json:"-"`
	Name         string `json:"category"`
	Added        string `json:"added,omitempty"`
	IsDefault    bool   `json:"isDefault"`
	Cows         []*Cow `json:"cows,omitempty"`
}

func init() {
	// Register model
	orm.RegisterModel(new(Cow))
}

/*	FUNCTIONS	 */

func (cow *Cow) GetJournal() {
	o := orm.NewOrm()
	var journalEntries []*JournalEntry
	_, err := o.QueryTable(new(JournalEntry)).Filter("TVDS__Cow__TVD", cow.TVD).OrderBy("-date").All(&journalEntries)
	if err != nil {
		beego.Error("Could not load JournalEntries.", err.Error())
	}
	cow.Journal = journalEntries
	if len(journalEntries) > 0 {
		cow.LatestJournalEntry = journalEntries[0]
	}
}
