// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/scmo/foodchain-backend/controllers"

	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/context"
	"github.com/scmo/foodchain-backend/services"
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
		beego.NSNamespace("/contribution",
			beego.NSInclude(
				&controllers.ContributionController{},
			),
		),
		beego.NSNamespace("/inspectioncriterion",
			beego.NSInclude(
				&controllers.InspectionCriterionController{},
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
	if strings.HasPrefix(ctx.Input.URL(), "/v1/user/login") {
		return
	}
	if services.Validate(ctx.Request.Header.Get("Authorization")) {
		return
	}
	ctx.Abort(401, "unauthorized")
}





