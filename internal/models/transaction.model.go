package models

import (
	"time"
)

type TransactionType string
type TransactionStatus string

const (
	TxDeposit              TransactionType = "deposit"
	TxWithdrawal           TransactionType = "withdrawal"
	TxTransfer             TransactionType = "transfer"
	TxSavingsContribution  TransactionType = "savings_contribution"
	TxRefund               TransactionType = "refund"
	TxPenalty              TransactionType = "penalty"

	TxPending    TransactionStatus = "pending"
	TxSuccessful TransactionStatus = "successful"
	TxFailed     TransactionStatus = "failed"
	TxReversed   TransactionStatus = "reversed"
)

type Transaction struct {
	ID                  string    `bson:"_id" json:"id"` // UUID
	Reference           string    `bson:"reference" json:"reference"`
	UserID              string    `bson:"user_id" json:"userId"`
	SourceWalletID      *string   `bson:"source_wallet_id,omitempty" json:"sourceWalletId,omitempty"`
	DestinationWalletID *string   `bson:"destination_wallet_id,omitempty" json:"destinationWalletId,omitempty"`
	Type                TransactionType `bson:"type" json:"type"`
	Amount              int64     `bson:"amount" json:"amount"` // kobo
	Fee                 int64     `bson:"fee" json:"fee"`
	Status              TransactionStatus `bson:"status" json:"status"`
	Narration           string    `bson:"narration" json:"narration"`
	IdempotencyKey      *string   `bson:"idempotency_key,omitempty" json:"idempotencyKey,omitempty"`
	CreatedAt           time.Time `bson:"created_at" json:"createdAt"`
	CompletedAt         *time.Time `bson:"completed_at,omitempty" json:"completedAt,omitempty"`
}

