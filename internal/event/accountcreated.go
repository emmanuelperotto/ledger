package event

import (
    "github.com/emmanuelperotto/ledger/internal"
    "github.com/google/uuid"
)

type (
    AccountCreatedPayload struct {
        Email   string  `json:"email"`
        Balance float64 `json:"balance"`
    }

    AccountCreated struct {
        Payload AccountCreatedPayload

        AggregateId uuid.UUID
    }
)

func NewAccountCreated(aggregateId uuid.UUID, payload AccountCreatedPayload) AccountCreated {
    return AccountCreated{
        Payload:     payload,
        AggregateId: aggregateId,
    }
}

func (a AccountCreated) EventType() internal.EventType {
    return internal.AccountCreatedEvent
}

func (a AccountCreated) Data() any {
    return a.Payload
}

func (a AccountCreated) AggregateType() internal.AggregateType {
    return internal.AccountAggregateType
}

func (a AccountCreated) AggregateID() uuid.UUID {
    return a.AggregateId
}
