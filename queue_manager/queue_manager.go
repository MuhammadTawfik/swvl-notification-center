package queue_manager

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to establish connection to message broker")
	ch, err := conn.Channel()
	failOnError(err, "Failed to get channel for connection")

	return conn, ch
}

func GetQueue(name string, ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name,  //name string,
		true,  //durable bool,
		false, //autoDelete bool,
		false, //exclusive bool,
		false, //noWait bool,
		nil)   //args amqp.Table)

	failOnError(err, "Failed to declare queue")

	return &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func GetPQueue(name string, max_priority int, ch *amqp.Channel) *amqp.Queue {
	args := amqp.Table{
		"x-max-priority": max_priority,
	}
	q, err := ch.QueueDeclare(
		name,  //name string,
		true,  //durable bool,
		false, //autoDelete bool,
		false, //exclusive bool,
		false, //noWait bool,
		args)  //args amqp.Table)

	failOnError(err, "Failed to declare queue")

	return &q
}
