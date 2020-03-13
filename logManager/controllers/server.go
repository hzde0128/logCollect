package controllers

import "github.com/astaxie/beego"

type ServerController struct {
	beego.Controller
}

func (c *ServerController) Get() {
	c.TplName = "server.tpl"
}