package models

import (
	"time"
)


type NotificationType string
type NotificationStatus string

const (
	NotifEmail NotificationType = "email"
	NotifSMS   NotificationType = "sms"
	NotifPush  NotificationType = "push"
	NotifInApp NotificationType = "in_app"

	NotifPending NotificationStatus = "pending"
	NotifSent    NotificationStatus = "sent"
	NotifFailed  NotificationStatus = "failed"
	NotifRead    NotificationStatus = "read"
)


type Notification struct {
	ID        string             `bson:"_id" json:"id"`
	UserID    string             `bson:"user_id" json:"userId"`
	Type      NotificationType   `bson:"type" json:"type"`
	Subject   string             `bson:"subject" json:"subject"`
	Message   string             `bson:"message" json:"message"`
	Status    NotificationStatus `bson:"status" json:"status"`
	SentAt    *time.Time         `bson:"sent_at,omitempty" json:"sentAt,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}