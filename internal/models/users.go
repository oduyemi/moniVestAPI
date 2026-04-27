package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName         string             `bson:"first_name" json:"firstName"`
	LastName         string             `bson:"last_name" json:"lastName"`
	Email        string             `bson:"email" json:"email"`
	Phone        string             `bson:"phone" json:"phone"`
	Role        string              `bson:"role" json:"role"` // user, admin
	Password     string             `bson:"password" json:"-"`
	IsVerified   bool               `bson:"is_verified" json:"isVerified"`
	Status         string             `bson:"status" json:"status"` // active, suspended, closed
	OTP          string             `bson:"otp,omitempty" json:"-"`
	OTPExpiresAt time.Time          `bson:"otp_expires_at,omitempty" json:"-"`
	RefreshToken string `bson:"refresh_token,omitempty" json:"-"`
	CreatedAt    time.Time          `bson:"created_at" json:"createdAt"`
	LastLogin    time.Time          `bson:"last_login" json:"lastLogin"`
}


func CreateUserIndexes(collection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	indexes := []mongo.IndexModel{
		{
			Keys: bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{"created_at": 1},
			Options: options.Index().
				SetExpireAfterSeconds(3600). // 1 hour
				SetPartialFilterExpression(bson.M{
					"is_verified": false,
				}),
		},
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	return err
}