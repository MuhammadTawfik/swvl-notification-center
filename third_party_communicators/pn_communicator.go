package third_party_communicators

import (
	"encoding/json"
	"github.com/MuhammadTawfik/notifications/dispatcher"
	"github.com/MuhammadTawfik/notifications/queue_manager"
	"github.com/MuhammadTawfik/notifications/third_party_integrations"
	"github.com/streadway/amqp"
	"log"
)

type PnCommunicator struct{}

func (p PnCommunicator) startOne(consumer_id int, ch *amqp.Channel, queue_name string) {

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
			third_party_integrations.FirebaseService{}.Send(notf.UserID, notf.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func (s PnCommunicator) StartMany(count int, server_url string, max_priority int, queue_name string) {
	for i := 1; i <= count; i++ {
		_, ch := queue_manager.GetChannel(server_url)
		dataQueue := queue_manager.GetPQueue(queue_name, max_priority, ch)
		go s.startOne(i, ch, dataQueue.Name)
	}

	// defer conn.Close()
	// defer ch.Close()
}
