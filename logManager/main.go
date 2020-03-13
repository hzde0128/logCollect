package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
	_ "github.com/hzde0128/logCollect/logManager/routers"
	_ "github.com/hzde0128/logCollect/logManager/models"
	"github.com/astaxie/beego"
)

func main() {
	// 上一页/下一页
	beego.AddFuncMap("prepage", prepage)
	beego.AddFuncMap("nextpage", nextpage)

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

func prepage(idx int) (page int){
	if idx > 1 {
		page = idx - 1
	} else{
		page = idx
	}
	return
}

func nextpage(idx, count int) (page int) {
	if idx < count {
		page = idx + 1
	} else {
		page = count
	}
	return
}