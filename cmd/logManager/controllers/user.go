package controllers

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"logCollect/utils"

	"logCollect/cmd/logManager/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// LoginController 登录
type LoginController struct {
	beego.Controller
}

// LogoutController 登出
type LogoutController struct {
	beego.Controller
}

// Get 登录页面
func (c *LoginController) Get() {
	// 获取Cookie
	username := c.Ctx.GetCookie("username")
	if username != "" {
		c.Data["username"] = username
		c.Data["check"] = "checked"
	}
	c.TplName = "login.tpl"
}

// Post 登录处理页面
func (c *LoginController) Post() {
	// 处理登录请求
	// 1.获取用户输入的数据
	userName := c.GetString("userName")
	password := c.GetString("password")
	// 检查是否记住用户名
	remember := c.GetString("remember")
	beego.Info(userName, password, remember)

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
	//md5验证
	data := []byte(password)
	has := md5.Sum(data)
	if user.Password != fmt.Sprintf("%x", has) {
		beego.Info("用户名或密码错误")
		c.TplName = "login.tpl"
		return
	}

	// 设置Cookie
	if remember == "on" {
		c.Ctx.SetCookie("username", userName, 7*24*time.Hour)
	} else {
		c.Ctx.SetCookie("username", userName, -1)
	}
	// 设置session
	c.SetSession("username", userName)
	// 4.返回视图
	c.Redirect("/admin/", http.StatusFound)

	// 创建sessionId
	sessionID := utils.Md5String(userName)
	// 将用户信息加入redis中
	bm, err := utils.RedisConn()
	if err != nil {
		beego.Info("Redis连接失败", err)
	}
	// session保留10分钟
	err = bm.Put(sessionID, userName, time.Second*600)
	if err != nil {
		beego.Info("存入redis失败", err)
	}
}

// Get 退出登录
func (c *LogoutController) Get() {
	// 1.删除session
	c.DelSession("username")

	// 2.跳转到登录页面
	c.Redirect("/", http.StatusFound)
}
