package RabbitMQplug

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

func createDead(RabbitMQ *RabbitMQ)error{
	var err error
	//创建路由
	err = changeDeclare(RabbitMQ)
	if(err != nil){
		return errors.New(fmt.Sprintf("RabbitMQ，路由创建失败: %s \n",err))
	}
	//创建queue+"_live"
	_,err = queueDeclareLive(RabbitMQ)
	if(err != nil){
		return errors.New(fmt.Sprintf("RabbitMQ，延迟LIVE队列创建失败: %s \n",err))
	}
	//创建queue+"_die" ttl latter_exchange:exchange routing-key:queue+"_live"
	_,err = queueDeclareDie(RabbitMQ)
	if(err != nil){
		return errors.New(fmt.Sprintf("RabbitMQ，延迟死信队列创建失败: %s \n",err))
	}
	//绑定路由
	err = changeBind(RabbitMQ,"_live")
	if(err != nil){
		return errors.New(fmt.Sprintf("RabbitMQ，绑定LIVE队列失败: %s \n",err))
	}
	err = changeBind(RabbitMQ,"_die")
	if(err != nil){
		return errors.New(fmt.Sprintf("RabbitMQ，绑定死信队列失败: %s \n",err))
	}
	return nil

}
func pushDead(RabbitMQ *RabbitMQ,body string)error{
	return queuePush(RabbitMQ,"_die",body)
}
func pullDead(RabbitMQ *RabbitMQ) (<-chan amqp.Delivery,error){
	return queuePull(RabbitMQ,"_live")
}
