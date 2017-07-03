package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/models"
	"github.com/scmo/apayment-backend/services"
	"time"
)

// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlcyI6WyJGYXJtZXIiXSwiZXhwIjoxNTM1MDk4OTc4LCJpYXQiOjE0OTkxMDI1NzgsImlzcyI6ImxvY2FsaG9zdDo5MDAwIiwic3ViIjoiZmFybWVyMSJ9.HDN_TowtegvrblHCgC-Fp_AnkuYZmX82LApl_HipvkM

// Operations about PointGroups
type JournalController struct {
	beego.Controller
}

func (controller *JournalController) getUser() *models.User {
	claims, err := services.ParseToken(controller.Ctx.Request.Header.Get("Authorization"))
	if err != nil {
		controller.CustomAbort(401, "Unauthorized")
	}
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		controller.CustomAbort(404, err.Error())
	}
	return user
}

// @Title Get Monthly Stat
// @Description Get monthly statistics
// @Param   Authorization     header   string true       "JWT token"
// @Param   month     query   int8 true       "month"
// @Param   year     query   int8 true       "year"
// @router /monthlystats [get]
func (this *JournalController) GetMonthlyStats() {

	user := this.getUser()

	month, err := this.GetInt("month")
	if err != nil {
		this.CustomAbort(400, "Month not sent")
	}
	year, err := this.GetInt("year")
	if err != nil {
		this.CustomAbort(400, "Year not sent")
	}

	if month < 1 || month > 12 || year < 1900 || year > time.Now().Year() {
		this.CustomAbort(400, "Value not possible for one of the parameter. Year must be a 4 digit integer. Month parameter must be a interger between 1 and 31.")
	}

	monthlyStats, err := services.GetMonthlyStats(user.TVD, month, year)
	if err != nil {
		this.CustomAbort(501, "Internal Error "+err.Error())
	}
	this.Data["json"] = monthlyStats
	this.ServeJSON()
}

// @Title Add Entry Journal
// @Description Add new Entry to the Journal
// @Param   Authorization     header   string true       "JWT token"
// @Param   journalEntry     body   model.JournalEntry true       "JournalEntry"
// @router / [post]
func (this *JournalController) AddJournalEntry() {
	user := this.getUser()

	var journalEntry models.JournalEntry
	json.Unmarshal(this.Ctx.Input.RequestBody, &journalEntry)

	beego.Debug(journalEntry)

	this.Data["json"] = user
	this.ServeJSON()
}
