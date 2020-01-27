package third_party_integrations

type ThirdPartyService interface{
	Send(id string, msg string)
	BulkSend(ids []string, msg string)
}

func GetService(service_name string) ThirdPartyService{
	if service_name == "sms" {
		return SmsService{}
	} else {
		return FirebaseService{}
	}
}
