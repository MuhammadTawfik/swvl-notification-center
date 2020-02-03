package front_line_consumers

import (
	"encoding/json"
	"fmt"
	"github.com/MuhammadTawfik/notifications/queue_manager"
	"log"
	"time"
)

const url = "amqp://guest:guest@rabbitmq"
const queue_name = "notification_requests"

type Notification struct {
	Typ        string
	Body       string
	UserID     string
	CreatedAt  int64
	SendBefore time.Duration
	Counter    int
}

func StartOne(consumer_id int) {
	fmt.Println("started started")
	fmt.Println("started started")
	fmt.Println("started started")
	fmt.Println("started started")
	fmt.Println("started started")
	fmt.Println("started started")

	conn, ch := queue_manager.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	dataQueue := queue_manager.GetQueue(queue_name, ch)

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
			var notf Notification
			json.Unmarshal([]byte(d.Body), &notf)
			log.Printf("notf.Counter")
			log.Printf("%d", consumer_id)
			fmt.Println(consumer_id)
			fmt.Println(notf.Counter)
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
