package main

import (
    "fmt"
    "github.com/MuhammadTawfik/notifications/front_line_consumers"
    "github.com/MuhammadTawfik/notifications/sender_simulator"
    "github.com/MuhammadTawfik/notifications/third_party_communicators"
    // "time"
    // "github.com/MuhammadTawfik/notifications/locker"
)

func main() {
    // m := third_party_integrations.GetService("kjfdkffdf")
    // m.Send("sdlkflskdf", "dsfhkjshdfkjsdhfkjhds")
    fmt.Println("hello swvl")
    // go front_line_consumers.StartOne(1111)
    // go front_line_consumers.StartOne(2222)
    var url = "amqp://guest:guest@rabbitmq"
    var queue_name = "notification_requests"
    var max_priority = 10
    var sms_queue_name = "sms_processed_notifications"
    var pn_queue_name = "pn_processed_notifications"
    sms_comm := third_party_communicators.GetCommunicator("sms")
    sms_comm.StartMany(2, url, max_priority, sms_queue_name)
    pn_comm := third_party_communicators.GetCommunicator("push_notification")
    pn_comm.StartMany(2, url, max_priority, pn_queue_name)

    front_line_consumers.StartMany(2, url, queue_name)
    sender_simulator.StartMany(2, url, queue_name)

    var a string
    fmt.Scanln(&a)
    // l := locker.CreateNewLock()
    // fmt.Println(l.GetData)
    // fmt.Println("locking")
    // l.Lock("something")
    // // fmt.Println(l.GetData)
    // fmt.Println("**************************")
    // v := l.IsLocked("something")
    // fmt.Println(v)
    // fmt.Println(l.Data)
    // fmt.Println("**************************")
    // t := l.IsLocked("something1111")
    // fmt.Println(t)
    // // fmt.Println(l.Data)
    // // fmt.Println("**************************")
    // l.LockFor("something1111", 2*time.Second)
    // // fmt.Println(l.Data)
    // // t = l.IsLocked("")
    // t = l.IsLocked("something1111")
    // fmt.Println("first time")
    // fmt.Println(t)
    // fmt.Println("**************************")
    // time.Sleep(3 * time.Second)
    // t = l.IsLocked("something1111")
    // fmt.Println(t)
    // fmt.Println(l.Data)
    // var a string
    // fmt.Scanln(&a)
}
