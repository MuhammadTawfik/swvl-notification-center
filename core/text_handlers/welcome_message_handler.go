package core

type WelcomeMessage struct{}

func (e WelcomeMessage) GenerateMessage(event_id int, receiver int) string {
	// generate message per user
	// find user language should go in here
	return "Welcome in SWVL"
}
