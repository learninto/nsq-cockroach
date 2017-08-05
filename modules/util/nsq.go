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
	config.LookupdPollInterval = 1000 * time.Millisecond //设置重连时间
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

// 消费者
func NsqConsumer(_topic string, _channel string, handler nsq.Handler) {
	logs.Info("Nsq consumer conn : ", setting.NsqConn)

	c, err := nsq.NewConsumer(_topic, _channel, setting.NsqConf) // 新建一个消费者
	if err != nil {
		panic(err)
	}

	c.SetLogger(nil, 0)   //系统日志
	c.AddHandler(handler) // 添加消费者接口

	////建立NSQLookupd连接
	//if err := c.ConnectToNSQLookupd(setting.NsqConsumerConn); err != nil {
	//	panic(err)
	//}

	//建立多个nsqd连接
	// if err := c.ConnectToNSQDs([]string{"127.0.0.1:4150", "127.0.0.1:4152"}); err != nil {
	//  panic(err)

	// }

	// 建立一个nsqd连接
	if err := c.ConnectToNSQD(setting.NsqConn); err != nil {
		panic(err)
	}

	<-c.StopChan
}
