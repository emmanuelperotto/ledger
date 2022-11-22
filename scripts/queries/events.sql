-- name: CreateEvent :one
INSERT INTO events (
    "type", entity_type, entity_id, event_data
) VALUES (
    $1, $2, $3, $3
 )
RETURNING *;