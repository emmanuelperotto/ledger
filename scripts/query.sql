-- name: ListEvents :many
SELECT * FROM "outbox" ORDER BY created_at
LIMIT 50;

-- name: CreateEvent :one
INSERT INTO outbox (
    event_id, event_type, aggregate_type, aggregate_id, payload
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteEvents :exec
DELETE FROM outbox WHERE event_id = ANY(@event_ids::uuid[]);
