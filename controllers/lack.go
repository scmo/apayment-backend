package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/models"
	"github.com/scmo/apayment-backend/services"
)

// Operations about Lacks
type LackController struct {
	beego.Controller
}

// @Title Create a Lack
// @Description Endpoint to create a new Lack
// @Param	body		body 	models.Lack	true		"body for lackcontent"
// @Success 200 {Object} models.Lack
// @Failure 403 body is empty
// @router / [post]
func (this *LackController) Post() {
	var lack models.Lack
	json.Unmarshal(this.Ctx.Input.RequestBody, &lack)

	// TODO Validiation
	err := services.CreateLack(&lack)
	if err != nil {
		beego.Error("Create Lack ", err.Error())
		this.CustomAbort(500, "Post Lack Error")
	}
	this.Data["json"] = lack
	this.ServeJSON()
}

// @Title GetAll
// @Description get all Lacks
// @Success 200 {Object} models.Lack
// @router / [get]
func (this *LackController) GetAll() {
	this.Data["json"] = services.GetAllLacks()
	this.ServeJSON()
}
