package main

import (
    "context"
    "database/sql"
    repository "github.com/emmanuelperotto/ledger/internal/repository/sqlc"
    "log"
)

func main() {
    log.Println("Starting Message Relay service")

    db, err := sql.Open("postgres", "postgresql://user:example@localhost:5432/ledger?sslmode=disable")
    if err != nil {
        log.Panic("can't connect to DB")
    }

    if err := db.Ping(); err != nil {
        log.Panic("ping DB didn't work", err)
    }

    queries := repository.New(db)

    events, err := queries.ListEvents(context.Background())
    if err != nil {
        return
    }

    log.Println("events returned")
    for _, event := range events {
        log.Println(event)
    }
}
