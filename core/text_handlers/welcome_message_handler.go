package core

type WelcomeMessage struct{}

func (e WelcomeMessage) GenerateMessage(event_id int, receivers [] int) string{
	// find user language should go in here
	return "Welcome in SWVL"	
}

