package Interface

import (
	"github.com/verylaolu/lzn_rabbitmq/RabbitMQplug"
	"log"
)

func (Pull Pull_OBJ)Dead(){
	var RabbitMQ = RabbitMQplug.RabbitMQ{}
	RabbitMQ.MQ_ttl=15
	RabbitMQ.MQ_exchange="exchange_dead"
	RabbitMQ.MQ_queue="queue_dead"
	RabbitMQ.MQ_type="dead"  //dead //normal
	RabbitMQ.RabbitConfig = RabbitMQ_Connect
	Rabbit_Pointer,err:= RabbitMQplug.Initial(RabbitMQ)
	//RabbitMQplug.Declare(Rabbit_Pointer)  //第一次创建需要使用， 之后不需要
	if err != nil{
		panic(err)
	}
	Client(Rabbit_Pointer)

}
func (Pull Pull_OBJ)Delay(){
	var RabbitMQ = RabbitMQplug.RabbitMQ{}
	RabbitMQ.MQ_exchange="exchange_delay"
	RabbitMQ.MQ_queue="queue_delay"
	RabbitMQ.MQ_type="delay"  //dead //normal
	RabbitMQ.RabbitConfig = RabbitMQ_Connect
	Rabbit_Pointer,err:= RabbitMQplug.Initial(RabbitMQ)
	//RabbitMQplug.Declare(Rabbit_Pointer)  //第一次创建需要使用， 之后不需要
	if err != nil{
		panic(err)
	}
	Client(Rabbit_Pointer)
}
func (Pull Pull_OBJ)Normal(){
	var RabbitMQ = RabbitMQplug.RabbitMQ{}
	RabbitMQ.MQ_ttl=15
	RabbitMQ.MQ_exchange="exchange_normal"
	RabbitMQ.MQ_queue="queue_normal"
	RabbitMQ.MQ_type="normal"  //dead //normal
	RabbitMQ.RabbitConfig = RabbitMQ_Connect
	Rabbit_Pointer,err:= RabbitMQplug.Initial(RabbitMQ)
	//RabbitMQplug.Declare(Rabbit_Pointer)  //第一次创建需要使用， 之后不需要
	if err != nil{
		panic(err)
	}
	Client(Rabbit_Pointer)
}
func Client(Rabbit_Pointer *RabbitMQplug.RabbitMQ){
	msgs,err:=RabbitMQplug.Pull(Rabbit_Pointer)
	if err != nil{
		panic(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			if string(d.Body) != ""{
				d.Ack(false)   //确认使用
				log.Printf("ACK")
			}else{
				d.Nack(false,true)  //回滚队列
				log.Printf("NACK")
			}

		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
