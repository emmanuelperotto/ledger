package internal

import "errors"

type (
    AggregateType string
    EventType     string
    CommandType   string
)

var (
    ErrUnprocessableEvent = errors.New("unprocessable event data")

    AccountAggregateType AggregateType = "account"

    AccountCreatedEvent  EventType = "account_created"
    AccountCreditedEvent EventType = "account_credited"
    AccountDebitedEvent  EventType = "account_debited"

    CreateAccountCommand CommandType = "create_account"
    CreditAccountCommand CommandType = "credit_account"
    DebitAccountCommand  CommandType = "debit_account"
)
