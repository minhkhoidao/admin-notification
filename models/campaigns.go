package models

import "time"

type CampaignType string
type CampaignEvent string

const (
	CampaignManual CampaignType = "campaign_manual"
	CampaignAuto   CampaignType = "campaign_auto"

	NewUserEvent        CampaignEvent = "new_user_event"
	OrderArrivingEvent  CampaignEvent = "order_arriving_event"
	OrderUpcommingEvent CampaignEvent = "order_upcomming_event"
)

type Campaigns struct {
	ID             uint          `json:"id" gorm:"primaryKey"`
	Template       TemplateType  `json:"template"`
	Type           CampaignType  `json:"campaign_type"`
	Event          CampaignEvent `json:"campaign_event"`
	Content        string        `json:"content"`
	SaveTemplate   bool          `json:"save_template"`
	NotificationID uint          `json:"notification_id"` // Foreign key for Notification
	Notification   Notification  `json:"notification" gorm:"foreignKey:NotificationID;references:ID"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type CampaignsRequest struct {
	Template       TemplateType  `json:"template"`
	Type           CampaignType  `json:"campaign_type"`
	Event          CampaignEvent `json:"campaign_event"`
	Content        string        `json:"content"`
	SaveTemplate   bool          `json:"save_template"`
	NotificationID uint          `json:"notification_id"` // Foreign key for Notification
}
