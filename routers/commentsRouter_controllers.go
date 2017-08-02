package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:APaymentTokenController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:APaymentTokenController"],
		beego.ControllerComments{
			Method: "Transfer",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:APaymentTokenController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:APaymentTokenController"],
		beego.ControllerComments{
			Method: "GetAllTransactions",
			Router: `/transactions`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetCategories",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "AddCategory",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "UpdateCategory",
			Router: `/:categoryId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "DeleteCategory",
			Router: `/:categoryId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ContributionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:cid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlCategoryController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlCategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlCategoryController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlCategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlPointController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlPointController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlPointController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ControlPointController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:CowController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:CowController"],
		beego.ControllerComments{
			Method: "GetCows",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:JournalController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:JournalController"],
		beego.ControllerComments{
			Method: "AddJournalEntry",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:JournalController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:JournalController"],
		beego.ControllerComments{
			Method: "GetMonthlyStats",
			Router: `/monthlystats`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LackController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LackController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LackController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LackController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LegalFormController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:LegalFormController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PingController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PingController"],
		beego.ControllerComments{
			Method: "Ping",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PlantTypeController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PlantTypeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PointGroupController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PointGroupController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PointGroupController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:PointGroupController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:requestId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "UpdateGVE",
			Router: `/gve`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "GetAllForInspection",
			Router: `/inspection`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "AddInspection",
			Router: `/inspection`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "AddInspector",
			Router: `/inspector`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:RequestController"],
		beego.ControllerComments{
			Method: "Pay",
			Router: `/pay`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Profile",
			Router: `/profile`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/apayment-backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
