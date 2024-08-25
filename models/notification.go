package models

import "time"

type TemplateType string

const (
	NewUserPromotion TemplateType = "new_user_promotion"
	PromotionPush    TemplateType = "promotion_push"
	OrderOnItsWay    TemplateType = "order_on_its_way"
	OrderArriving    TemplateType = "order_arriving"
)

type Notification struct {
	ID           uint         `json:"id" gorm:"primaryKey"`
	Template     TemplateType `json:"template"`
	CampaignType CampaignType `json:"campaign_type"`
	Content      string       `json:"content"`
	CampaignID   uint         `gorm:"not null"`
	Campaign     *Campaigns   `gorm:"foreignKey:CampaignID"`
	CreatedAt    time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

type NotificationRequest struct {
	Template     TemplateType `json:"template"`
	CampaignType CampaignType `json:"campaign_type"`
	Content      string       `json:"content"`
	CampaignID   uint         `json:"campaign_id"`
}

type NotificationResponse struct {
	Data  []Notification `json:"data"`
	Total int64          `json:"total"`
}
