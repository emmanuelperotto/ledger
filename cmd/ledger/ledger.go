package main

import (
    "github.com/emmanuelperotto/ledger/internal/aggregate"
    "github.com/emmanuelperotto/ledger/internal/command"
    "github.com/google/uuid"
    "log"
)

func main() {
    account := createAccount()

    addBalance(&account, 1)
    addBalance(&account, 1)

    log.Println("final aggregate state", account)
}

func createAccount() aggregate.Account {
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
    return account
}

func addBalance(account *aggregate.Account, value float64) {
    cmd := command.NewCreditAccount(account.Id, value)

    events, err := account.ProcessCommand(cmd)
    if err != nil {
        log.Println("error crediting money to account", err)
    }

    for _, event := range events {
        if err := account.ApplyEvent(event); err != nil {
            log.Println("error applying event", err)
        }
    }
}
