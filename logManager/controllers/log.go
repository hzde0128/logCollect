package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
	"net/http"
)

type LogController struct {
	beego.Controller
}

// 获取日志列表
func (c *LogController) Get(){

	o := orm.NewOrm()

	// 获取主机列表
	qs := o.QueryTable("Server")

	servers := []models.Server{}

	beego.Info(qs)

	_, err := qs.All(&servers)
	if err != nil {
		beego.Info("查询失败,err:", err)
	}

	// 获取日志收集项
	var collects []models.Collect
	table := o.QueryTable("collect")
	table.Limit(5, 0).RelatedSel("Server").All(&collects)

	c.Data["collect"] = collects
	c.Data["server"] = servers

	c.TplName = "index.tpl"
}

func (c *LogController) Post() {
	// 接收数据
	address := c.GetString("select")
	beego.Info("Address: ", address)
	if address == "" {
		beego.Info("下拉框传递数据失败")
	}
	// 获取分类信息
	c.Redirect("/admin/", http.StatusFound)
}