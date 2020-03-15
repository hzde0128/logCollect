package controllers

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
)

// UpdateController 更新
type UpdateController struct {
	beego.Controller
}

// Get 修改日志收集项
func (c *UpdateController) Get() {
	// 从数据库获取当前的信息
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	beego.Info("Id: ", id)

	// 验证id是否非法
	o := orm.NewOrm()
	collect := models.Collect{}
	collect.Id = id
	err := o.Read(&collect, "id")
	if err != nil {
		beego.Info("无效的ID", err)
		return
	}

	o.QueryTable("Collect").RelatedSel("Server").Filter("id", id).All(&collect)

	// 显示内容
	c.Data["collect"] = collect

	c.Layout = "layout.tpl"
	c.TplName = "update.tpl"
}

// Post 处理修改日志收集项
func (c *UpdateController) Post() {
	server := c.GetString("server")
	path := c.GetString("path")
	topic := c.GetString("topic")
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	beego.Info("Path: ", path)
	beego.Info("Topic: ", topic)
	beego.Info("Server: ", server)

	// 修改数据库
	o := orm.NewOrm()
	collect := models.Collect{}
	servers := models.Server{}
	servers.Address = server
	err := o.Read(&servers, "Address")
	if err != nil {
		beego.Info("获取主机错误")
		return
	}
	collect.Topic = topic
	collect.Path = path
	collect.Server = &servers
	collect.Id = id

	_, err = o.Update(&collect)
	if err != nil {
		beego.Info("修改失败,", err)
		return
	}
	c.Redirect("/admin/", http.StatusFound)
}
