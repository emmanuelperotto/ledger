-- name: CreateEvent :one
INSERT INTO events (
    event_type, entity_type, entity_id, event_data
) VALUES (
    $1, $2, $3, $4
 )
RETURNING *;