
-- name: CreateAgency :one
INSERT INTO Agency (
  Name,Phone,Email
) VALUES (
  $1, $2, $3
)
RETURNING *;