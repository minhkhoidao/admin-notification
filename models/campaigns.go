package models

import (
	"time"
)

type CampaignType string
type CampaignStatus string

const (
	ManualCampaign    CampaignType = "Manual"
	AutomaticCampaign CampaignType = "Automatic"

	CampaignDraft      CampaignStatus = "Draft"
	CampaignActive     CampaignStatus = "Active"
	CampaignQueue      CampaignStatus = "Queue"
	CampaignProcessing CampaignStatus = "Processing"
	CampaignInactive   CampaignStatus = "Inactive"
	CampaignStopped    CampaignStatus = "Stopped"
)

type Campaign struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     *time.Time     `json:"deleted_at" gorm:"index"`
	Name          string         `json:"name"`
	CampaignType  CampaignType   `json:"campaign_type"`
	Status        CampaignStatus `json:"status"`
	StartAt       *time.Time     `json:"start_at"`
	EndAt         *time.Time     `json:"end_at"`
	Notifications []Notification `json:"notifications" gorm:"foreignKey:CampaignID"` // Ensure that GORM preloads notifications
	Priority      int            `json:"priority"`
	EventType     string         `json:"event_type"`
}
