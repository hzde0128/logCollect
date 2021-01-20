package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hzde0128/logCollect/logManager/utils"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	orm.DefaultTimeLoc = time.Local
	dsn := utils.GMysqlUser + ":" + utils.GMysqlPass + "@tcp(" + utils.GMysqlHost + ":" + utils.GMysqlPort + ")/" + utils.GMysqlDb
	beego.Info("dsn:", dsn)
	orm.RegisterDataBase("default", "mysql", dsn+"?charset=utf8")
	orm.RegisterModel(new(User), new(Server), new(Collect))
	orm.RunSyncdb("default", false, false)
}
