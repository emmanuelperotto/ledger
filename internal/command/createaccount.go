package command

import (
    "github.com/emmanuelperotto/ledger/internal"
    "github.com/google/uuid"
)

type (
    CreateAccount struct {
        Id      uuid.UUID
        Email   string
        Balance float64
    }
)

func NewCreateAccount(id uuid.UUID, email string, balance float64) CreateAccount {
    return CreateAccount{
        Id:      id,
        Email:   email,
        Balance: balance,
    }
}

func (c CreateAccount) AggregateType() internal.AggregateType {
    return internal.AccountAggregateType
}

func (c CreateAccount) AggregateID() uuid.UUID {
    return c.Id
}

func (c CreateAccount) CommandType() internal.CommandType {
    return internal.CreateAccountCommand
}
