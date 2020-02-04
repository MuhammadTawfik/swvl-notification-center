package main

import (
    "fmt"
    "github.com/MuhammadTawfik/notifications/front_line_consumers"
    "github.com/MuhammadTawfik/notifications/sender_simulator"
    "github.com/MuhammadTawfik/notifications/third_party_communicators"
)

func main() {
    // m := third_party_integrations.GetService("kjfdkffdf")
    // m.Send("sdlkflskdf", "dsfhkjshdfkjsdhfkjhds")
    fmt.Println("hello swvl")
    go sender_simulator.Send()
    go front_line_consumers.StartOne(1111)
    go front_line_consumers.StartOne(2222)
    sms_comm := third_party_communicators.GetCommunicator("sms")
    go sms_comm.StartOne(3333)
    pn_comm := third_party_communicators.GetCommunicator("push_notification")
    go pn_comm.StartOne(4444)
    var a string
    fmt.Scanln(&a)
}
