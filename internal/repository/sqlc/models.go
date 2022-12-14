// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package repository

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Event struct {
	ID         int32
	EventType  string
	EntityType string
	EntityID   string
	EventData  json.RawMessage
	CreatedAt  time.Time
}

type Outbox struct {
	EventID       int32
	EventType     string
	AggregateType string
	AggregateID   string
	Payload       json.RawMessage
	CreatedAt     sql.NullTime
}
