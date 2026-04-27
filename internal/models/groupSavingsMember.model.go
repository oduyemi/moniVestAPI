package models

import (
	"time"
)

type MemberRole string
type MemberStatus string

const (
	MemberRoleOwner  MemberRole = "owner"
	MemberRoleAdmin  MemberRole = "admin"
	MemberRoleMember MemberRole = "member"

	MemberStatusActive  MemberStatus = "active"
	MemberStatusInvited MemberStatus = "invited"
	MemberStatusLeft    MemberStatus = "left"
	MemberStatusRemoved MemberStatus = "removed"
)


type GroupSavingsMember struct {
	ID                 string       `bson:"_id" json:"id"`
	SavingsPlanID      string       `bson:"savings_plan_id" json:"savingsPlanId"`
	UserID             string       `bson:"user_id" json:"userId"`
	Role               MemberRole   `bson:"role" json:"role"`
	ContributionTarget *int64       `bson:"contribution_target,omitempty" json:"contributionTarget,omitempty"`
	JoinedAt           time.Time    `bson:"joined_at" json:"joinedAt"`
	Status             MemberStatus `bson:"status" json:"status"`
}