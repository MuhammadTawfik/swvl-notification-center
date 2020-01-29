package core

type NotificationTextHandler interface {
	GenerateMessage(event_id int, receivers []int) string
}
