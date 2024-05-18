-- name: insertTransaction :one
INSERT INTO transactions (buyer_token, seller_token, buyer_address, seller_address, transaction_token)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
