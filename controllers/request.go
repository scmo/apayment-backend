package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scmo/foodchain-backend/models"
	"encoding/json"
	"github.com/scmo/foodchain-backend/services"
)

// Operations about Contributions
type RequestController struct {
	beego.Controller
}

// @Title Create a new Request
// @Description Endpoint to create a new Request
// @Param	body		body 	models.Request	true		"body for request content"
// @Success 200 {Object} models.Request
// @Failure 403 body is empty
// @router / [post]
func (this *RequestController) Post() {
	var request models.Request
	json.Unmarshal(this.Ctx.Input.RequestBody, &request)

	claims, _ := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	user, err := services.GetUserByUsername(claims.Username)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}
	request.User = user

	services.CreateRequest(&request)
	//// TODO Validiation
	this.ServeJSON()
}

// @Title Get
// @Description find request by requestID
// @Param	requestId		path 	string	true		"the requestid you want to get"
// @Success 200 {object} models.Request
// @Failure 403 :requestId is empty
// @router /:requestId [get]
func (this *RequestController) Get() {
	objectId := this.Ctx.Input.Param(":requestId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = ob
		}
	}
	this.ServeJSON()
}

// @Title GetAll
// @Description get all request
// @Success 200 {object} models.Request
// @router / [get]
func (this *RequestController) GetAll() {
	requests := services.GetAllRequests()
	this.Data["json"] = requests
	this.ServeJSON()
}
