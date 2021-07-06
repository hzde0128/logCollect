package models

import (
	"logCollect/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	utils.Init()
	orm.DefaultTimeLoc = time.Local
	dsn := utils.GMysqlUser + ":" + utils.GMysqlPass + "@tcp(" + utils.GMysqlHost + ":" + utils.GMysqlPort + ")/" + utils.GMysqlDb
	beego.Info("dsn:", dsn)
	orm.RegisterDataBase("default", "mysql", dsn+"?charset=utf8")
	orm.RegisterModel(new(User), new(Server), new(Collect))
	orm.RunSyncdb("default", false, false)
}
