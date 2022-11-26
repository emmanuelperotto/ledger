package event

import "github.com/emmanuelperotto/ledger/internal/aggregate"

type (
    EvtType string

    // Event is a domain event describing a change that has happened to an aggregate.
    //
    // An event struct and type name should:
    //   1) Be in past tense (CustomerMoved)
    //   2) Contain the intent (CustomerMoved vs CustomerAddressCorrected).
    //
    // The event should contain all the data needed when applying/handling it.
    Event interface {
        // EventType returns the type of the event.
        EventType() EvtType

        // AggregateType is the type of the aggregate that the event can be
        // applied to.
        AggregateType() aggregate.AggrType

        //Data returns the payload of the event
        Data() any
    }
)
