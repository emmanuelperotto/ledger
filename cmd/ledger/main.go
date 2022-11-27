package main

import (
    "context"
    "database/sql"
    "encoding/json"
    "github.com/emmanuelperotto/ledger/internal/aggregate"
    "github.com/emmanuelperotto/ledger/internal/command"
    repository "github.com/emmanuelperotto/ledger/internal/repository/sqlc"
    "github.com/google/uuid"
    "log"
)

func main() {
    ctx := context.Background()
    q, db := connectDB()
    account, err := createAccount(ctx, db, q)
    if err != nil {
        log.Println("error creating account", err)
    }

    addBalance(&account, 1)
    addBalance(&account, 1)

    log.Println("final aggregate state", account)
}

func connectDB() (*repository.Queries, *sql.DB) {
    db, err := sql.Open("postgres", "postgresql://user:example@localhost:5432/ledger?sslmode=disable")
    if err != nil {
        log.Panic("can't connect to DB")
    }

    if err := db.Ping(); err != nil {
        log.Panic("ping DB didn't work", err)
    }

    return repository.New(db), db
}

func createAccount(ctx context.Context, db *sql.DB, queries *repository.Queries) (aggregate.Account, error) {
    //init tx
    tx, err := db.Begin()
    if err != nil {
        return aggregate.Account{}, err
    }
    defer func(tx *sql.Tx) {
        _ = tx.Rollback()
    }(tx)
    qtx := queries.WithTx(tx)

    cmd := command.NewCreateAccount(uuid.New(), "emmanuelperoto@gmail.com", 10.0)
    account := aggregate.NewAccount()

    events, err := account.ProcessCommand(cmd)
    if err != nil {
        log.Println("error processing account creation", err)
    }

    for _, event := range events {
        bytes, err := json.Marshal(event.Data())
        if err != nil {
            return account, err
        }

        //insert into event store
        eventDb, err := qtx.CreateEvent(ctx, repository.CreateEventParams{
            EventType:  string(event.EventType()),
            EntityType: string(event.AggregateType()),
            EntityID:   event.AggregateID().String(),
            EventData:  bytes,
        })

        if err != nil {
            return account, err
        }

        //insert into outbox
        if _, err := qtx.AddToOutbox(ctx, repository.AddToOutboxParams{
            EventID:       eventDb.ID,
            EventType:     string(event.EventType()),
            AggregateType: string(event.AggregateType()),
            AggregateID:   event.AggregateID().String(),
            Payload:       bytes,
        }); err != nil {
            return account, err
        }
        if err := account.ApplyEvent(event); err != nil {
            log.Println("error applying event", err)
        }
    }

    return account, tx.Commit()
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
