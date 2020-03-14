package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
)

type DeleteController struct {
	beego.Controller
}

// 删除日志收集项
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
