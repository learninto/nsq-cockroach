package routers

import (
	"github.com/astaxie/beego"
	"nsq-cockroach/controllers"
	"nsq-cockroach/controllers/rest"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/rest",
			beego.NSInclude(
				&rest.ProducerController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
