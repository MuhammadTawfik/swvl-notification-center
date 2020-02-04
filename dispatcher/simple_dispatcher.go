package dispatcher

import (
	"encoding/json"
	"github.com/MuhammadTawfik/notifications/queue_manager"
	"math/rand"
	"time"
)

// type Notification struct {
// 	Typ        string
// 	Body       string
// 	UserID     string
// 	CreatedAt  int64
// 	SendBefore time.Duration
//	Priority   int
// 	Counter    int
// }

const url = "amqp://guest:guest@rabbitmq"
const sms_queue_name = "sms_processed_notifications"
const pn_queue_name = "pn_processed_notifications"
const max_priority = 10

type SimpleDispatcher struct{}

func (s SimpleDispatcher) Dispatch(notification *Notification) {
	// this where we should adjust the priorities, based on some algorithm
	// but now I will just make it random number
	notification.Priority = uint8(rand.Intn(10))
	prepare_user_id(notification)
	insert_to_sender(notification)

}

func insert_to_sender(notification *Notification) {
	conn, ch := queue_manager.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	smsDataQueue := queue_manager.GetPQueue(sms_queue_name, max_priority, ch)
	pnDataQueue := queue_manager.GetPQueue(pn_queue_name, max_priority, ch)

	data, _ := json.Marshal(notification)

	if notification.Typ == "sms" {
		queue_manager.Publish(ch, smsDataQueue.Name, data, notification.Priority)

	} else {
		queue_manager.Publish(ch, pnDataQueue.Name, data, notification.Priority)
	}

}

func prepare_user_id(notification *Notification) {
	// this also a fake function which going to return the required
	//tokens needed to to communicate with the third party service
	if notification.Typ == "sms" {
		notification.UserID = get_user_mobile(notification.UserID)

	} else {
		notification.UserID = get_user_notification_token(notification.UserID)
	}
}

func get_user_notification_token(user_id string) string {
	// this is simulation for accessing the database and getting his number
	time.Sleep(700 * time.Millisecond)
	return user_id
}

func get_user_mobile(user_id string) string {
	// this is simulation for accessing the database and getting his number
	time.Sleep(700 * time.Millisecond)
	return user_id
}
