package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.TplName = "home.html"
	//c.TplName = "login.html"
	c.TplName = "show.html"
}
