package aggregate

import (
    "errors"
    "github.com/emmanuelperotto/ledger/internal"
    "github.com/emmanuelperotto/ledger/internal/command"
    "github.com/emmanuelperotto/ledger/internal/event"
    "github.com/google/uuid"
    "log"
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
    case command.CreditAccount:
        events := []event.Event{
            event.NewAccountCredited(cmd.AccountId, cmd.Value),
        }
        return events, nil
    default:
        return nil, errors.New("invalid command")
    }
}

func (a *Account) ApplyEvent(e event.Event) error {
    log.Println("current aggregate state", a)
    switch e.EventType() {
    case internal.AccountCreatedEvent:
        if content, ok := e.Data().(event.AccountCreatedPayload); ok {
            a.Id = e.AggregateID()
            a.Email = content.Email
            a.Balance = content.Balance
            return nil
        }
        return internal.ErrUnprocessableEvent

    case internal.AccountCreditedEvent:
        if content, ok := e.Data().(event.AccountCreditedPayload); ok {
            a.Balance += content.Value
            return nil
        }
        return internal.ErrUnprocessableEvent
    default:
        return nil
    }
}
