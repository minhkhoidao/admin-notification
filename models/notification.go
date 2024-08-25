package models

type TemplateType string

const (
	NewUserPromotion TemplateType = "new_user_promotion"
	PromotionPush    TemplateType = "promotion_push"
	OrderOnItsWay    TemplateType = "order_on_its_way"
	OrderArriving    TemplateType = "order_arriving"
)

type Notification struct {
	ID           int          `json:"id" gorm:"primaryKey"`
	Template     TemplateType `json:"template"`
	CampaignType string       `json:"campaign_type"`
	Content      string       `json:"content"`
}

type NotificationRequest struct {
	Template     TemplateType `json:"template"`
	CampaignType string       `json:"campaign_type"`
	Content      string       `json:"content"`
}

type NotificationResponse struct {
	Data  []Notification `json:"data"`
	Total int64          `json:"total"`
}
