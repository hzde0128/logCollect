package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
	"net/http"
	"time"
)

type LoginController struct {
	beego.Controller
}

// 登录页面
func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

// 登录处理页面
func (c *LoginController) Post() {
	// 处理登录请求
	// 1.获取用户输入的数据
	userName := c.GetString("userName")
	password := c.GetString("password")
	beego.Info(userName, password)

	// 2.数据处理
	if userName == "" || password == "" {
		c.TplName = "login.tpl"
		return
	}

	// 3.数据查询
	// 3.1获取orm对象
	o := orm.NewOrm()

	// 3.2获取查询对象
	user := models.User{}

	// 3.3查询数据
	user.Username = userName

	err := o.Read(&user, "userName")
	if err != nil {
		beego.Info("用户名或密码错误")
		c.TplName = "login.tpl"
		return
	}
	if user.Password != password {
		beego.Info("用户名或密码错误")
		c.TplName = "login.tpl"
		return
	}

	// 设置Cookie
	c.Ctx.SetCookie("username", userName, time.Second * 3600)

	// 4.返回视图
	c.Redirect("/admin/", http.StatusFound)
}