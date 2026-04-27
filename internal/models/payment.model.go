package models

import (
	"time"
)

type PaymentStatus string

const (
	PayInitialized PaymentStatus = "initialized"
	PaySuccessful  PaymentStatus = "successful"
	PayFailed      PaymentStatus = "failed"
	PayAbandoned   PaymentStatus = "abandoned"
)

type Payment struct {
	ID                string        `bson:"_id" json:"id"`
	UserID            string        `bson:"user_id" json:"userId"`
	WalletID          string        `bson:"wallet_id" json:"walletId"`
	Provider          string        `bson:"provider" json:"provider"` // paystack
	ProviderReference string        `bson:"provider_reference" json:"providerReference"`
	Amount            int64         `bson:"amount" json:"amount"`
	Currency          string        `bson:"currency" json:"currency"`
	Channel           string        `bson:"channel" json:"channel"`
	Status            PaymentStatus `bson:"status" json:"status"`
	PaidAt            *time.Time    `bson:"paid_at,omitempty" json:"paidAt,omitempty"`
	CreatedAt         time.Time     `bson:"created_at" json:"createdAt"`
}
