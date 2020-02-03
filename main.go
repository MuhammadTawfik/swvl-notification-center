package main

import (
    "fmt"
    "github.com/MuhammadTawfik/notifications/front_line_consumers"
    "github.com/MuhammadTawfik/notifications/sender_simulator"
)

func main() {
    // m := third_party_integrations.GetService("kjfdkffdf")
    // m.Send("sdlkflskdf", "dsfhkjshdfkjsdhfkjhds")
    fmt.Println("hello swvl")
    go sender_simulator.Send()
    go front_line_consumers.StartOne(1111)
    front_line_consumers.StartOne(2222)
    var a string
    fmt.Scanln(&a)
}
