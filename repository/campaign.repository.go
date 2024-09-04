package repository

import (
	"admin-backend/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *CampaignRepository {
	return &CampaignRepository{db: db}
}

// Get all campaigns with their notifications
func (r *CampaignRepository) GetAllCampaigns(ctx context.Context) ([]models.Campaign, error) {
	var campaigns []models.Campaign
	if err := r.db.WithContext(ctx).
		Preload("Notifications"). // Preload related notifications
		Find(&campaigns).Error; err != nil {
		return nil, err
	}
	return campaigns, nil
}

func (r *CampaignRepository) GetCampaignWithNotifications(ctx context.Context, campaignID uint) (*models.Campaign, error) {
	var campaign models.Campaign
	if err := r.db.WithContext(ctx).
		Preload("Notifications"). // Preload related notifications
		First(&campaign, campaignID).Error; err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (r *CampaignRepository) GetCampaignsForProcessing(ctx context.Context) ([]models.Campaign, error) {
	var campaigns []models.Campaign
	now := time.Now()

	if err := r.db.WithContext(ctx).
		Where("status = ? AND start_at <= ? AND end_at >= ?", models.CampaignActive, now, now).
		Order("priority DESC").
		Find(&campaigns).Error; err != nil {
		return nil, err
	}
	return campaigns, nil
}

func (r *CampaignRepository) UpdateCampaignStatus(ctx context.Context, campaign *models.Campaign, status models.CampaignStatus) error {
	campaign.Status = status
	return r.db.WithContext(ctx).Save(campaign).Error
}

func (r *CampaignRepository) CreateCampaign(ctx context.Context, campaign *models.Campaign) error {
	return r.db.WithContext(ctx).Create(campaign).Error
}
