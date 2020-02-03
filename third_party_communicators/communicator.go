package third_party_communicators

import (
	"encoding/json"
	"fmt"
	"github.com/MuhammadTawfik/notifications/dispatcher"
	"github.com/MuhammadTawfik/notifications/queue_manager"
	"github.com/MuhammadTawfik/notifications/third_party_integrations"
	"log"
)

const url = "amqp://guest:guest@rabbitmq"
const queue_name = "processed_notifications"
const max_priority = 10

// type Notification struct {
// 	Typ        string
// 	Body       string
// 	UserID     string
// 	CreatedAt  int64
// 	SendBefore time.Duration
// 	Counter    int
// 	Priority   int
// }

func StartOne(consumer_id int) {

	conn, ch := queue_manager.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	dataQueue := queue_manager.GetPQueue(queue_name, max_priority, ch)

	msgs, err := ch.Consume(
		dataQueue.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	func() {
		for d := range msgs {
			var notf dispatcher.Notification
			json.Unmarshal([]byte(d.Body), &notf)
			log.Printf("communicatorrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr")
			log.Printf("notf.Counter")
			log.Printf("%d", consumer_id)
			fmt.Println(consumer_id)
			fmt.Println(notf.Counter)
			fmt.Println(notf.Priority)
			third_party_integrations.GetService(notf.Typ).Send(notf.UserID, notf.Body)
			log.Printf("communicatorrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr")
			fmt.Println("************************************************************")
			// dot_count := bytes.Count(d.Body, []byte("."))
			// t := time.Duration(dot_count)
			// time.Sleep(t * time.Second)
			// log.Printf("Done")
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
