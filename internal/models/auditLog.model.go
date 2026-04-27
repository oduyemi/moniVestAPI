package models

import (
	"time"
)


type AuditLog struct {
	ID           string                 `bson:"_id" json:"id"`
	UserID       *string                `bson:"user_id,omitempty" json:"userId,omitempty"`
	ActorRole    string                 `bson:"actor_role" json:"actorRole"`
	Action       string                 `bson:"action" json:"action"`
	ResourceType string                 `bson:"resource_type" json:"resourceType"`
	ResourceID   string                 `bson:"resource_id" json:"resourceId"`
	Metadata     map[string]interface{} `bson:"metadata" json:"metadata"`
	IPAddress    string                 `bson:"ip_address" json:"ipAddress"`
	UserAgent    string                 `bson:"user_agent" json:"userAgent"`
	CreatedAt    time.Time              `bson:"created_at" json:"createdAt"`
}