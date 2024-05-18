// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package storage

import (
	"context"
)

const insertTransaction = `-- name: insertTransaction :one
INSERT INTO transactions (buyer_address, seller_address, transaction_token)
VALUES ($1, $2, $3)
RETURNING id, buyer_address, seller_address, transaction_token
`

type insertTransactionParams struct {
	BuyerAddress     string
	SellerAddress    string
	TransactionToken string
}

func (q *Queries) insertTransaction(ctx context.Context, arg insertTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, insertTransaction, arg.BuyerAddress, arg.SellerAddress, arg.TransactionToken)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.BuyerAddress,
		&i.SellerAddress,
		&i.TransactionToken,
	)
	return i, err
}