package models

import (
	"github.com/scmo/apayment-backend/services/tvd"
)

type Cow struct {
	Name               string                     `json:"name"`
	Tvd                string                     `json:"tvd"`
	Details            *tvd.CattleLivestockDataV2 `json:"details"`
	Journal            *Journal                   `json:"journal"`
	Farm_id            string                     `json:"farm_id,omitempty"`
	Categories         []*Category                `json:"categories,omitempty"`
	Added              string                     `json:"added,omitempty"`
	MoreDetails        *tvd.CattleDetailData      `json:"moreDetails,omitempty"`
	LatestJournalEntry *JournalEntry              `json:"latest_journal_entry,omitempty"`
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
	//orm.RegisterModel(new(Category), new(CowTVD))
}

/*	FUNCTIONS	 */

func (cow *Cow) GetLatestJournalEntry() {
	cow.LatestJournalEntry = &JournalEntry{}
}
