package Interface

import (
	"git.gbicom.com/develop/go-rabbitmq/RabbitMQplug"
	"log"
)

func test(){

	MQ_Full := RabbitMQplug.MQ_FULL{}
	//  创建MQ  ////////////////////////////////////////////////////////////////
	Config := MQ_Full.DefaultConfig("10.10.107.101","","30072","admin","admin")
	MQ_OBJ , _ := MQ_Full.Connent(Config)

	//  初始化交换机数据  ////////////////////////////////////////////////////////////////
	ExchangeParam := MQ_Full.DefaultExchangeDeclareParam("交换机")  //初始化交换机数据
	//ExchangeParam = MQ_Full.DelayExchangeDeclareParam(ExchangeParam)//如果是延迟队列，执行方法，修改初始化数据
	MQ_Full.ExchangeDeclare(MQ_OBJ,ExchangeParam) //初始化交换机

	//  初始化队列  ////////////////////////////////////////////////////////////////
	QueueParam := MQ_Full.DefaultQueueDeclareParam("队列")
	//QueueParam = MQ_Full.DieQueueDeclareParam(QueueParam,"交换机","路由",10) //死信队列，修改初始化参数  过期时间单位：秒
	_, _ = MQ_Full.QueueDeclare(MQ_OBJ,QueueParam) //初始化队列

	//  队列绑定交换机  ////////////////////////////////////////////////////////////////
	ExchangeBindParam := MQ_Full.DefaultExchangeBindParam(QueueParam.Name,ExchangeParam.Name,"路由")
	MQ_Full.ExchangeBind(MQ_OBJ,ExchangeBindParam)//队列绑定交换机

	//  推送数据  ////////////////////////////////////////////////////////////////
	PushParam := MQ_Full.DefaultQueuePushParam(ExchangeParam.Name,QueueParam.Name,"推送的数据")
	//PushParam = MQ_Full.DelayQueuePushParam(PushParam,10) //如果是延迟队列，执行方法  过期时间单位：秒
	MQ_Full.QueuePush(MQ_OBJ,PushParam)

	//  获取数据  ////////////////////////////////////////////////////////////////
	PullParam :=MQ_Full.DefaultQueuePullParam(QueueParam.Name)
	//PullParam.AutoAck = false //需要ACK
	//PullParam.AutoAck = true  //不需要ACK
	msgs,_:=MQ_Full.QueuePull(MQ_OBJ,PullParam)
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("获取消息内容: %s", msg.Body)
			//需要ACK
			if(PullParam.AutoAck==false){
				if string(msg.Body) != ""{
					MQ_Full.ACK(msg,false)//确认消息销毁
					log.Printf("ACK")
				}else{
					MQ_Full.NACK(msg,false,true) //操作失败回滚队列
					log.Printf("NACK")
				}
			}

		}
	}()
	<-forever
}
