package event

// Event is a domain event describing a change that has happened to an aggregate.
//
// An event struct and type name should:
//   1) Be in past tense (CustomerMoved)
//   2) Contain the intent (CustomerMoved vs CustomerAddressCorrected).
//
// The event should contain all the data needed when applying/handling it.
type Event interface {
    // EventType returns the type of the event.
    EventType() string

    // Data is the data attached to the event.
    Data() any

    // AggregateType is the type of the aggregate that the event can be
    // applied to.
    AggregateType() string
}
