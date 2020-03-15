package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
)

// LogController 日志
type LogController struct {
	beego.Controller
}

// DeleteController 删除日志
type DeleteController struct {
	beego.Controller
}

// Get 获取日志列表
func (c *LogController) Get() {
	o := orm.NewOrm()

	// 获取主机列表
	qs := o.QueryTable("Server")

	servers := []models.Server{}

	beego.Info(qs)

	_, err := qs.All(&servers)
	if err != nil {
		beego.Info("查询失败,err:", err)
	}

	// 分页设置
	pageSize := 8

	query := o.QueryTable("collect")
	count, err := query.Count()

	// 获取页面数，向上取整
	page, err := strconv.Atoi(c.GetString("page"))
	if err != nil {
		page = 1
	}
	start := pageSize * (page - 1)

	pageCount := int(math.Ceil(float64(count) / float64(pageSize)))

	// 获取日志收集项
	var collects []models.Collect
	table := o.QueryTable("collect")
	table.Limit(pageSize, start).RelatedSel("Server").All(&collects)

	beego.Info("总数：", count)

	c.Data["pageCount"] = pageCount
	c.Data["count"] = count
	c.Data["collect"] = collects
	c.Data["server"] = servers
	c.Data["prepage"] = 1
	c.Data["page"] = page

	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
}

// Post 日志处理
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

// Get 删除日志收集项
func (c *DeleteController) Get() {
	// 获取用户输入的ID
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("数据非法")
		return
	}
	beego.Info("获取到的ID：", id)

	// 执行删除操作
	o := orm.NewOrm()
	collect := models.Collect{Id: id}
	_, err = o.Delete(&collect)
	if err != nil {
		beego.Info("删除失败，", err)
		return
	}

	// 跳转到日志列表页
	c.Redirect("/admin/", http.StatusFound)
}
