package core


type NotificationTextHandler interface{
	GenerateMessage(event_id int, receivers [] int) string
	// find_receivers(event_id int)
}

// func text_handler_factory(string) notification_text_handler{

// }