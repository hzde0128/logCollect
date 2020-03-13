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
	user.Password = "admin"
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败")
	}

	beego.Run()
}

