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
	//err := services.CreateContribution(&contribution)
	//if err != nil {
	//	beego.Error("CreateContribution ", err.Error())
	//	this.CustomAbort(500, "Post Contribution Error")
	//}
	//this.Data["json"] = contribution
	this.ServeJSON()
}
