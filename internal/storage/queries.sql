-- name: InitializeTransaction :one
INSERT INTO transactions (post_token, supplier_id, demand_id)
VALUES ($1, $2, $3)
RETURNING *;
