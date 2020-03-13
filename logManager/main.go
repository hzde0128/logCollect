package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
	_ "github.com/hzde0128/logCollect/logManager/routers"
	_ "github.com/hzde0128/logCollect/logManager/models"
	"github.com/astaxie/beego"
)

func main() {
	orm.Debug = true

	// 初始化一个管理用户
	o := orm.NewOrm()

	user := models.User{}
	user.Username = "admin"

	err := o.Read(&user, "username")
	if err != nil{
		beego.Info("用户admin不存在，添加用户")
		user.Password = "admin"
		o.Insert(&user)
	}

	beego.Run()
}

