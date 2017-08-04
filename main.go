package main

import (
	"github.com/astaxie/beego"
	_ "nsq-cockroach/routers"
)

func main() {
	beego.Run()
}
