package RabbitMQplug

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

func createDelay(RabbitMQ *RabbitMQ)error{
	var err error
	//创建路由
	err = changeDeclareDelay(RabbitMQ)
	if(err != nil){
		err = errors.New(fmt.Sprintf("RabbitMQ，路由创建失败: %s \n",err))
	}
	//queue
	_,err = queueDeclareNormal(RabbitMQ)
	if(err != nil){
		err =  errors.New(fmt.Sprintf("RabbitMQ，队列创建失败: %s \n",err))
	}
	//绑定路由
	err = changeBind(RabbitMQ,"")
	if(err != nil){
		err =  errors.New(fmt.Sprintf("RabbitMQ，绑定队列失败: %s \n",err))
	}
	return err

}
func pushDelay(RabbitMQ *RabbitMQ,ttl int,body string,)error{
	return queuePushDelay(RabbitMQ,ttl,body)
}
func pullDelay(RabbitMQ *RabbitMQ) (<-chan amqp.Delivery,error){
	return queuePull(RabbitMQ,"")
}
