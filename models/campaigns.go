package models

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
	ID       int          `json:"id" gorm:"primaryKey"`
	Template TemplateType `json:"template"`
	CampaignType
	CampaignEvent
	Content      string `json:"content"`
	SaveTemplate bool   `json:"save_template"`
}

type CampaignsRequest struct {
	Template TemplateType `json:"template"`
	CampaignType
	CampaignEvent
	Content      string `json:"content"`
	SaveTemplate bool   `json:"save_template"`
}
