package internal

type (
    AggregateType string
    EventType     string
    CommandType   string
)

var (
    AccountAggregateType AggregateType = "account"

    AccountCreatedEvent EventType = "account_created"

    CreateAccountCommand CommandType = "create_account"
)
