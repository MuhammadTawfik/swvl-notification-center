package dispatcher

import (
	"encoding/json"
	"github.com/MuhammadTawfik/notifications/queue_manager"
	"os"
	"strconv"
	"time"
)

var url = os.Getenv("RABBITMQ_URL")
var sms_queue_name = os.Getenv("SMS_QUEUE_NAME")
var pn_queue_name = os.Getenv("PN_QUEUE_NAME")
var max_priority, _ = strconv.Atoi(os.Getenv("QUEUES_MAX_PRIORITY"))

type Notification struct {
	Typ        string
	Body       string
	UserID     string
	CreatedAt  int64
	SendBefore time.Duration
	Priority   uint8
	Counter    int
}

type Dispatcher interface {
	Dispatch(notification *Notification)
}

func GetDispatcher(notification_type string) Dispatcher {
	if notification_type == "sms" {
		return SimpleSmsDispatcher{}
	} else {
		return SimplePnDispatcher{}
	}

}

func insert_to_be_sent(notification *Notification, queue_name string) {
	conn, ch := queue_manager.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	DataQueue := queue_manager.GetPQueue(queue_name, max_priority, ch)

	data, _ := json.Marshal(notification)

	queue_manager.Publish(ch, DataQueue.Name, data, notification.Priority)
}
