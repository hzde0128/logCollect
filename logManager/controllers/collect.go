package controllers

import "github.com/astaxie/beego"

type CollectController struct {
	beego.Controller
}

func (c *CollectController) Get(){
	c.TplName = "collect.tpl"
}