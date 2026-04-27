package models

import (
	"time"
)


type WebhookEvent struct {
	ID             string    `bson:"_id" json:"id"`
	Provider       string    `bson:"provider" json:"provider"`
	EventType      string    `bson:"event_type" json:"eventType"`
	EventReference string    `bson:"event_reference" json:"eventReference"`
	Payload        map[string]interface{} `bson:"payload" json:"payload"`
	Processed      bool      `bson:"processed" json:"processed"`
	ProcessedAt    *time.Time `bson:"processed_at,omitempty" json:"processedAt,omitempty"`
	CreatedAt      time.Time `bson:"created_at" json:"createdAt"`
}
