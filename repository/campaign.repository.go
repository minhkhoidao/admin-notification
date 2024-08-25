package repository

import (
	"admin-backend/models"
	"context"
	"fmt"
	"log"

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
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Recovered in CreateNewCampaign: %v", r)
		}
	}()

	if campaign.NotificationID == 0 {
		tx.Rollback()
		return fmt.Errorf("notification ID is not set")
	}

	var notification models.Notification
	if err := tx.First(&notification, campaign.NotificationID).Error; err != nil {
		// Handle error, e.g., return an error response
		tx.Rollback()
		return err
	}

	if err := tx.Create(&campaign).Error; err != nil {
		log.Printf("Error creating campaign: %v", err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
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
