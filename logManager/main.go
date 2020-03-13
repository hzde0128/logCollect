package main

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/models"
	_ "github.com/hzde0128/logCollect/logManager/models"
	_ "github.com/hzde0128/logCollect/logManager/routers"
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
		// md5存放
		data := []byte("admin")
		has := md5.Sum(data)
		user.Password = fmt.Sprintf("%x", has)
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