package sender_simulator

import (
	"encoding/json"
	"github.com/MuhammadTawfik/notifications/queue_manager"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var notification_types = map[int]string{
	0: "sms",
	1: "push notification",
}

const url = "amqp://guest:guest@rabbitmq"
const queue_name = "notification_requests"
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Notification struct {
	Typ        string
	Body       string
	UserID     string
	CreatedAt  int64
	SendBefore time.Duration
}

func Send() {
	conn, ch := queue_manager.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	dataQueue := queue_manager.GetQueue(queue_name, ch)
	ticker := time.Tick(500 * time.Millisecond)

	for range ticker {
		notfiction := &Notification{
			Typ:        notification_types[rand.Intn(2)],
			Body:       RandStringBytes(rand.Intn(100)),
			UserID:     strconv.Itoa(rand.Intn(1000000000000)),
			CreatedAt:  time.Now().Unix(),
			SendBefore: 5 * time.Millisecond,
		}

		data, _ := json.Marshal(notfiction)

		msg := amqp.Publishing{
			Body: data,
		}

		ch.Publish(
			"",             //exchange string,
			dataQueue.Name, //key string,
			false,          //mandatory bool,
			false,          //immediate bool,
			msg)            //msg amqp.Publishing)

		log.Printf("Reading sent. Value: %v\n", msg)
	}

}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
