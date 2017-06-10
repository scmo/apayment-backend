package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:cid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlCategoryController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlCategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlCategoryController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlCategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlPointController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlPointController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlPointController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlPointController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LackController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LackController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LackController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LackController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LegalFormController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LegalFormController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PingController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PingController"],
		beego.ControllerComments{
			Method: "Ping",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PlantTypeController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PlantTypeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PointGroupController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PointGroupController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PointGroupController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PointGroupController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:requestId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "GetAllForInspection",
			Router: `/inspection`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "AddInspector",
			Router: `/inspector`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "AddInspection",
			Router: `/inspection`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "UpdateGVE",
			Router: `/gve`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Profile",
			Router: `/profile`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "FullProfile",
			Router: `/fullprofile`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
