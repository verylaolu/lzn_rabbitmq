package RabbitMQplug

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)
type MQ_FULL struct {
}
type MQ_OBJ struct {
	Conn Conn
	Config Config
}
type Conn struct{
	Connect *amqp.Connection
	Channel *amqp.Channel
}
type Config struct {
	Host string
	Virtual_Host string
	Port string
	Username string
	Password string
}
type ExchangeDeclareParam struct {
	Name string
	Kind string
	Durable bool
	AutoDelete bool
	Internal bool
	NoWait bool
	Args amqp.Table
}
type ExchangeBindParam struct {
	QuereName string
	ExchangeName string
	RoutingKey string
	NoWait bool
	Args amqp.Table
}
type QueueDeclareParam struct {
	Name string
	Durable bool
	AutoDelete bool
	Exclusive bool
	NoWait bool
	Args amqp.Table
}
type QueuePushParam struct {
	ExchangeName string
	QueueName string
	Mandatory bool
	Immediate bool
	Publish struct{
		Headers amqp.Table
		ContentType string
		Body []byte
	}
}
type QueuePullParam struct {
	QueueName string
	Consumer string
	AutoAck bool
	Exclusive bool
	NoLocal bool
	NoWait bool
	Args amqp.Table
}

//初始化MQ配置信息
func (MQ_FULL) DefaultConfig(Host string, Virtual_Host string, Port string, Username string, Password string)Config{
	var Config Config
	Config.Host = Host
	Config.Virtual_Host = Virtual_Host
	Config.Port = Port
	Config.Username = Username
	Config.Password = Password
	return Config
}
//初始化交换机声明参数设置
func (MQ_FULL) DefaultExchangeDeclareParam(Name string)ExchangeDeclareParam{
 	var ExchangeDeclareParam ExchangeDeclareParam
	ExchangeDeclareParam.Name= Name
	ExchangeDeclareParam.Kind= "direct" //"x-delayed-message",
	ExchangeDeclareParam.Durable= true
	ExchangeDeclareParam.AutoDelete= false
	ExchangeDeclareParam.Internal=false
	ExchangeDeclareParam.NoWait= false
	ExchangeDeclareParam.Args= nil  //arg :=amqp.Table{"x-delayed-type":"direct",}
	return ExchangeDeclareParam
}
//初始化交换机与队列绑定参数设置
func (MQ_FULL) DefaultExchangeBindParam(QuereName string,ExchangeName string,RoutingKey string)ExchangeBindParam{
	var ExchangeBindParam ExchangeBindParam
	ExchangeBindParam.QuereName =QuereName
	ExchangeBindParam.ExchangeName =ExchangeName
	ExchangeBindParam.RoutingKey = RoutingKey
	ExchangeBindParam.NoWait = false
	ExchangeBindParam.Args = nil
	return ExchangeBindParam
}
//初始化获取队列数据参数设置
func(MQ_FULL)DefaultQueuePullParam(QueueName string)QueuePullParam{
	var QueuePullParam QueuePullParam
	QueuePullParam.QueueName = QueueName
	QueuePullParam.Consumer = ""
	QueuePullParam.AutoAck = false
	QueuePullParam.Exclusive = false
	QueuePullParam.NoLocal = false
	QueuePullParam.NoWait = false
	QueuePullParam.Args = nil
	return QueuePullParam
}
//初始化推送数据到队列参数设置
func(MQ_FULL)DefaultQueuePushParam(ExchangeName string,QueueName string,Body string)QueuePushParam{
	var QueuePushParam QueuePushParam
	QueuePushParam.ExchangeName = ExchangeName
	QueuePushParam.QueueName = QueueName
	QueuePushParam.Mandatory = false
	QueuePushParam.Immediate = false
	QueuePushParam.Publish.ContentType = "text/plain"
	QueuePushParam.Publish.Headers = nil
	QueuePushParam.Publish.Body = []byte(Body)
	return QueuePushParam
}
//初始化队列声明参数设置
func (MQ_FULL) DefaultQueueDeclareParam(Name string)QueueDeclareParam{
	var QueueDeclareParam QueueDeclareParam
	QueueDeclareParam.Name= Name
	QueueDeclareParam.Durable= false
	QueueDeclareParam.AutoDelete= false
	QueueDeclareParam.Exclusive= false
	QueueDeclareParam.NoWait= false
	QueueDeclareParam.Args = nil
	//arg :=amqp.Table{
	//	"x-dead-letter-exchange":exchangename,
	//	"x-dead-letter-routing-key":MQ_queue+"_live",
	//	"x-message-ttl":MQ_ttl*1000,
	//}
	return QueueDeclareParam
}

//链接RABBITMQ链接
func (MQ_FULL) Connent(config Config)(*MQ_OBJ,error){
	var err error
	var mq_obj MQ_OBJ
	mq_obj.Config = config
	//Connect
	RabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", config.Username, config.Password,config.Host,  config.Port,config.Virtual_Host)
	mq_obj.Conn.Connect,err = amqp.Dial(RabbitUrl)
	if(err != nil){
		return &mq_obj,errors.New(fmt.Sprintf("RabbitMQ链接打开失败: %s \n",err))
	}
	//Channel
	mq_obj.Conn.Channel,err = mq_obj.Conn.Connect.Channel()
	if(err != nil){
		return &mq_obj,errors.New(fmt.Sprintf("RabbitMQ管道打开失败: %s \n",err))
	}
	return &mq_obj,nil
}
//初始化交换机
func(MQ_FULL) ExchangeDeclare(MQ_OBJ *MQ_OBJ,param ExchangeDeclareParam)error{
	return MQ_OBJ.Conn.Channel.ExchangeDeclare(
		param.Name, // name
		param.Kind,      // kind
		param.Durable,          // durable
		param.AutoDelete,         // auto-deleted
		param.Internal,         // internal
		param.NoWait,         // no-wait
		param.Args,           // arguments
	)
}
//交换机绑定队列
func (MQ_FULL) ExchangeBind(MQ_OBJ *MQ_OBJ,param ExchangeBindParam)error{
	return MQ_OBJ.Conn.Channel.QueueBind(
		param.QuereName, // queue name
		param.RoutingKey,     // routing key
		param.ExchangeName, // exchange
		param.NoWait,
		param.Args)
}
//延时队列交换机参数设置
func (MQ_FULL) DelayExchangeDeclareParam(ExchangeDeclareParam ExchangeDeclareParam)ExchangeDeclareParam{
	ExchangeDeclareParam.Kind="x-delayed-message"
	ExchangeDeclareParam.Args=amqp.Table{"x-delayed-type":"direct",}
	return ExchangeDeclareParam
}


//初始化队列
func (MQ_FULL) QueueDeclare(MQ_OBJ *MQ_OBJ, param QueueDeclareParam) (amqp.Queue,error){

	return MQ_OBJ.Conn.Channel.QueueDeclare(
		param.Name, // name
		param.Durable,   // durable
		param.AutoDelete,   // delete when unused
		param.Exclusive,   // exclusive
		param.NoWait,   // no-wait
		param.Args,     // arguments
	)
}
//死信队列参数设置
func (MQ_FULL) DieQueueDeclareParam(QueueDeclareParam QueueDeclareParam,Exchange string,RoutingKey string,TTL int)QueueDeclareParam{
	QueueDeclareParam.Args =  amqp.Table{
		"x-dead-letter-exchange":Exchange,
		"x-dead-letter-routing-key":RoutingKey,
		"x-message-ttl":TTL*1000,
	}
	return QueueDeclareParam
}

//延迟队列推送数据参数设置
func(MQ_FULL)DelayQueuePushParam(QueuePushParam QueuePushParam,TTL int)QueuePushParam{
	QueuePushParam.Publish.Headers = amqp.Table{
		"x-delay":TTL*1000,
	}
	return QueuePushParam
}

//推送数据至队列
func (MQ_FULL) QueuePush(MQ_OBJ *MQ_OBJ, param QueuePushParam)error{
	return MQ_OBJ.Conn.Channel.Publish(
		param.ExchangeName,     // exchange
		param.QueueName, // routing key
		param.Mandatory,  // mandatory
		param.Immediate,  // immediate
		amqp.Publishing {
			Headers:param.Publish.Headers,
			ContentType: param.Publish.ContentType,
			Body:        param.Publish.Body,
		})

}
//获取队列数据
func (MQ_FULL) QueuePull(MQ_OBJ *MQ_OBJ,param QueuePullParam) (<-chan amqp.Delivery,error) {
	return  MQ_OBJ.Conn.Channel.Consume(
		param.QueueName, // queue
		param.Consumer,     // consumer
		param.AutoAck,   // auto-ack
		param.Exclusive,  // exclusive
		param.NoLocal,  // no-local
		param.NoWait,  // no-wait
		param.Args,    // args
	)
}
//确认消息
func(MQ_FULL)ACK(Delivery amqp.Delivery,Multiple bool)error{
	return Delivery.Ack(Multiple)
}
//回滚队列
func(MQ_FULL)NACK(Delivery amqp.Delivery,Multiple bool,Requeue bool)error{
	return Delivery.Nack(Multiple,Requeue)
}