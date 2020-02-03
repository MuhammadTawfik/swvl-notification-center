package dispatcher

import (
	"time"
)

type Notification struct {
	Typ        string
	Body       string
	UserID     string
	CreatedAt  int64
	SendBefore time.Duration
	Priority   int
	Counter    int
}

type Dispatcher interface {
	Dispatch(notification *Notification)
}

func GetDispatcher() Dispatcher {
	return SimpleDispatcher{}
}
