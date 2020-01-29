package core

type NotificationTextHandler interface {
	GenerateMessage(event_id int, receiver int) string
}
