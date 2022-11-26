package command

import "github.com/emmanuelperotto/ledger/internal/aggregate"

type (
    CreateAccount struct {
        Email   string
        Balance float64
    }
)

func NewCreateAccount(email string, balance float64) CreateAccount {
    return CreateAccount{
        Email:   email,
        Balance: balance,
    }
}

func (c CreateAccount) AggregateType() aggregate.AggrType {
    return aggregate.AccountType
}

func (c CreateAccount) CommandType() CmdType {
    return "create_account"
}
