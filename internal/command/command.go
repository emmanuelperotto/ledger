package command

import (
    "github.com/emmanuelperotto/ledger/internal"
    "github.com/google/uuid"
)

type (
    Command interface {
        // AggregateID returns the ID of the aggregate that the command should be
        // handled by.
        AggregateID() uuid.UUID

        // AggregateType returns the type of the aggregate that the command can be
        // handled by.
        AggregateType() internal.AggregateType

        // CommandType returns the type of the command.
        CommandType() internal.CommandType
    }
)
