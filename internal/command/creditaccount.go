package command

import (
    "github.com/emmanuelperotto/ledger/internal"
    "github.com/google/uuid"
)

type (
    CreditAccount struct {
        AccountId uuid.UUID
        Value     float64
    }
)

func NewCreditAccount(accountId uuid.UUID, value float64) CreditAccount {
    return CreditAccount{
        AccountId: accountId,
        Value:     value,
    }
}

func (c CreditAccount) AggregateID() uuid.UUID {
    return c.AccountId
}

func (c CreditAccount) AggregateType() internal.AggregateType {
    return internal.AccountAggregateType
}

func (c CreditAccount) CommandType() internal.CommandType {
    return internal.CreditAccountCommand
}
