package setting

import (
	"github.com/astaxie/beego"
)

func InitConfig() {
	NsqConn = beego.AppConfig.String("nsq_conn")
	RunMode = beego.AppConfig.String("appver")
}
