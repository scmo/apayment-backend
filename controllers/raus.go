package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/services"
	"time"
)

// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlcyI6WyJGYXJtZXIiXSwiZXhwIjoxNTM1MDc5MTk2LCJpYXQiOjE0OTkwODI3OTYsImlzcyI6ImxvY2FsaG9zdDo5MDAwIiwic3ViIjoiZmFybWVyMSJ9.HNlNd1NI8Oop0uS74FzbRE4Q-PqaH3nW5tCU82-VPP0

// Operations about PointGroups
type RausController struct {
	beego.Controller
}

// @Title Get Monthly Stat
// @Description Get monthly statistics
// @Param   Authorization     header   string true       "JWT token"
// @Param   month     query   int8 true       "month"
// @Param   year     query   int8 true       "year"
// @router /monthlystats [get]
func (this *RausController) GetMonthlyStats() {

	claims, err := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	if err != nil {
		this.CustomAbort(401, "Unauthorized")
	}
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}

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
// @Param   journalEntry     body   model.JournalEntry true       "JournalEntry"
// @router /journal [post]
func (this *RausController) AddJournalEntry() {

}
