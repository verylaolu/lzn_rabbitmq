package RabbitMQplug

import (
	"github.com/streadway/amqp"
)
func queueDeclareDie(RabbitMQ *RabbitMQ) (amqp.Queue,error){
	arg :=amqp.Table{
		"x-dead-letter-exchange":RabbitMQ.MQ_exchange,
		"x-dead-letter-routing-key":RabbitMQ.MQ_queue+"_live",
		"x-message-ttl":RabbitMQ.MQ_ttl*1000,
	}

	return RabbitMQ.Channel.QueueDeclare(
		RabbitMQ.MQ_queue+"_die", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		arg,     // arguments
	)


}
func queueDeclareLive(RabbitMQ *RabbitMQ) (amqp.Queue,error){

	return RabbitMQ.Channel.QueueDeclare(
		RabbitMQ.MQ_queue+"_live", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

}
func queueDeclareNormal(RabbitMQ *RabbitMQ) (amqp.Queue,error){

	return RabbitMQ.Channel.QueueDeclare(
		RabbitMQ.MQ_queue, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

}
func queuePush(RabbitMQ *RabbitMQ,suffix string,body string)error{
	return RabbitMQ.Channel.Publish(
		RabbitMQ.MQ_exchange,     // exchange
		RabbitMQ.MQ_queue+suffix, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})

}
func queuePushDelay(RabbitMQ *RabbitMQ,ttl int,body string)error{
	return RabbitMQ.Channel.Publish(
		RabbitMQ.MQ_exchange,     // exchange
		RabbitMQ.MQ_queue, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			Headers:amqp.Table{
				"x-delay":ttl*1000,
			},
			ContentType: "text/plain",
			Body:        []byte(body),
		})

}

func queuePull(RabbitMQ *RabbitMQ,suffix string) (<-chan amqp.Delivery,error) {
	return  RabbitMQ.Channel.Consume(
		RabbitMQ.MQ_queue+suffix, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
}