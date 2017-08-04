package setting

import (
	"github.com/astaxie/beego"
)

func InitConfig() {
	NsqConn = beego.AppConfig.String("nsq_url")
	RunMode = beego.AppConfig.String("appver")
}
