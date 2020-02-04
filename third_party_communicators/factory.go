package third_party_communicators

import (
	"fmt"
	"log"
)

type Communicator interface {
	StartMany(count int, server_url string, max_priority int, queue_name string)
}

func GetCommunicator(service_type string) Communicator {
	if service_type == "sms" {
		return SmsCommunicator{}
	} else {
		return PnCommunicator{}
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
