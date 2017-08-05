// @APIVersion 1.0.0
// @Title beego nsq
// @Description beego nsq
// @Contact learninton@gmail.com
// @TermsOfServiceUrl https://github.com/learninton/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
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
				&rest.ConsumerController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
