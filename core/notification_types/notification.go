package core

import (
	"github.com/MuhammadTawfik/notifications/core/text_handlers"
)

type notification interface{
	// event_id int
	text_handler() core.NotificationTextHandler
	is_sms() bool // this is done normally with enums 
	time_to_send() int
	multiple_recievers() bool
	GenenrateMessage() string
	Receivers() []int
}
