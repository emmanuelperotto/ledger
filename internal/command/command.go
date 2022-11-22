package command

type Command interface {
    // AggregateType returns the type of the aggregate that the command can be
    // handled by.
    AggregateType() string

    // CommandType returns the type of the command.
    CommandType() string
}
