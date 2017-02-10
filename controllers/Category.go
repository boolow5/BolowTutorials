package controllers

import (
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["title"] = "Categories"
	SetBasicTemplate(&this.Controller, "pages/categories.tpl")
}
