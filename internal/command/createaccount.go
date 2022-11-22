package command

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

func (c CreateAccount) AggregateType() string {
    return "account"
}

func (c CreateAccount) CommandType() string {
    return "create_account"
}
