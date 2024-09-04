package models

import (
	"time"
)

type NotificationStatus string

const (
	NotificationPending NotificationStatus = "Pending"
	NotificationSent    NotificationStatus = "Sent"
	NotificationFailed  NotificationStatus = "Failed"
)

type Notification struct {
	ID          uint               `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	DeletedAt   *time.Time         `json:"deleted_at" gorm:"index"`
	Content     string             `json:"content"`
	CampaignID  *uint              `json:"campaign_id"` // Foreign key
	Status      NotificationStatus `json:"status"`
	SentAt      *time.Time         `json:"sent_at,omitempty"`
	FailedAt    *time.Time         `json:"failed_at,omitempty"`
	Campaign    *Campaign          `json:"campaign" gorm:"foreignKey:CampaignID"` // Reference to Campaign
	Template    string             `json:"template"`
	Description string             `json:"description"`
}
