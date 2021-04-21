package Interface

import (
	"github.com/verylaolu/lzn_rabbitmq/RabbitMQplug"
)


func (Push Push_OBJ)Dead(){
	var RabbitMQ = RabbitMQplug.RabbitMQ{}
	RabbitMQ.MQ_ttl=15
	RabbitMQ.MQ_exchange="exchange_dead"
	RabbitMQ.MQ_queue="queue_dead"
	RabbitMQ.MQ_type="dead"  //dead //normal
	RabbitMQ.RabbitConfig = RabbitMQ_Connect
	Rabbit_Pointer,err:= RabbitMQplug.Initial(RabbitMQ)
	if err != nil{
		panic(err)
	}
	err = RabbitMQplug.Declare(Rabbit_Pointer)  //第一次创建需要使用， 之后不需要
	if err != nil{
		panic(err)
	}
	var m1,m2,m3,m4,m5,m6  RabbitMQplug.RabbitMsg
	m1.Body="11111111"
	m2.Body="22222222"
	m3.Body="33333333"
	m4.Body="44444444"
	m5.Body="55555555"
	m6.Body="66666666"
	err = RabbitMQplug.Push(Rabbit_Pointer,m1)
	err = RabbitMQplug.Push(Rabbit_Pointer,m2)
	err = RabbitMQplug.Push(Rabbit_Pointer,m3)
	err = RabbitMQplug.Push(Rabbit_Pointer,m4)
	err = RabbitMQplug.Push(Rabbit_Pointer,m5)
	err = RabbitMQplug.Push(Rabbit_Pointer,m6)
	if err != nil{
		panic(err)
	}
}
func (Push Push_OBJ)Delay(){
	var RabbitMQ = RabbitMQplug.RabbitMQ{}
	RabbitMQ.MQ_exchange="exchange_delay"
	RabbitMQ.MQ_queue="queue_delay"
	RabbitMQ.MQ_type="delay"  //dead //normal
	RabbitMQ.RabbitConfig = RabbitMQ_Connect
	Rabbit_Pointer,err:= RabbitMQplug.Initial(RabbitMQ)
	if err != nil{
		panic(err)
	}
	err = RabbitMQplug.Declare(Rabbit_Pointer)  //第一次创建需要使用， 之后不需要
	if err != nil{
		panic(err)
	}
	var m1,m2,m3,m4,m5,m6  RabbitMQplug.RabbitMsg
	m1.Body="11111111"
	m1.Ttl=10
	m2.Body="22222222"
	m2.Ttl=5
	m3.Body="33333333"
	m3.Ttl=20
	m4.Body="44444444"
	m4.Ttl=15
	m5.Body="55555555"
	m5.Ttl=30
	m6.Body="66666666"
	m6.Ttl=25
	//2,1,4,3,6,5
	err = RabbitMQplug.Push(Rabbit_Pointer,m1)
	err = RabbitMQplug.Push(Rabbit_Pointer,m2)
	err = RabbitMQplug.Push(Rabbit_Pointer,m3)
	err = RabbitMQplug.Push(Rabbit_Pointer,m4)
	err = RabbitMQplug.Push(Rabbit_Pointer,m5)
	err = RabbitMQplug.Push(Rabbit_Pointer,m6)
	if err != nil{
		panic(err)
	}
}
func (Push Push_OBJ)Normal(){
	var RabbitMQ = RabbitMQplug.RabbitMQ{}
	RabbitMQ.MQ_ttl=15
	RabbitMQ.MQ_exchange="exchange_normal"
	RabbitMQ.MQ_queue="queue_normal"
	RabbitMQ.MQ_type="normal"  //dead //normal
	RabbitMQ.RabbitConfig = RabbitMQ_Connect
	Rabbit_Pointer,err:= RabbitMQplug.Initial(RabbitMQ)
	if err != nil{
		panic(err)
	}
	err = RabbitMQplug.Declare(Rabbit_Pointer)  //第一次创建需要使用， 之后不需要
	if err != nil{
		panic(err)
	}
	var m1,m2,m3,m4,m5,m6  RabbitMQplug.RabbitMsg
	m1.Body="11111111"
	m2.Body="22222222"
	m3.Body="33333333"
	m4.Body="44444444"
	m5.Body="55555555"
	m6.Body="66666666"
	err = RabbitMQplug.Push(Rabbit_Pointer,m1)
	err = RabbitMQplug.Push(Rabbit_Pointer,m2)
	err = RabbitMQplug.Push(Rabbit_Pointer,m3)
	err = RabbitMQplug.Push(Rabbit_Pointer,m4)
	err = RabbitMQplug.Push(Rabbit_Pointer,m5)
	err = RabbitMQplug.Push(Rabbit_Pointer,m6)
	if err != nil{
		panic(err)
	}
}
