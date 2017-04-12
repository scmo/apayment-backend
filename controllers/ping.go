package controllers

import (
	"github.com/astaxie/beego"
)

type PingController struct {
	beego.Controller
}

// @Title Ping
// @Success 200 {string} pong
// @router / [get]
func (this *PingController) Ping() {
	this.Data["json"] = "pong"
	this.ServeJSON()
}
