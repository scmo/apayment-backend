package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/services"
)

type LegalFormController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all LegalForms
// @Success 200 {Object} models.LegalForm
// @router / [get]
func (this *LegalFormController) GetAll() {
	this.Data["json"] = services.GetAllLegalForms()
	this.ServeJSON()
}
