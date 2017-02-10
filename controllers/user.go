package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/boolow5/BolowTutorials/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) GetLogin() {
	SetBasicTemplate(&this.Controller, "pages/user/login.tpl")
}

func (this *UserController) PostLogin() {
	flash := beego.NewFlash()
	// if already loggedin go back home
	loggedin := this.GetSession("loggedin")
	fmt.Println("loggedin")
	fmt.Println(loggedin)
	if loggedin != nil {
		if bool(loggedin.(bool)) == true {
			flash.Notice("You're already logged in")
			flash.Store(&this.Controller)
			this.Redirect("/", 302)
			return
		}
	}

	user_form := models.User{}
	user_form.Username = this.GetString("username")
	user_form.Password = this.GetString("password")

	if user_form.Username == "" || user_form.Password == "" {
		flash.Error("Both username and password are required")
		flash.Store(&this.Controller)
		this.Redirect("/login", 302)
		return
	}
	loggedInUser, err := user_form.Authenticate()
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/login", 302)
		return
	}
	user_form.Password = "hidden for security reasons"
	this.SetSession("current_user", loggedInUser)
	this.SetSession("loggedin", user_form.Username)

	loggedin = this.GetSession("loggedin")
	fmt.Println("loggedin")

	fmt.Println("Congradulations! You loggedin successfully")
	flash.Success("Congradulations! You loggedin successfully")
	flash.Store(&this.Controller)
	this.Redirect("/", 302)
}

func (this *UserController) Logout() {
	flash := beego.NewFlash()
	this.DelSession("current_user")
	this.DelSession("loggedin")
	flash.Success("Successfully Logged out")
	flash.Store(&this.Controller)
	this.Redirect("/login", 302)
}

// Registration
func (this *UserController) GetRegister() {
	SetBasicTemplate(&this.Controller, "pages/user/register.tpl")
}

func (this *UserController) PostRegister() {
	fmt.Println("PostRegister:")
	flash := beego.NewFlash()
	if IsAuthenticated(&this.Controller) {
		fmt.Println("Your're already logged in")
		flash.Notice("You're already logged in")
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
		return
	}
	user_form := models.User{}
	user_form.Username = this.GetString("username")
	user_form.Password = this.GetString("password")

	fmt.Println("form data")
	fmt.Println()
	if user_form.Username == "" || user_form.Password == "" {
		fmt.Println("Both username and password are required")
		flash.Error("Both username and password are required")
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)
		return
	}
	if user_form.Exists() {
		fmt.Println("Username already exists")
		flash.Error("Username already exists")
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)
		return
	}
	fmt.Println(user_form)
	saved, err := user_form.Add()
	if err != nil || !saved {
		fmt.Println(err)
		flash.Error("Failed to register, please try again later")
		flash.Store(&this.Controller)
		this.Redirect("/register", 302)
		return
	}
	fmt.Println("Congradulations! You registered successfully")
	flash.Success("Congradulations! You registered successfully")
	flash.Store(&this.Controller)
	this.Redirect("/", 302)
}
