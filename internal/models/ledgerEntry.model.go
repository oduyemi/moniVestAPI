package models

import (
	"time"
)

type EntryType string

const (
	Debit  EntryType = "debit"
	Credit EntryType = "credit"
)


type LedgerEntry struct {
	ID             string    `bson:"_id" json:"id"`
	TransactionID  string    `bson:"transaction_id" json:"transactionId"`
	WalletID       string    `bson:"wallet_id" json:"walletId"`
	EntryType      EntryType `bson:"entry_type" json:"entryType"`
	Amount         int64     `bson:"amount" json:"amount"`
	BalanceBefore  int64     `bson:"balance_before" json:"balanceBefore"`
	BalanceAfter   int64     `bson:"balance_after" json:"balanceAfter"`
	CreatedAt      time.Time `bson:"created_at" json:"createdAt"`
}