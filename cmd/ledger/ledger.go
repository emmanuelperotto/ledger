package main

import (
    "github.com/emmanuelperotto/ledger/internal/aggregate"
    "github.com/emmanuelperotto/ledger/internal/command"
    "github.com/google/uuid"
    "log"
)

func main() {
    cmd := command.NewCreateAccount(uuid.New(), "emmanuelperoto@gmail.com", 10.0)
    account := aggregate.NewAccount()

    events, err := account.ProcessCommand(cmd)
    if err != nil {
        log.Println("error processing account creation", err)
    }

    for _, event := range events {
        if err := account.ApplyEvent(event); err != nil {
            log.Println("error applying event", err)
        }
    }

    log.Println(account)
}
