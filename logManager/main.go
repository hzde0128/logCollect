package main

import (
	_ "logCollect/logManager/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

