package models

import (
	"time"

)

type SavingsType string
type ContributionFrequency string
type SavingsStatus string

const (
	SavingsMoniBank   SavingsType = "monibank"
	SavingsMoniTarget SavingsType = "monitarget"

	FreqDaily   ContributionFrequency = "daily"
	FreqWeekly  ContributionFrequency = "weekly"
	FreqMonthly ContributionFrequency = "monthly"
	FreqManual  ContributionFrequency = "manual"

	SavingsActive    SavingsStatus = "active"
	SavingsCompleted SavingsStatus = "completed"
	SavingsPaused    SavingsStatus = "paused"
	SavingsCancelled SavingsStatus = "cancelled"
)

type SavingsPlan struct {
	ID                     string                 `bson:"_id" json:"id"`
	OwnerID                string                 `bson:"owner_id" json:"ownerId"`
	WalletID               string                 `bson:"wallet_id" json:"walletId"`
	Type                   SavingsType            `bson:"type" json:"type"`
	Name                   string                 `bson:"name" json:"name"`
	TargetAmount           *int64                 `bson:"target_amount,omitempty" json:"targetAmount,omitempty"`
	CurrentAmount          int64                  `bson:"current_amount" json:"currentAmount"`
	LockUntil              *time.Time             `bson:"lock_until,omitempty" json:"lockUntil,omitempty"`
	ContributionFrequency  ContributionFrequency  `bson:"contribution_frequency" json:"contributionFrequency"`
	AutoSaveEnabled        bool                   `bson:"auto_save_enabled" json:"autoSaveEnabled"`
	Status                 SavingsStatus          `bson:"status" json:"status"`
	IsGroup                bool                   `bson:"is_group" json:"isGroup"`
	CreatedAt              time.Time              `bson:"created_at" json:"createdAt"`
	UpdatedAt              time.Time              `bson:"updated_at" json:"updatedAt"`
}