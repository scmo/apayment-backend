package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/scmo/foodchain-backend/services"
	"github.com/scmo/foodchain-backend/models"
)

// Operations about Contributions
type InspectionCriterionController struct {
	beego.Controller
}


// @Title CreateInspectionCriterion
// @Description Endpoint to create a new Inspection Criterion
// @Param	body		body 	models.InspectionCriterion	true		"body for InspectionCriterion content"
// @Success 200 {Object} models.InspectionCriterion
// @Failure 403 body is empty
// @router / [post]
func (this *InspectionCriterionController) Post() {
	var inspectionCriterion models.InspectionCriterion
	json.Unmarshal(this.Ctx.Input.RequestBody, &inspectionCriterion)
	beego.Debug(inspectionCriterion)
	// TODO Validiation
	err := services.CreateInspectionCriterion(&inspectionCriterion)
	if err != nil {
		beego.Error("CreateInspectionCriterion ", err.Error())
		this.CustomAbort(500, "Post InspectionCriterion Error")
	}
	this.Data["json"] = inspectionCriterion
	this.ServeJSON()
}