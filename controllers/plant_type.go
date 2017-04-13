package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scmo/foodchain-backend/services"
)

type PlantTypeController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all PlantTypes
// @Success 200 {Object} models.PlantType
// @router / [get]
func (this *PlantTypeController) GetAll() {
	this.Data["json"] = services.GetAllPlantTypes()
	this.ServeJSON()
}