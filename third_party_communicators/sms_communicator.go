package third_party_communicators

import (
	"encoding/json"
	"github.com/MuhammadTawfik/notifications/dispatcher"
	"github.com/MuhammadTawfik/notifications/queue_manager"
	"github.com/MuhammadTawfik/notifications/third_party_integrations"
	"log"
)

type SmsCommunicator struct{}

// type Notification struct {
// 	Typ        string
// 	Body       string
// 	UserID     string
// 	CreatedAt  int64
// 	SendBefore time.Duration
// 	Counter    int
// 	Priority   int
// }

func (s SmsCommunicator) StartOne(consumer_id int) {

	conn, ch := queue_manager.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	dataQueue := queue_manager.GetPQueue(sms_queue_name, max_priority, ch)

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
			// log.Printf("communicatorrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr")
			// log.Printf("notf.Counter")
			// log.Printf("%d", consumer_id)
			// fmt.Println(consumer_id)
			// fmt.Println(notf.Counter)
			// fmt.Println(notf.Priority)
			third_party_integrations.SmsService{}.Send(notf.UserID, notf.Body)
			// log.Printf("communicatorrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr")
			// fmt.Println("************************************************************")
			// dot_count := bytes.Count(d.Body, []byte("."))
			// t := time.Duration(dot_count)
			// time.Sleep(t * time.Second)
			// log.Printf("Done")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
