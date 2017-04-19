package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scmo/foodchain-backend/models"
	"encoding/json"
	"github.com/scmo/foodchain-backend/services"
)

// Operations about ControlCategories
type PointGroupController struct {
	beego.Controller
}

// @Title Create a Point Group
// @Description Endpoint to create a new Point Group
// @Param	body		body 	models.PointGroup	true		"body for point group content"
// @Success 200 {Object} models.PointGroup
// @Failure 403 body is empty
// @router / [post]
func (this *PointGroupController) Post() {
	var pointGroup models.PointGroup
	json.Unmarshal(this.Ctx.Input.RequestBody, &pointGroup)

	// TODO Validiation
	err := services.CreatePointGroup(&pointGroup)
	if err != nil {
		beego.Error("Create Point Group ", err.Error())
		this.CustomAbort(500, "Post Point Group Error")
	}
	this.Data["json"] = pointGroup
	this.ServeJSON()
}

// @Title GetAll
// @Description get all Control Categories
// @Success 200 {Object} models.ControlCategory
// @router / [get]
func (this *PointGroupController) GetAll() {
	this.Data["json"] = services.GetAllPointGroups()
	this.ServeJSON()
}