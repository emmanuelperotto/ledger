package aggregate

import (
    "errors"
    "github.com/emmanuelperotto/ledger/internal/command"
    "github.com/emmanuelperotto/ledger/internal/event"
)

var (
    AccountType AggrType = "account"
)

type (
    Account struct {
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
            event.NewAccountCreatedContent(cmd.Email, cmd.Balance),
        }

        return events, nil
    default:
        return nil, errors.New("invalid command")
    }
}

func (a *Account) ApplyEvent(e event.Event) error {
    switch e.EventType() {
    case event.AccountCreated:
        if content, ok := e.Data().(event.AccountCreatedContent); ok {
            a.Email = content.Email
            a.Balance = content.Balance
            return nil
        }

        return errors.New("unprocessable event data")
    default:
        return nil
    }
}
