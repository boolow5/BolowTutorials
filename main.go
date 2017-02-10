package main

import (
	"github.com/astaxie/beego"
	_ "github.com/boolow5/BolowTutorials/routers"
)

func main() {
	//beego.AddFuncMap("i18n", i18n.Tr)
	beego.AddFuncMap("eq", func(a, b interface{}) bool {
		return a == b
	})

	beego.AddFuncMap("neq", func(a, b interface{}) bool {
		return a != b
	})
	beego.AddFuncMap("not", func(a bool) bool {
		return !a
	})

	beego.Run()
}
