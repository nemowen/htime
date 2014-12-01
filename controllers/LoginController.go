package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {

	this.Data["last_release_date"] = beego.AppConfig.String("last_release_date")
	this.Data["static_server"] = beego.AppConfig.String("static_server")

	this.TplNames = "admin/login.html"

}
