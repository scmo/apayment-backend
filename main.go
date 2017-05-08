package main

import (
	_ "github.com/scmo/foodchain-backend/routers"
	"github.com/astaxie/beego"
	"github.com/scmo/foodchain-backend/db"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/scmo/foodchain-backend/ethereum"

)

func init() {
	// TODO: ethereum struct
	ethereum.Init()
	// Setup DB
	db.Init()


}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//beego.InsertFilter("/*", beego.BeforeRouter, routers.HandleJWT)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.Run()
}




