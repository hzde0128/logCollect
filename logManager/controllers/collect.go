package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
	"github.com/hzde0128/logCollect/logManager/utils"
	"net/http"
)

// CollectController 日志收集
type CollectController struct {
	beego.Controller
}

// LogEntry 需要收集的日志的配置信息
type LogEntry struct {
	Path  string `json:"path"`  // 日志存放的路径
	Topic string `json:"topic"` // 日志要发往Kafka中的哪个Topic
}

// Get 日志收集列表展示
func (c *CollectController) Get() {
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

// Post 添加日志收集处理
func (c *CollectController) Post() {
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
	if err != nil {
		beego.Info("非法的服务器地址", err)
		return
	}
	collect.Server = &servers
	collect.Path = filepath
	collect.Topic = topic

	_, err = o.Insert(&collect)
	if err != nil {
		beego.Info("添加失败,", err)
		return
	}

	key := "/logagent/" + server + "/collect"
	var conf = LogEntry{
		Path:  filepath,
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
