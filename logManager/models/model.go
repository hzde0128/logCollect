package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
"time"
)

type User struct {
	Id int
	Username string `orm:"unique"`
	Password string
}

func init(){
	orm.DefaultTimeLoc = time.Local
	orm.RegisterDataBase("default", "sqlite3", "logmanager.db")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, false)
}