package RabbitMQplug

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

func createNormal(RabbitMQ *RabbitMQ)error{
	var err error
	//创建路由
	err = changeDeclare(RabbitMQ)
	if(err != nil){
		return errors.New(fmt.Sprintf("RabbitMQ，路由创建失败: %s \n",err))
	}
	//queue
	_,err = queueDeclareNormal(RabbitMQ)
	if(err != nil){
		return errors.New(fmt.Sprintf("RabbitMQ，队列创建失败: %s \n",err))
	}
	//绑定路由
	err = changeBind(RabbitMQ,"")
	if(err != nil){
		return errors.New(fmt.Sprintf("RabbitMQ，绑定队列失败: %s \n",err))
	}
	return nil

}
func pushNormal(RabbitMQ *RabbitMQ,body string)error{
	return queuePush(RabbitMQ,"",body)
}
func pullNormal(RabbitMQ *RabbitMQ) (<-chan amqp.Delivery,error){
	return queuePull(RabbitMQ,"")
}
