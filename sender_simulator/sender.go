package sender_simulator

import (
	"encoding/json"
	"github.com/MuhammadTawfik/notifications/queue_manager"
	"github.com/streadway/amqp"
	"math/rand"
	"strconv"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var notification_types = map[int]string{
	0: "sms",
	1: "push notification",
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Notification struct {
	Typ        string
	Body       string
	UserID     string
	CreatedAt  int64
	SendBefore time.Duration
	Counter    int
}

func startOne(sender_id int, ch *amqp.Channel, queue_name string) {

	ticker := time.Tick(50 * time.Millisecond)
	var i = 0
	for range ticker {
		i++
		notfiction := &Notification{
			Typ:        notification_types[rand.Intn(2)],
			Body:       RandStringBytes(rand.Intn(100)),
			UserID:     strconv.Itoa(rand.Intn(1000000000000)),
			CreatedAt:  time.Now().Unix(),
			SendBefore: 5 * time.Millisecond,
			Counter:    i,
		}

		data, _ := json.Marshal(notfiction)
		queue_manager.Publish(ch, queue_name, data, 0)
	}

}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func StartMany(count int, server_url string, queue_name string) {

	for i := 1; i <= count; i++ {
		_, ch := queue_manager.GetChannel(server_url)
		dataQueue := queue_manager.GetQueue(queue_name, ch)
		go startOne(i, ch, dataQueue.Name)
	}

	// the following lines are for some purpose, should be removed after adjusting the docker image

	_, ch := queue_manager.GetChannel(server_url)
	dataQueue := queue_manager.GetQueue(queue_name, ch)
	startOne(0, ch, dataQueue.Name)
	// defer conn.Close()
	// defer ch.Close()
}
