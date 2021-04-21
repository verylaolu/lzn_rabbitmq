package RabbitMQplug

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)
type RabbitMsg struct {
	Body string
	Ttl int
}
type RabbitMQ struct {
	Connect *amqp.Connection
	Channel *amqp.Channel
	MQ_type string
	MQ_exchange string
	MQ_queue string
	MQ_ttl int
	RabbitConfig  RabbitConfig
}
type RabbitConfig struct {
	Host string
	Virtual_Host string
	Port string
	Username string
	Password string
}

func Initial(RabbitMQ RabbitMQ)(*RabbitMQ,error){
	var err error
	if(RabbitMQ.MQ_exchange == ""|| RabbitMQ.MQ_queue==""){
		return &RabbitMQ,errors.New(fmt.Sprintf("路由或队列信息错误: [exchange:%s;queue:%s] \n",RabbitMQ.MQ_exchange,RabbitMQ.MQ_queue))
	}
	//Connect
	RabbitMQ.Connect,err = makeConnect(RabbitMQ.RabbitConfig)
	if(err != nil){
		return &RabbitMQ,errors.New(fmt.Sprintf("RabbitMQ链接打开失败: %s \n",err))
	}
	//Channel
	RabbitMQ.Channel,err = makeChannel(RabbitMQ.Connect)
	if(err != nil){
		return &RabbitMQ,errors.New(fmt.Sprintf("RabbitMQ管道打开失败: %s \n",err))
	}
	return &RabbitMQ,nil
}
func Declare(RabbitMQ *RabbitMQ)error{
	var err error
	switch RabbitMQ.MQ_type {
		case "dead":
			if RabbitMQ.MQ_ttl<=0{
				err = errors.New(fmt.Sprintf("死信延迟队列过期时间设置错误: [ttl:%s] \n",RabbitMQ.MQ_ttl))
			}
			err = createDead(RabbitMQ)
		case "delay":
			err = createDelay(RabbitMQ)
		default:
			err = createNormal(RabbitMQ)
	}
	return err
}
func Push(RabbitMQ *RabbitMQ,msg RabbitMsg) error {
	var err error
	switch RabbitMQ.MQ_type {
	case "dead":
		err = pushDead(RabbitMQ,msg.Body)
	case "delay":
		err =  pushDelay(RabbitMQ,msg.Ttl,msg.Body)
	default:
		err =  pushNormal(RabbitMQ,msg.Body)
	}
	return err
}
func Pull(RabbitMQ *RabbitMQ) (<-chan amqp.Delivery,error) {
	if RabbitMQ.MQ_type=="dead"{
		return pullDead(RabbitMQ)
	}else{
		return pullNormal(RabbitMQ)
	}
}

func makeConnect(config RabbitConfig)(*amqp.Connection,error){
	RabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", config.Username, config.Password,config.Host,  config.Port,config.Virtual_Host)
	return amqp.Dial(RabbitUrl)
}
func makeChannel(mq *amqp.Connection)(*amqp.Channel,error){
	return mq.Channel()
}