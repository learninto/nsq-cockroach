package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "nsq-cockroach/initial"
	_ "nsq-cockroach/routers"
	"runtime"
)

func main() {

	// 文档
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	// 日志
	logs.SetLogger("console")

	// 系统
	runtime.GOMAXPROCS(runtime.NumCPU())

	beego.Run()
}
