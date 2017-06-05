package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/models"
	"encoding/json"
	"github.com/scmo/apayment-backend/services"
)

// Operations about ControlPoints
type ControlPointController struct {
	beego.Controller
}

// @Title Create a Control Point
// @Description Endpoint to create a new Control Point
// @Param	body		body 	models.ControlPoint	true		"body for control point content"
// @Success 200 {Object} models.ControlPoint
// @Failure 403 body is empty
// @router / [post]
func (this *ControlPointController) Post() {
	var controlPoint models.ControlPoint
	json.Unmarshal(this.Ctx.Input.RequestBody, &controlPoint)

	// TODO Validiation
	err := services.CreateControlPoint(&controlPoint)
	if err != nil {
		beego.Error("Create Control Point ", err.Error())
		this.CustomAbort(500, "Post Control Point Error")
	}
	this.Data["json"] = controlPoint
	this.ServeJSON()
}

// @Title GetAll
// @Description get all Control Points
// @Success 200 {Object} models.ControlPoint
// @router / [get]
func (this *ControlPointController) GetAll() {
	this.Data["json"] = services.GetAllControlPoints()
	this.ServeJSON()
}