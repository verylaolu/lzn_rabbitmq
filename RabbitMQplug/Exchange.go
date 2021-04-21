package RabbitMQplug

import "github.com/streadway/amqp"

func changeDeclare(RabbitMQ *RabbitMQ)error{

	return RabbitMQ.Channel.ExchangeDeclare(
		RabbitMQ.MQ_exchange, // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
}
func changeDeclareDelay(RabbitMQ *RabbitMQ)error{
	arg :=amqp.Table{
		"x-delayed-type":"direct",
	}
	return RabbitMQ.Channel.ExchangeDeclare(
		RabbitMQ.MQ_exchange, // name
		"x-delayed-message",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		arg,           // arguments
	)
}
func changeBind(RabbitMQ *RabbitMQ,suffix string)error{
	return RabbitMQ.Channel.QueueBind(
		RabbitMQ.MQ_queue+suffix, // queue name
		RabbitMQ.MQ_queue+suffix,     // routing key
		RabbitMQ.MQ_exchange, // exchange
		false,
		nil)
}
