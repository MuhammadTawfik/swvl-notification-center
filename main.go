package main

import (
    "fmt"
    "github.com/MuhammadTawfik/notifications/core/notification_types"
)

func main() {
    m := core.Welcome{}

    fmt.Println(m.GenenrateMessage(1))
}
