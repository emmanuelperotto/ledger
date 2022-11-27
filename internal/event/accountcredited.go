package event

import (
    "github.com/emmanuelperotto/ledger/internal"
    "github.com/google/uuid"
)

type (
    AccountCreditedPayload struct {
        Value float64 `json:"value"`
    }

    AccountCredited struct {
        accountId uuid.UUID
        payload   AccountCreditedPayload
    }
)

func NewAccountCredited(accountId uuid.UUID, value float64) AccountCredited {
    return AccountCredited{
        accountId: accountId,
        payload: AccountCreditedPayload{
            Value: value,
        },
    }
}

func (a AccountCredited) EventType() internal.EventType {
    return internal.AccountCreditedEvent
}

func (a AccountCredited) AggregateType() internal.AggregateType {
    return internal.AccountAggregateType
}

func (a AccountCredited) AggregateID() uuid.UUID {
    return a.accountId
}

func (a AccountCredited) Data() any {
    return a.payload
}
