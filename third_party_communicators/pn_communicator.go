package third_party_communicators

import (
	"encoding/json"
	"fmt"
	"github.com/MuhammadTawfik/notifications/dispatcher"
	"github.com/MuhammadTawfik/notifications/queue_manager"
	"github.com/MuhammadTawfik/notifications/third_party_integrations"
	"log"
	"time"
)

type PnCommunicator struct{}

// type Notification struct {
// 	Typ        string
// 	Body       string
// 	UserID     string
// 	CreatedAt  int64
// 	SendBefore time.Duration
// 	Counter    int
// 	Priority   int
// }

func (s PnCommunicator) StartOne(consumer_id int) {

	conn, ch := queue_manager.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	dataQueue := queue_manager.GetPQueue(pn_queue_name, max_priority, ch)

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
			notification_tokens := get_user_notification_tokens(notf.UserID)
			third_party_integrations.FirebaseService{}.BulkSend(notification_tokens, notf.Body)
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

func get_user_notification_tokens(user_id string) []string {
	// this is simulation for accessing the database and getting his number
	time.Sleep(700 * time.Millisecond)
	return []string{user_id}
}
