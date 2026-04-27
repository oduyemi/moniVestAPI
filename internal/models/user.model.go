package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRole string
type UserStatus string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"

	StatusActive    UserStatus = "active"
	StatusSuspended UserStatus = "suspended"
	StatusClosed    UserStatus = "closed"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName      string             `bson:"first_name" json:"firstName"`
	LastName       string             `bson:"last_name" json:"lastName"`
	Email          string             `bson:"email" json:"email"`
	Phone          string             `bson:"phone" json:"phone"`
	Role           UserRole           `bson:"role" json:"role"` // default: user
	Password       string             `bson:"password" json:"-"`
	IsVerified     bool               `bson:"is_verified" json:"isVerified"`
	Status         UserStatus         `bson:"status" json:"status"`
	OTP            string             `bson:"otp,omitempty" json:"-"`
	OTPExpiresAt   time.Time          `bson:"otp_expires_at,omitempty" json:"-"`
	RefreshToken   string             `bson:"refresh_token,omitempty" json:"-"`
	CreatedAt      time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
	LastLogin      time.Time          `bson:"last_login,omitempty" json:"lastLogin,omitempty"`
}

func (u *User) SetDefaults() {
	if u.Role == "" {
		u.Role = RoleUser
	}
	if u.Status == "" {
		u.Status = StatusActive
	}
	u.CreatedAt = time.Now()
}

func CreateUserIndexes(collection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexes := []mongo.IndexModel{
		{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{"created_at": 1},
			Options: options.Index().
				SetExpireAfterSeconds(3600).
				SetPartialFilterExpression(bson.M{
					"is_verified": false,
				}),
		},
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	return err
}