// @APIVersion 1.0.0
// @Title Foodchain API
// @Description Autogenerated API Documentation of the foodchain backend
// @Contact moritz.schneider3@uzh.ch
// @TermsOfServiceUrl http://foodchain.ch
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/scmo/apayment-backend/controllers"

	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/context"
	"github.com/scmo/apayment-backend/services"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/request",
			beego.NSInclude(
				&controllers.RequestController{},
			),
		),
		beego.NSNamespace("/controlcategory",
			beego.NSInclude(
				&controllers.ControlCategoryController{},
			),
		),
		beego.NSNamespace("/pointgroup",
			beego.NSInclude(
				&controllers.PointGroupController{},
			),
		),
		beego.NSNamespace("/controlpoint",
			beego.NSInclude(
				&controllers.ControlPointController{},
			),
		),
		beego.NSNamespace("/lack",
			beego.NSInclude(
				&controllers.LackController{},
			),
		),
		beego.NSNamespace("/contribution",
			beego.NSInclude(
				&controllers.ContributionController{},
			),
		),
		beego.NSNamespace("/legalform",
			beego.NSInclude(
				&controllers.LegalFormController{},
			),
		),
		beego.NSNamespace("/planttype",
			beego.NSInclude(
				&controllers.PlantTypeController{},
			),
		),
		beego.NSNamespace("/ping",
			beego.NSInclude(
				&controllers.PingController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

var HandleJWT = func(ctx *context.Context) {
	if strings.Compare(ctx.Request.Method, "OPTIONS") == 0 {
		return
	}
	if strings.HasPrefix(ctx.Input.URL(), "/v1/user/login") {
		return
	}
	if services.Validate(ctx.Request.Header.Get("Authorization")) {
		return
	}
	ctx.Abort(401, "unauthorized")
}





