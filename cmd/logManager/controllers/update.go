package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"logCollect/utils"

	"logCollect/cmd/logManager/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	collect.ID = id
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
	collect.ID = id

	_, err = o.Update(&collect)
	if err != nil {
		beego.Info("修改失败,", err)
		return
	}

	// 更新etcd
	key := "/logagent/" + server + "/collect"
	var conf = LogEntry{
		Path:  path,
		Topic: topic,
	}

	// 转化为json
	val, _ := json.Marshal(conf)
	beego.Info(string(val))
	_, err = utils.PutConf(key, "["+string(val)+"]")
	if err != nil {
		beego.Info("推送到etcd失败", err)
	}
	beego.Info("成功推送到etcd")

	c.Redirect("/admin/", http.StatusFound)
}
