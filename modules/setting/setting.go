package setting

import (
	"github.com/nsqio/go-nsq"
)

var (
	NsqConn string      // nsq 链接
	NsqConf *nsq.Config // nsq 配置
	RunMode string      // 运行模式
)
