// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO outbox (
    event_id, event_type, aggregate_type, aggregate_id, payload
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING event_id, event_type, aggregate_type, aggregate_id, payload, created_at
`

type CreateEventParams struct {
	EventID       uuid.UUID
	EventType     string
	AggregateType string
	AggregateID   string
	Payload       pqtype.NullRawMessage
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Outbox, error) {
	row := q.db.QueryRowContext(ctx, createEvent,
		arg.EventID,
		arg.EventType,
		arg.AggregateType,
		arg.AggregateID,
		arg.Payload,
	)
	var i Outbox
	err := row.Scan(
		&i.EventID,
		&i.EventType,
		&i.AggregateType,
		&i.AggregateID,
		&i.Payload,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEvents = `-- name: DeleteEvents :exec
DELETE FROM outbox WHERE event_id = ANY($1::uuid[])
`

func (q *Queries) DeleteEvents(ctx context.Context, eventIds []uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEvents, pq.Array(eventIds))
	return err
}

const listEvents = `-- name: ListEvents :many
SELECT event_id, event_type, aggregate_type, aggregate_id, payload, created_at FROM "outbox" ORDER BY created_at
LIMIT 50
`

func (q *Queries) ListEvents(ctx context.Context) ([]Outbox, error) {
	rows, err := q.db.QueryContext(ctx, listEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Outbox{}
	for rows.Next() {
		var i Outbox
		if err := rows.Scan(
			&i.EventID,
			&i.EventType,
			&i.AggregateType,
			&i.AggregateID,
			&i.Payload,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}