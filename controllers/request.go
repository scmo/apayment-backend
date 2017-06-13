package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/models"
	"encoding/json"
	"github.com/scmo/apayment-backend/services"
	"strconv"
	"github.com/scmo/apayment-backend/ethereum"
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
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}
	request.User = user

	err = services.CreateRequest(&request, ethereum.GetAuth(user.Address))
	if err != nil {
		this.CustomAbort(500, err.Error())
	}

	// Update GVE for the request
	err = services.SetGVE(&request)
	if (err != nil ) {
		this.CustomAbort(500, err.Error())
	}
	this.ServeJSON()
}

// @Title Get
// @Description find request by requestID
// @Param jwtToken header string true "jwt Token for Authorization"
// @Param	requestId		path 	string	true		"the requestid you want to get"
// @Success 200 {object} models.Request
// @Failure 403 :requestId is empty
// @router /:requestId [get]
func (this *RequestController) Get() {
	input := this.Ctx.Input.Param(":requestId")
	requestId, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	this.Data["json"] = services.GetRequestById(requestId)
	this.ServeJSON()
}

// @Title GetAll
// @Description get all request
// @Success 200 {object} models.Request
// @router / [get]
func (this *RequestController) GetAll() {
	requests := []*models.Request{}

	claims, _ := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}

	if (user.HasRole("Farmer")) {
		requests = services.GetAllRequestsByUserId(user.Id)
	} else if ( user.HasRole("Admin") || user.HasRole("Canton")) {
		requests = services.GetAllRequests()
	} else {
		this.CustomAbort(401, "Unauthorized")
	}

	this.Data["json"] = requests
	this.ServeJSON()
}

// @Title GetAll
// @Description get all request which have an inspector assigned
// @Success 200 {object} models.Request
// @router /inspection [get]
func (this *RequestController) GetAllForInspection() {
	requests := []*models.Request{}
	claims, _ := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}

	if (user.HasRole("Inspector")) {
		requests = services.GetAllRequestsForInspectionByInspectorId(user.Id)
	} else if ( user.HasRole("Admin") || user.HasRole("Canton")) {
		requests = services.GetAllRequestsForInspection()
	} else {
		this.CustomAbort(401, "Unauthorized")
	}

	this.Data["json"] = requests
	this.ServeJSON()
}


// @Title Add Inspector
// @Description add Inspector to Requestion
// @Param	body		body 	models.Request	true		"body for requestion content"
// @Success 200 {object} models.Request
// @router /inspector [put]
func (this *RequestController) AddInspector() {
	var request models.Request
	json.Unmarshal(this.Ctx.Input.RequestBody, &request)

	claims, _ := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}

	if ( user.HasRole("Admin") || user.HasRole("Canton")) {
		//requests = services.GetAllRequests()
		services.AddInspectorToRequest(&request, ethereum.GetAuth(user.Address))
	} else {
		this.CustomAbort(401, "Unauthorized")
	}

	this.Data["json"] = request
	this.ServeJSON()
}

// @Title Add Inspection
// @Description Add the report of the inspection
// @Param	body		body 	models.Request	true		"body for requestion content"
// @Success 200 {object} models.Request
// @router /inspection [post]
func (this *RequestController) AddInspection() {
	var inspection models.Inspection
	json.Unmarshal(this.Ctx.Input.RequestBody, &inspection)

	claims, _ := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}
	if ( user.HasRole("Admin") || user.HasRole("Inspector")) {
		//inspection.InspectorId = user.Id
	} else {
		this.CustomAbort(401, "Unauthorized")
	}
	err = services.AddLacksToRequest(&inspection, ethereum.GetAuth(user.Address))
	if err != nil {
		this.CustomAbort(500, err.Error())
	}
	this.Data["json"] = inspection
	this.ServeJSON()
}

// @Title Update GVE
// @Description Update GVE of request
// @Param	body		body 	models.Request	true		"body for requestion content"
// @Success 200 {object} models.Request
// @Failure 403 :requestId is empty
// @router /gve [put]
func (this *RequestController) UpdateGVE() {
	var request models.Request

	json.Unmarshal(this.Ctx.Input.RequestBody, &request)

	claims, _ := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}

	if ( ( user.HasRole("Admin") || user.HasRole("Canton") ) || user.Address == request.User.Address) {
		err = services.SetGVE(&request)
		if (err != nil ) {
			this.CustomAbort(500, err.Error())
		}
	}

	this.Data["json"] = request
	this.ServeJSON()
}


// @Title Pay DirectPayment
// @Description Update GVE of request
// @Param	body		body 	models.Request	true		"body for requestion content"
// @Success 200 {object} models.Request
// @router /pay [post]
func (this *RequestController) Pay() {

	var r models.Request
	json.Unmarshal(this.Ctx.Input.RequestBody, &r)

	claims, _ := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}

	if ( ( user.HasRole("Admin") || user.HasRole("Canton") ) == false ) {
		this.CustomAbort(401, "Unauthorized")
	}

	request := services.GetRequestById(r.Id)

	beego.Debug("lets make the payment")
	//if ( len(request.Payments) == 0 ) {
	//	beego.Debug("make first payment")
	//} else if ( len(request.Payments) == 1 ) {
	//	beego.Debug("make second payment")
	//} else if ( len(request.Payments) == 2 ) {
	//	beego.Debug("make third payment")
	//}
	this.Data["json"] = request
	this.ServeJSON()
}