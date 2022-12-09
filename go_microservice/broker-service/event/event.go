package event

import (
	// adv messaging queue protocol
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //name of exchange
		"topic",      //type
		true,         // durable? keep it running
		false,        // auto-delete? get rid of it when done?
		false,        // internal?
		false,        // no-wait?
		nil,          // arguments
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable, get rid when done
		false, // auto-delete
		true,  // exclusive channel? dont share it
		false, // no-wait?
		nil,   // arguments?
	)
}
