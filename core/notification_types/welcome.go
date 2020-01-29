package core

import (
	"github.com/MuhammadTawfik/notifications/core/text_handlers"
)
type Welcome struct{
	event_id int
}


func (n Welcome) text_handler() core.NotificationTextHandler{
	return core.WelcomeMessage{}
}

func (n Welcome) is_sms() bool{
	return true
}

func (n Welcome) time_to_send() int{
	return 50
}

func (n Welcome) multiple_recievers() bool{
	return true
}

func (n Welcome) GenenrateMessage() string{
	return n.text_handler().GenerateMessage(n.event_id, n.Receivers())
}

func (n Welcome) Receivers() [] int{
	// here the logic of finding the users who should recive this message
	return [] int{1,2}
}


