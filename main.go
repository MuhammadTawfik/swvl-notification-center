package main

import (
    "fmt"
    "github.com/MuhammadTawfik/notifications/front_line_consumers"
    "github.com/MuhammadTawfik/notifications/sender_simulator"
    "github.com/MuhammadTawfik/notifications/third_party_communicators"
    "os"
    "strconv"
)

func main() {
    fmt.Println("hello swvl")
    url := os.Getenv("RABBITMQ_URL")
    queue_name := os.Getenv("NOTIFICATION_REQUEST_QUEUE_NAME")
    max_priority, _ := strconv.Atoi(os.Getenv("QUEUES_MAX_PRIORITY"))
    sms_queue_name := os.Getenv("SMS_QUEUE_NAME")
    pn_queue_name := os.Getenv("PN_QUEUE_NAME")
    front_line_consumers_count, _ := strconv.Atoi(os.Getenv("NUMBER_OF_FRONT_LINE_CONSUMERS"))
    sms_communicators_count, _ := strconv.Atoi(os.Getenv("NUMBER_OF_SMS_COMMUNICATORS"))
    pn_communicators_count, _ := strconv.Atoi(os.Getenv("NUMBER_OF_PN_COMMUNICATORS"))
    sender_simulators_count, _ := strconv.Atoi(os.Getenv("NUMBER_OF_SENDER_SIMULATOR"))
    fmt.Println(sender_simulators_count)

    sms_comm := third_party_communicators.GetCommunicator("sms")
    sms_comm.StartMany(sms_communicators_count, url, max_priority, sms_queue_name)
    pn_comm := third_party_communicators.GetCommunicator("push_notification")
    pn_comm.StartMany(pn_communicators_count, url, max_priority, pn_queue_name)
    front_line_consumers.StartMany(front_line_consumers_count, url, queue_name)
    sender_simulator.StartMany(sender_simulators_count, url, queue_name)

    // var a string
    // fmt.Scanln(&a)
}
