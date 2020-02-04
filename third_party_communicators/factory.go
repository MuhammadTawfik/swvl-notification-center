package third_party_communicators

import (
	"fmt"
	"log"
)

const url = "amqp://guest:guest@rabbitmq"
const max_priority = 10
const sms_queue_name = "sms_processed_notifications"
const pn_queue_name = "pn_processed_notifications"

type Communicator interface {
	StartOne(consumer_id int)
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
