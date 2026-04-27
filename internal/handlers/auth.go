package handlers

import (
	"context"
	"strings"
	"time"

	"moniVestAPI/internal/models"
	"moniVestAPI/internal/repository"
	"moniVestAPI/internal/services"
	"moniVestAPI/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


// REGISTER
func Register(c *fiber.Ctx) error {
	collection := repository.GetUserCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var body struct {
		FirstName       string
		LastName        string
		Email           string
		Password        string
		ConfirmPassword string
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	email := strings.ToLower(strings.TrimSpace(body.Email))

	if body.FirstName == "" || body.LastName == "" || email == "" || body.Password == "" || body.ConfirmPassword == "" {
		return c.Status(400).JSON(fiber.Map{"error": "All fields are required"})
	}

	if body.Password != body.ConfirmPassword {
		return c.Status(400).JSON(fiber.Map{"error": "Passwords must match"})
	}

	hashed, err := utils.HashPassword(body.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	otp, _ := services.GenerateOTP()
	expiry := time.Now().Add(5 * time.Minute)
	user := models.User{
		FirstName:    strings.TrimSpace(body.FirstName),
		LastName:     strings.TrimSpace(body.LastName),
		Email:        email,
		Password:     hashed,
		IsVerified:   false,
		OTP:          &otp,
		OTPExpiresAt: &expiry,
	}

	user.SetDefaults()

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return c.Status(400).JSON(fiber.Map{"error": "Email already exists"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "User creation failed"})
	}

	go services.SendOTPEmail(email, otp)
	return c.JSON(fiber.Map{"message": "OTP sent"})
}


// VERIFY OTP
func VerifyOTP(c *fiber.Ctx) error {
	collection := repository.GetUserCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var body struct {
		Email string
		OTP   string
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	email := strings.ToLower(strings.TrimSpace(body.Email))
	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "User not found"})
	}

	if user.IsVerified {
		return c.Status(400).JSON(fiber.Map{"error": "Already verified"})
	}

	if user.OTP == nil || *user.OTP != body.OTP {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid OTP"})
	}

	if user.OTPExpiresAt == nil || time.Now().After(*user.OTPExpiresAt) {
		return c.Status(400).JSON(fiber.Map{"error": "OTP expired"})
	}

	now := time.Now()
	_, err = collection.UpdateOne(ctx,
		bson.M{"_id": user.ID},
		bson.M{
			"$set": bson.M{
				"is_verified": true,
				"updated_at":  now,
			},
			"$unset": bson.M{
				"otp":            "",
				"otp_expires_at": "",
			},
		},
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Verification failed"})
	}

	go services.CreateDefaultWallets(user.ID)
	return c.JSON(fiber.Map{"message": "Account verified"})
}


// LOGIN
func Login(c *fiber.Ctx) error {
	collection := repository.GetUserCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var body struct {
		Email    string
		Password string
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	email := strings.ToLower(strings.TrimSpace(body.Email))
	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "User not found"})
	}

	if !utils.CheckPassword(user.Password, body.Password) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if !user.IsVerified {
		return c.Status(403).JSON(fiber.Map{"error": "Verify OTP first"})
	}

	accessToken, _ := services.GenerateAccessToken(user.ID.Hex())
	refreshToken, _ := services.GenerateRefreshToken(user.ID.Hex())
	rt := refreshToken
	now := time.Now()
	_, _ = collection.UpdateOne(ctx,
		bson.M{"_id": user.ID},
		bson.M{
			"$set": bson.M{
				"last_login":    now,
				"refresh_token": &rt,
			},
		},
	)

	return c.JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// REFRESH TOKEN
func RefreshToken(c *fiber.Ctx) error {
	collection := repository.GetUserCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var body struct {
		RefreshToken string
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	claims, err := services.ParseToken(body.RefreshToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid refresh token"})
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token payload"})
	}

	objID, _ := primitive.ObjectIDFromHex(userID)
	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil || user.RefreshToken == nil || *user.RefreshToken != body.RefreshToken {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid session"})
	}

	newAccess, _ := services.GenerateAccessToken(userID)
	return c.JSON(fiber.Map{"access_token": newAccess})
}


// RESEND OTP
func ResendOTP(c *fiber.Ctx) error {
	collection := repository.GetUserCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var body struct {
		Email string
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	email := strings.ToLower(strings.TrimSpace(body.Email))
	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "User not found"})
	}

	if user.IsVerified {
		return c.Status(400).JSON(fiber.Map{"error": "User already verified"})
	}

	if user.OTPExpiresAt != nil &&
		time.Now().Before(user.OTPExpiresAt.Add(-4*time.Minute)) {
		return c.Status(429).JSON(fiber.Map{"error": "Wait before requesting new OTP"})
	}

	otp, err := services.GenerateOTP()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate OTP"})
	}

	expiry := time.Now().Add(5 * time.Minute)
	_, err = collection.UpdateOne(ctx,
		bson.M{"_id": user.ID},
		bson.M{
			"$set": bson.M{
				"otp":            &otp,
				"otp_expires_at": &expiry,
			},
		},
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update OTP"})
	}

	go services.SendOTPEmail(email, otp)

	return c.JSON(fiber.Map{"message": "OTP resent"})
}


// LOGOUT
func Logout(c *fiber.Ctx) error {
	collection := repository.GetUserCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userID := c.Locals("user_id").(string)
	objID, _ := primitive.ObjectIDFromHex(userID)
	_, _ = collection.UpdateOne(ctx,
		bson.M{"_id": objID},
		bson.M{"$unset": bson.M{"refresh_token": ""}},
	)

	return c.JSON(fiber.Map{"message": "Logged out"})
}