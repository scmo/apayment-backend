package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/scmo/apayment-backend/models"
	"github.com/scmo/apayment-backend/services"
)

// Operations about Contributions
type ContributionController struct {
	beego.Controller
}

// @Title CreateContribution
// @Description Endpoint to create a new Contribution
// @Param	body		body 	models.Contribution	true		"body for contribution content"
// @Success 200 {Object} models.Contribution
// @Failure 403 body is empty
// @router / [post]
func (this *ContributionController) Post() {
	var contribution models.Contribution
	json.Unmarshal(this.Ctx.Input.RequestBody, &contribution)

	// TODO Validiation
	err := services.CreateContribution(&contribution)
	if err != nil {
		beego.Error("CreateContribution ", err.Error())
		this.CustomAbort(500, "Post Contribution Error")
	}
	this.Data["json"] = contribution
	this.ServeJSON()
}

// @Title Get
// @Description get contribution by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {Object} models.Contribution
// @Failure 403 :cid is empty
// @router /:cid [get]
func (this *ContributionController) Get() {
	cid, err := this.GetInt64(":cid")
	if err != nil {
		beego.Error("GetInt64 ", err.Error())
		this.CustomAbort(500, "No Contribution Id provided")
	}
	contribution, err := services.GetContributionById(cid)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}
	this.Data["json"] = contribution
	this.ServeJSON()
}

// @Title GetAll
// @Description get all Contributions
// @Success 200 {Object} models.Contribution
// @router / [get]
func (this *ContributionController) GetAll() {
	this.Data["json"] = services.GetAllContributions()
	this.ServeJSON()
}
