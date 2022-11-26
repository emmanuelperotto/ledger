package aggregate

import (
    "errors"
    "github.com/emmanuelperotto/ledger/internal"
    "github.com/emmanuelperotto/ledger/internal/command"
    "github.com/emmanuelperotto/ledger/internal/event"
    "github.com/google/uuid"
)

type (
    Account struct {
        Id      uuid.UUID
        Email   string
        Balance float64
    }
)

func NewAccount() Account {
    return Account{}
}

func (a *Account) ProcessCommand(cmd command.Command) ([]event.Event, error) {
    switch cmd := cmd.(type) {
    case command.CreateAccount:
        events := []event.Event{
            event.NewAccountCreated(cmd.AggregateID(), event.AccountCreatedPayload{
                Email:   cmd.Email,
                Balance: cmd.Balance,
            }),
        }

        return events, nil
    default:
        return nil, errors.New("invalid command")
    }
}

func (a *Account) ApplyEvent(e event.Event) error {
    switch e.EventType() {
    case internal.AccountCreatedEvent:
        if content, ok := e.Data().(event.AccountCreatedPayload); ok {
            a.Id = e.AggregateID()
            a.Email = content.Email
            a.Balance = content.Balance
            return nil
        }

        return errors.New("unprocessable event data")
    default:
        return nil
    }
}
