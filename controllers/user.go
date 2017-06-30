package controllers

import (
	"encoding/json"
	"github.com/scmo/apayment-backend/models"

	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/services"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserController) Post() {
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)

	// TODO Validiation
	err := services.CreateUser(&user)
	if err != nil {
		beego.Error("CreateUser ", err.Error())
		this.CustomAbort(500, "Post User Error")
	}
	this.Data["json"] = user
	this.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {Object} models.User
// @router / [get]
func (this *UserController) GetAll() {
	claims, _ := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}
	if user.HasRole("Admin") || user.HasRole("Canton") {
		role := this.Ctx.Input.Query("role")
		if role == "Inspector" {
			this.Data["json"], err = services.GetAllUsersByRole(role)
			if err != nil {
				beego.Error("Error while fetching users by Role", err)
			}
		} else if 1 != 1 {
			// maybe another condition
		} else {
			this.Data["json"] = services.GetAllUsers()
		}

	} else {
		this.CustomAbort(401, "Unauthorized")
	}
	this.ServeJSON()
}

// @Title Get my Profile
// @Description get user based on JWT Token
// @Success 200 {object} models.User
// @Failure 404
// @router /profile [get]
func (this *UserController) Profile() {
	claims, _ := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}
	this.Data["json"] = user
	this.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (this *UserController) Get() {
	uid, err := this.GetInt64(":uid")
	if err != nil {
		beego.Error("GetInt64 ", err.Error())
		this.CustomAbort(500, "No User Id provided")

	}
	user, err := services.GetUserById(uid)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}
	this.Data["json"] = user
	this.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	// TODO
	u.Data["json"] = "User updated"
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	// TODO
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title RegisterUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /register [post]
func (this *UserController) Register() {
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)

	// TODO Validiation
	err := services.CreateUser(&user)
	if err != nil {
		beego.Error("CreateUser ", err.Error())
		this.CustomAbort(500, err.Error())
	}

	this.Data["json"] = user
	this.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (this *UserController) Login() {
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)

	user, err := services.CheckLogin(user.Username, user.Password)
	if err != nil {
		// TODO: Return appropiate Errer, when 'no row found'
		beego.Error("CheckLogin ", err.Error())
		this.CustomAbort(500, "Login Error")
	}

	signedToken := services.IssueToken(&user)
	//user.JwtToken = signedToken
	beego.Debug(signedToken)
	this.Data["json"] = signedToken
	this.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
