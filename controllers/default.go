package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "htime.me"
	this.Data["Email"] = "nemowen@gmail.com"
	this.TplNames = "index.html"

}
