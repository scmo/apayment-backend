package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/models"
	"github.com/scmo/apayment-backend/services"
)

// Operations about ControlCategories
type ControlCategoryController struct {
	beego.Controller
}

// @Title Create a Control Category
// @Description Endpoint to create a new Control Category
// @Param	body		body 	models.ControlCategory	true		"body for control category content"
// @Success 200 {Object} models.ControlCategory
// @Failure 403 body is empty
// @router / [post]
func (this *ControlCategoryController) Post() {
	var controlCategory models.ControlCategory
	json.Unmarshal(this.Ctx.Input.RequestBody, &controlCategory)

	// TODO Validiation
	err := services.CreateControlCategory(&controlCategory)
	if err != nil {
		beego.Error("CreateContribution ", err.Error())
		this.CustomAbort(500, "Post Contribution Error")
	}
	this.Data["json"] = controlCategory
	this.ServeJSON()
}

// @Title GetAll
// @Description get all Control Categories
// @Success 200 {Object} models.ControlCategory
// @router / [get]
func (this *ControlCategoryController) GetAll() {
	this.Data["json"] = services.GetAllControlCategories()
	this.ServeJSON()
}
