package repository

import (
	"admin-backend/models"
	"context"

	"gorm.io/gorm"
)

type campaignRepository struct {
	db *gorm.DB
}

type CampaignRepository interface {
	CreateNewCampaign(ctx context.Context, campaign *models.Campaigns) error
	GetAllCampaign(ctx context.Context, page int, pageSize int) ([]models.Campaigns, int64, error)
}

func NewCampaignRepository(db *gorm.DB) *campaignRepository {
	return &campaignRepository{
		db: db,
	}
}

func (r *campaignRepository) CreateNewCampaign(ctx context.Context, campaign *models.Campaigns) error {
	err := r.db.WithContext(ctx).Create(&campaign).Error
	if err != nil {
		return err
	}
	return nil

}

func (r *campaignRepository) GetAllCampaign(ctx context.Context, page int, pageSize int) ([]models.Campaigns, int64, error) {
	var campaigns []models.Campaigns
	var count int64

	err := r.db.WithContext(ctx).Find(&campaigns).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Limit(pageSize).Offset((page - 1) * pageSize).Find(&campaigns).Error
	if err != nil {
		return nil, 0, err
	}

	return campaigns, count, nil
}
