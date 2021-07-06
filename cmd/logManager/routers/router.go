package routers

import (
	"net/http"

	"logCollect/cmd/logManager/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	// 路由过滤器，在进入后台之前进行session判断
	beego.InsertFilter("/admin/*", beego.BeforeRouter, filterFunc)
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/admin/", &controllers.LogController{})
	beego.Router("/admin/collect/", &controllers.CollectController{})
	beego.Router("/admin/collect/delete", &controllers.DeleteController{})
	beego.Router("/admin/collect/?:id", &controllers.UpdateController{})
	beego.Router("/admin/server/", &controllers.ServerController{})
	beego.Router("/admin/server/delete", &controllers.DeleteSvrController{})
}

var filterFunc = func(ctx *context.Context) {
	username := ctx.Input.Session("username")
	if username == nil {
		ctx.Redirect(http.StatusFound, "/")
	}
}
