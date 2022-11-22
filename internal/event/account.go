package event

type (
    AccountCreated struct {
        Email   string
        Balance float64
    }
)

func NewAccountCreated(email string, balance float64) AccountCreated {
    return AccountCreated{
        Email:   email,
        Balance: balance,
    }
}

func (a AccountCreated) EventType() string {
    return "account_created"
}

func (a AccountCreated) Data() any {
    return a
}

func (a AccountCreated) AggregateType() string {
    return "account"
}
