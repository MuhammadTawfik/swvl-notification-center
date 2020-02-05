package front_line_consumers

import (
	"encoding/json"
	"fmt"
	"github.com/MuhammadTawfik/notifications/dispatcher"
	"github.com/MuhammadTawfik/notifications/queue_manager"
	"github.com/streadway/amqp"
	"log"
)

func startOne(consumer_id int, ch *amqp.Channel, queue_name string) {

	msgs, err := ch.Consume(
		queue_name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	func() {
		for d := range msgs {
			var notf dispatcher.Notification
			json.Unmarshal([]byte(d.Body), &notf)
			dispatcher.GetDispatcher(notf.Typ).Dispatch(&notf)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func StartMany(count int, server_url string, queue_name string) {

	for i := 1; i <= count; i++ {
		_, ch := queue_manager.GetChannel(server_url)
		dataQueue := queue_manager.GetQueue(queue_name, ch)
		go startOne(i, ch, dataQueue.Name)
	}

	// defer conn.Close()
	// defer ch.Close()
}
