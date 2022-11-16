package main

import (
    "log"
    "time"
)

func main() {
    log.Println("Starting Message Relay service")
    for {
        log.Println("fetching events to deliver...")
        time.Sleep(1 * time.Second)
        log.Println("events delivered...")
    }
}
