package repository

import (
	"moniVestAPI/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserCollection() *mongo.Collection {
	return config.UserCollection
}