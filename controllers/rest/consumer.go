package rest

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/learninton/beegolibs/restlib"
	"github.com/nsqio/go-nsq"
	"net/http"
	"nsq-cockroach/models"
	"nsq-cockroach/modules/util"
)

func init() {
	util.InitNsqConfig()
}

// 元素：消费
type ConsumerController struct {
	restlib.RestlibController
}

// @Title 消费
// @Description 消费
// @Param	topic	query	string 	true	主题
// @Param	channel	query	string 	true	通道
// @Success	Success	{"":""}
// @Failure	Failure	{"errcode": 错误码,"errmsg": "错误信息"}
// @router /consumer [get]
func (this *ConsumerController) ConsumerGet() {

	topic, channel := this.checkConsumerGetParameter()

	go util.NsqConsumer(topic, channel, &ConsumerT{})

	this.Success("Success")
}

/* 消费者 入参校验 */
func (this *ConsumerController) checkConsumerGetParameter() (string, string) {
	topic := this.GetString("topic")
	if topic == "" {
		this.Error("The topic parameters are not valid", http.StatusNotFound)
	}
	channel := this.GetString("channel")
	if channel == "" {
		this.Error("The channel parameters are not valid", http.StatusNotFound)
	}

	logs.Info("Consumer topic：%s，Consumer channel：%s", topic, channel)

	return topic, channel
}

type ConsumerT struct{}

//处理消息
func (*ConsumerT) HandleMessage(message *nsq.Message) error {

	var ppi models.ProducerPostIo
	err := json.Unmarshal(message.Body, &ppi)
	if err != nil {
		return err
	}

	logs.Info("Got a message Body: ", ppi.Body)
	logs.Info("Got a message StorageType: ", ppi.StorageType)
	logs.Info("Got a message StorageName: ", ppi.StorageName)

	return nil
}
