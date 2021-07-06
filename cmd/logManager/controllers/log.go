package controllers

import (
	"math"
	"net/http"
	"strconv"

	"logCollect/utils"

	"logCollect/cmd/logManager/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	if err != nil {
		beego.Info("查询失败,", err)
		return
	}

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

	o := orm.NewOrm()
	// 获取对应的主机
	server := models.Server{}
	o.Raw("SELECT address FROM server WHERE id = (SELECT server_id FROM collect WHERE id = ?)", id).QueryRow(&server)

	beego.Info(server)
	// 执行删除操作
	collect := models.Collect{ID: id}
	_, err = o.Delete(&collect)
	if err != nil {
		beego.Info("删除失败，", err)
		return
	}

	// 拼接key
	key := "/logagent/" + server.Address + "/collect"
	//// 同时删除etcd配置
	_, err = utils.DelConf(key)
	if err != nil {
		beego.Info("删除失败", err)
	}
	beego.Info("删除成功")
	// 跳转到日志列表页
	c.Redirect("/admin/", http.StatusFound)
}
