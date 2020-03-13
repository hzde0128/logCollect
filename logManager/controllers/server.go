package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
	"net/http"
)

type ServerController struct {
	beego.Controller
}

func (c *ServerController) Get() {
	// 显示主机列表
	o := orm.NewOrm()
	qs := o.QueryTable("Server")

	servers := []models.Server{}

	beego.Info(qs)

	_, err := qs.All(&servers)
	if err != nil {
		beego.Info("查询失败,err:", err)
	}

	c.Data["server"] = servers
	c.TplName = "server.tpl"
}

// 添加主机
func (c *ServerController) Post() {
	// 获取用户输入的数据
	serverName := c.GetString("ServerName")
	serverAddress := c.GetString("ServerAddress")
	beego.Info("ServerName:",serverName)
	beego.Info("ServerAddress:",serverAddress)

	// 简单判断
	if serverName == "" || serverAddress == "" {
		beego.Info("主机名称和地址不能为空")
		c.Redirect("/admin/server/", http.StatusFound)
	}

	// 初始orm
	o := orm.NewOrm()
	server := models.Server{}
	server.Hostname = serverName
	server.Address = serverAddress

	// 入库
	_, err := o.Insert(&server)
	if err != nil {
		beego.Info("添加失败")
	}
	// 跳转到主机列表页
	c.Redirect("/admin/server/", http.StatusFound)
}