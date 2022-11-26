package command

import "github.com/emmanuelperotto/ledger/internal/aggregate"

type (
    CmdType string

    Command interface {
        // AggregateType returns the type of the aggregate that the command can be
        // handled by.
        AggregateType() aggregate.AggrType

        // CommandType returns the type of the command.
        CommandType() CmdType
    }
)
