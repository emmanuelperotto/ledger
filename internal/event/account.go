package event

import "github.com/emmanuelperotto/ledger/internal/aggregate"

var (
    AccountCreated EvtType = "account_created"
)

type (
    AccountCreatedContent struct {
        Email   string
        Balance float64
    }
)

func NewAccountCreatedContent(email string, balance float64) AccountCreatedContent {
    return AccountCreatedContent{
        Email:   email,
        Balance: balance,
    }
}

func (a AccountCreatedContent) EventType() EvtType {
    return "account_created"
}

func (a AccountCreatedContent) Data() any {
    return a
}

func (a AccountCreatedContent) AggregateType() aggregate.AggrType {
    return aggregate.AccountType
}
