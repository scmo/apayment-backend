package controllers

import (
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/services"
)

// Operations about PointGroups
type CowController struct {
	beego.Controller
}

// @Title Get Cows from user
// @Description Get all cows of user
// @Param   Authorization     header   string true       "JWT token"
// @Param   category     query   int64 false       "Category ID"
// @router / [get]
func (this *CowController) GetCows() {
	claims, err := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	if err != nil {
		this.CustomAbort(401, "Unauthorized")
	}
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}

	categoryId, err := this.GetInt64("category")
	if err != nil {
		beego.Debug(err)
	}
	if categoryId > 0 {
		beego.Debug("get cows from category")
	}

	//Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlcyI6WyJGYXJtZXIiXSwiZXhwIjoxNDk4OTMwOTAxLCJpYXQiOjE0OTg5MTI5MDEsImlzcyI6ImxvY2FsaG9zdDo5MDAwIiwic3ViIjoiZmFybWVyMSJ9.mWvfkW2w2VHLfODMlgwJF4YPyg-E9ZIlExW31RaMaYs

	cows, err := services.GetCowFromUser(user.TVD)
	_ = cows

	this.Data["json"] = "asdf"
	this.ServeJSON()
}

// Cow Category Controll
type CategoryController struct {
	beego.Controller
}

// @Title Add new Category
// @Param   Authorization     header   string true       "JWT token"
// @router / [get]
func (this *CategoryController) GetCategories() {
	claims, err := services.ParseToken(this.Ctx.Request.Header.Get("Authorization"))
	if err != nil {
		this.CustomAbort(401, "Unauthorized")
	}
	user, err := services.GetUserByUsername(claims.Subject)
	if err != nil {
		this.CustomAbort(404, err.Error())
	}
	categories, err := services.GetCategories(user.TVD)
	this.Data["json"] = categories
	this.ServeJSON()
}

// @Title Add new Category
// @Param   Authorization     header   string true       "JWT token"
// @Param	body		body 	models.Category	true		"body for category content"
// @router / [post]
func (this *CategoryController) AddCategory() {

}

// @Title Update Category
// @Param   Authorization     header   string true       "JWT token"
// @Param	body		body 	models.Category	true		"body for category content"
// @Param   categoryId     path   int64 true       "Category ID"
// @router /:categoryId [put]
func (this *CategoryController) UpdateCategory() {

}

// @Title Delete Category
// @Param   Authorization     header   string true       "JWT token"
// @Param   categoryId     path   int64 true       "Category ID"
// @router /:categoryId [put]
func (this *CategoryController) DeleteCategory() {

}
