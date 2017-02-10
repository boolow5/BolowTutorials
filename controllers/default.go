package controllers

import (
	"fmt"
	"html/template"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	SetBasicTemplate(&this.Controller, "pages/index.tpl")
}

func (this *MainController) Detail() {
	SetBasicTemplate(&this.Controller, "pages/detail.tpl")
}

func SetBasicTemplate(this *beego.Controller, templateName string) {
	this.Layout = "layouts/base.tpl"
	this.TplName = templateName
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Head"] = "partials/head.tpl"
	this.LayoutSections["Nav"] = "partials/nav.tpl"
	this.LayoutSections["Featured"] = "partials/featured.tpl"
	this.LayoutSections["Footer"] = "partials/footer.tpl"
	this.LayoutSections["FlashMessages"] = "partials/flashes.tpl"
	this.LayoutSections["Scripts"] = "partials/scripts.tpl"

	flash := beego.ReadFromRequest(this)

	this.Data["error"] = flash.Data["error"]
	this.Data["warning"] = flash.Data["warning"]
	this.Data["notice"] = flash.Data["notice"]
	this.Data["success"] = flash.Data["success"]
	this.Data["error"] = flash.Data["error"]

	// cross-site request forgery prevention
	this.Data["xsrf_token"] = this.XSRFToken()
	this.XSRFExpire = 7200
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	this.Data["IsLoggedin"] = false
	loggedin := this.GetSession("loggedin")
	if loggedin != nil {
		this.Data["IsLoggedin"] = true
	}

	// loggedin user data
	current_user := this.GetSession("current_user")
	if current_user != nil {
		this.Data["current_user"] = current_user
	}
}

func IsAuthenticated(this *beego.Controller) bool {
	fmt.Println("IsAuthenticated")
	loggedin := this.GetSession("loggedin")

	if loggedin == nil {
		fmt.Println("No")
		return false
	}

	fmt.Println("Yes")
	return true
}
