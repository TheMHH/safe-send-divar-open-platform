// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import ()

type Transaction struct {
	ID               int32
	BuyerAddress     string
	SellerAddress    string
	TransactionToken string
}
