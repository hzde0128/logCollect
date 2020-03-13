package routers

import (
	"github.com/hzde0128/logCollect/logManager/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginController{})
	beego.Router("/admin/", &controllers.LogController{})
	beego.Router("/admin/collect/", &controllers.CollectController{})
	beego.Router("/admin/collect/?:id", &controllers.UpdateController{})
	beego.Router("/admin/server/", &controllers.ServerController{})
}
