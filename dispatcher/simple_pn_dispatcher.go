package dispatcher

import (
	"math/rand"
	"time"
)

type SimplePnDispatcher struct{}

func (s SimplePnDispatcher) Dispatch(notification *Notification) {
	// this where we should adjust the priorities, based on some algorithm
	// but now I will just make it random number
	notification.Priority = uint8(rand.Intn(10))
	set_user_token(notification)
	insert_to_be_sent(notification, pn_queue_name)

}

func set_user_token(notification *Notification) {
	// this also a fake function which going to return the required
	//tokens needed to to communicate with the third party service
	time.Sleep(700 * time.Millisecond)

}
