package models

import (
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRole string
type UserStatus string

const (
	UserRoleUser  UserRole = "user"
	UserRoleAdmin UserRole = "admin"

	UserStatusActive    UserStatus = "active"
	UserStatusSuspended UserStatus = "suspended"
	UserStatusClosed    UserStatus = "closed"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName    string             `bson:"first_name" json:"firstName"`
	LastName     string             `bson:"last_name" json:"lastName"`
	Email        string             `bson:"email" json:"email"`
	Phone        string             `bson:"phone" json:"phone"`
	Role         UserRole           `bson:"role" json:"role"`
	Password     string             `bson:"password" json:"-"`
	IsVerified   bool               `bson:"is_verified" json:"isVerified"`
	Status       UserStatus         `bson:"status" json:"status"`
	OTP          *string            `bson:"otp,omitempty" json:"-"`
	OTPExpiresAt *time.Time         `bson:"otp_expires_at,omitempty" json:"-"`
	RefreshToken *string            `bson:"refresh_token,omitempty" json:"-"`

	CreatedAt time.Time  `bson:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
	LastLogin *time.Time `bson:"last_login,omitempty" json:"lastLogin,omitempty"`
}

func (u *User) SetDefaults() {
	now := time.Now()

	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	if u.Role == "" {
		u.Role = UserRoleUser
	}

	if u.Status == "" {
		u.Status = UserStatusActive
	}

	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
}

func (u *User) Touch() {
	now := time.Now()
	u.UpdatedAt = &now
}

func CreateUserIndexes(collection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexes := []mongo.IndexModel{
		{
			Keys: bson.M{"email": 1},
			Options: options.Index().
				SetUnique(true).
				SetCollation(&options.Collation{
					Locale:   "en",
					Strength: 2, 
				}),
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