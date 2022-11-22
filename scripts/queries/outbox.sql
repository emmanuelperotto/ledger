-- name: ListOutbox :many
SELECT * FROM "outbox" ORDER BY event_id
LIMIT 50;

-- name: AddToOutbox :one
INSERT INTO outbox (
    event_id, event_type, aggregate_type, aggregate_id, payload
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteFromOutbox :exec
DELETE FROM outbox WHERE event_id = ANY(@event_ids::integer[]);
