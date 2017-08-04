package rest

import (
	"github.com/astaxie/beego/logs"
	"github.com/learninton/beegolibs/restlib"
	"net/http"
	"nsq-cockroach/models"
	"nsq-cockroach/modules/util"
)

func init() {
	// 初始化配置
	util.InitNsqConfig()
}

// 元素：生产者
type ProducerController struct {
	restlib.RestlibController
}

// @Title 生产者
// @Description 生产者
// @Param	topic	query	string 	true	主题
// @Param	body	body	rest.ProducerPostIo	true	{"body":"主体","storage_name":"存储名称","storage_type":"存储类型[file]"}
// @Success	Success	{"body":"主体","storage_name":"存储名称","storage_type":"存储类型[file]"}
// @Failure	Failure	{"error_code": 错误码,"error_msg": "错误信息"}
// @router /producer [post]
func (this *ProducerController) ProducerPost() {

	requestBody, params, topic := this.checkPostParameter()

	// push message
	err := util.NsqPublish(topic, requestBody)
	if err != nil {
		this.Error("push message error", http.StatusInternalServerError)
	}

	this.Success(params)
}

/* 生产者 入参校验 */
func (this *ProducerController) checkPostParameter() ([]byte, models.ProducerPostIo, string) {
	var params models.ProducerPostIo
	requestBody, err := this.GetPostJson(&params)

	if err != nil {
		logs.Error("The producer request parameters are not valid", err)
		this.Error("The producer request parameters are not valid", http.StatusNotFound)
	}
	logs.Info("The producer request parameters" + string(requestBody))

	if params.Body == "" {
		this.Error("The body parameters are not valid", http.StatusNotFound)
	}

	if params.StorageName == "" {
		this.Error("The storage_name parameters are not valid", http.StatusNotFound)
	}

	if params.StorageType == "" {
		this.Error("The storage_type parameters are not valid", http.StatusNotFound)
	}

	topic := this.GetString("topic")
	if topic == "" {
		this.Error("The topic parameters are not valid", http.StatusNotFound)
	}

	return requestBody, params, topic
}
