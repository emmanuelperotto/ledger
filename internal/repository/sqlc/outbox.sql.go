// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: outbox.sql

package repository

import (
	"context"
	"encoding/json"

	"github.com/lib/pq"
)

const addToOutbox = `-- name: AddToOutbox :one
INSERT INTO outbox (
    event_id, event_type, aggregate_type, aggregate_id, payload
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING event_id, event_type, aggregate_type, aggregate_id, payload, created_at
`

type AddToOutboxParams struct {
	EventID       int32
	EventType     string
	AggregateType string
	AggregateID   string
	Payload       json.RawMessage
}

func (q *Queries) AddToOutbox(ctx context.Context, arg AddToOutboxParams) (Outbox, error) {
	row := q.db.QueryRowContext(ctx, addToOutbox,
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

const deleteFromOutbox = `-- name: DeleteFromOutbox :exec
DELETE FROM outbox WHERE event_id = ANY($1::integer[])
`

func (q *Queries) DeleteFromOutbox(ctx context.Context, eventIds []int32) error {
	_, err := q.db.ExecContext(ctx, deleteFromOutbox, pq.Array(eventIds))
	return err
}

const listOutbox = `-- name: ListOutbox :many
SELECT event_id, event_type, aggregate_type, aggregate_id, payload, created_at FROM "outbox" ORDER BY event_id
LIMIT 50
`

func (q *Queries) ListOutbox(ctx context.Context) ([]Outbox, error) {
	rows, err := q.db.QueryContext(ctx, listOutbox)
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
