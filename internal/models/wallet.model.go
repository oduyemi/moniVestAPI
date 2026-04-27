package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WalletType string
type WalletStatus string

const (
	WalletMoniFlex  WalletType = "moniflex"
	WalletMoniBank  WalletType = "monibank"
	WalletMoniTarget WalletType = "monitarget"

	WalletActive WalletStatus = "active"
	WalletLocked WalletStatus = "locked"
	WalletClosed WalletStatus = "closed"
)

type Wallet struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID        primitive.ObjectID `bson:"user_id" json:"userId"` // reference to User
	WalletType    WalletType         `bson:"wallet_type" json:"walletType"`
	WalletName    string             `bson:"wallet_name" json:"walletName"`
	AccountNumber string             `bson:"account_number" json:"accountNumber"`
	Balance       float64            `bson:"balance" json:"balance"` // use numeric
	Status        WalletStatus       `bson:"status" json:"status"`
	CreatedAt     time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
}

// Optional defaults
func (w *Wallet) SetDefaults() {
	if w.Status == "" {
		w.Status = WalletActive
	}
	w.CreatedAt = time.Now()
}

func CreateWalletIndexes(collection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexes := []mongo.IndexModel{
		{
			Keys: bson.M{"user_id": 1},
		},
		{
			Keys: bson.M{
			  "user_id":     1,
			  "wallet_type": 1,
			},
			Options: options.Index().SetUnique(true),
		  },
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	return err
}

