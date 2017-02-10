package routers

import (
	"github.com/astaxie/beego"
	"github.com/boolow5/BolowTutorials/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/detail", &controllers.MainController{}, "get:Detail")
	beego.Router("/login", &controllers.UserController{}, "get:GetLogin;post:PostLogin")
	beego.Router("/logout", &controllers.UserController{}, "*:Logout")
	beego.Router("/register", &controllers.UserController{}, "get:GetRegister;post:PostRegister")
}
