package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
	"net/http"
)

type CollectController struct {
	beego.Controller
}

func (c *CollectController) Get(){
	// 查询主机信息
	o := orm.NewOrm()
	qs := o.QueryTable("Server")

	var servers []models.Server
	_, err := qs.All(&servers)
	if err != nil {
		beego.Info("查询失败,err:", err)
	} else {
		c.Data["server"] = servers
	}

	c.Layout = "layout.tpl"
	c.TplName = "collect.tpl"
}

// 添加日志收集处理
func (c *CollectController) Post(){
	// 处理用户发过来的请求
	server := c.GetString("server")
	filepath := c.GetString("filePath")
	topic := c.GetString("topic")
	beego.Info("Server:", server)
	beego.Info("FilePath:", filepath)
	beego.Info("Topic:", topic)

	// 处理处理
	if filepath == "" || topic == "" {
		beego.Info("路径或者Topic为空")
		c.Redirect("/admin/collect", http.StatusFound)
	}

	// 数据入库
	o := orm.NewOrm()


	var servers models.Server
	collect := models.Collect{}

	servers.Address = server
	err := o.Read(&servers, "Address")
	if err != nil{
		beego.Info("非法的服务器地址")
		return
	}
	collect.Server = &servers
	collect.Path = filepath
	collect.Topic = topic

	_, err = o.Insert(&collect)
	if err != nil{
		beego.Info("添加失败,", err)
		return
	}

	c.Redirect("/admin/", http.StatusFound)
}
