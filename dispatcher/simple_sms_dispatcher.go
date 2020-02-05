package dispatcher

import (
	"math/rand"
	"time"
)

type SimpleSmsDispatcher struct{}

func get_user_mobile(user_id string) string {
	// this is simulation for accessing the database and getting his number
	time.Sleep(700 * time.Millisecond)
	return user_id
}

func (s SimpleSmsDispatcher) Dispatch(notification *Notification) {
	// this where we should adjust the priorities, based on some algorithm
	// but now I will just make it random number
	notification.Priority = uint8(rand.Intn(10))
	set_user_mobile_number(notification)
	insert_to_be_sent(notification, sms_queue_name)

}

func set_user_mobile_number(notification *Notification) {
	// this also a fake function which going to return the required
	//mobile_number needed to to communicate with the third party service
	time.Sleep(700 * time.Millisecond)

}
