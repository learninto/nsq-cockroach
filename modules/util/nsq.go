package util

import (
	"github.com/astaxie/beego/logs"
	"github.com/nsqio/go-nsq"
	"nsq-cockroach/modules/setting"
	"time"
)

// 初始化生产者
func InitNsqConfig() {
	config := nsq.NewConfig()
	config.DefaultRequeueDelay = 0
	config.MaxBackoffDuration = 20 * time.Millisecond
	config.LookupdPollInterval = 1000 * time.Millisecond
	config.RDYRedistributeInterval = 1000 * time.Millisecond
	config.MaxInFlight = 2500

	setting.NsqConf = config
}

//发布消息
func NsqPublish(topic string, message []byte) error {

	logs.Info("nsq_producer_address", setting.NsqConn)

	producer, err := nsq.NewProducer(setting.NsqConn, setting.NsqConf)
	if err != nil {
		//panic(err)
		return err
	}

	err1 := producer.Publish(topic, message)
	if err != nil {
		return err1
	}

	producer.Stop()
	return nil
}
