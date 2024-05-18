-- name: insertTransaction :one
INSERT INTO transactions (buyer_address, seller_address, transaction_token)
VALUES ($1, $2, $3)
RETURNING *;
