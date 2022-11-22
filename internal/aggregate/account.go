package aggregate

import (
    "errors"
    "github.com/emmanuelperotto/ledger/internal/command"
    "github.com/emmanuelperotto/ledger/internal/event"
)

type (
    Account struct {
    }
)

func (a Account) ProcessCommand(cmd command.Command) ([]event.Event, error) {
    switch cmd := cmd.(type) {
    case command.CreateAccount:
        events := []event.Event{
            event.NewAccountCreated(cmd.Email, cmd.Balance),
        }

        return events, nil
    default:
        return nil, errors.New("invalid command")
    }
}
