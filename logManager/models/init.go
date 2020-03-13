package models

import (
	"github.com/astaxie/beego/orm"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	orm.DefaultTimeLoc = time.Local
	//orm.RegisterDataBase("default", "sqlite3", "logmanager.db")
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/logcollect?charset=utf8")
	orm.RegisterModel(new(User), new(Server), new(Collect))
	orm.RunSyncdb("default", false, false)
}
