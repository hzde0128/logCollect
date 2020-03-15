package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
)

// ServerController 主机结构体
type ServerController struct {
	beego.Controller
}

// DeleteSvrController 删除主机
type DeleteSvrController struct {
	beego.Controller
}

// Get 显示主机列表
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
	pageSize := 5
	query := o.QueryTable("Server")
	count, err := query.Count()

	// 获取页面数，向上取整
	page, err := strconv.Atoi(c.GetString("page"))
	if err != nil {
		page = 1
	}
	start := pageSize * (page - 1)

	pageCount := int(math.Ceil(float64(count) / float64(pageSize)))
	table := o.QueryTable("Server")
	table.Limit(pageSize, start).All(&servers)

	c.Data["pageCount"] = pageCount
	c.Data["count"] = count
	c.Data["pagesize"] = pageSize
	c.Data["server"] = servers
	c.Data["prepage"] = 1
	c.Data["page"] = page

	c.Layout = "layout.tpl"
	c.TplName = "server.tpl"
}

// Post 添加主机
func (c *ServerController) Post() {
	// 获取用户输入的数据
	serverName := c.GetString("ServerName")
	serverAddress := c.GetString("ServerAddress")
	beego.Info("ServerName:", serverName)
	beego.Info("ServerAddress:", serverAddress)

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

// Get 删除主机
func (c *DeleteSvrController) Get() {
	// 获取用户输入的ID
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("数据非法")
		return
	}
	beego.Info("获取到的ID：", id)

	// 执行删除操作
	o := orm.NewOrm()
	server := models.Server{Id: id}
	_, err = o.Delete(&server)
	if err != nil {
		beego.Info("删除失败，", err)
		return
	}

	c.Redirect("/admin/server/", http.StatusFound)
}
