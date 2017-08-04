package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "消息队列-WEB"
	c.Data["Email"] = "learninton@gmail.com"
	c.TplName = "index.html"
}
