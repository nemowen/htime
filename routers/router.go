package routers

import (
	"github.com/astaxie/beego"
	"github.com/nemowen/htime/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/signup", &controllers.LoginController{}, "get:SignupPage;post:Signup")
}
