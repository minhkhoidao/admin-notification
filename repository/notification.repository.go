package repository

import (
	"admin-backend/models"
	"context"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

// Create a new notification
func (r *NotificationRepository) CreateNotification(ctx context.Context, notification *models.Notification) error {
	return r.db.WithContext(ctx).Create(notification).Error
}

func (r *NotificationRepository) GetAllNotifications(ctx context.Context) ([]models.Notification, error) {
	var notifications []models.Notification
	if err := r.db.WithContext(ctx).
		Preload("Campaign"). // Preload associated campaigns
		Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

// Get all notifications for a specific campaign
func (r *NotificationRepository) GetNotificationsByCampaignID(ctx context.Context, campaignID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	if err := r.db.WithContext(ctx).Where("campaign_id = ?", campaignID).Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

// Get all pending notifications for a specific campaign
func (r *NotificationRepository) GetPendingNotificationsByCampaignID(ctx context.Context, campaignID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	if err := r.db.WithContext(ctx).
		Where("campaign_id = ? AND status = ?", campaignID, models.NotificationPending).
		Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

// Update the status of a notification
func (r *NotificationRepository) UpdateNotificationStatus(ctx context.Context, notification *models.Notification, status models.NotificationStatus) error {
	notification.Status = status
	return r.db.WithContext(ctx).Save(notification).Error
}

// Get a notification by ID
func (r *NotificationRepository) GetNotificationByID(ctx context.Context, notificationID uint) (*models.Notification, error) {
	var notification models.Notification
	if err := r.db.WithContext(ctx).First(&notification, notificationID).Error; err != nil {
		return nil, err
	}
	return &notification, nil
}

// Update notification to associate it with a campaign
func (r *NotificationRepository) AssociateWithCampaign(ctx context.Context, notificationID uint, campaignID uint) error {
	// Update the CampaignID for the notification
	return r.db.WithContext(ctx).Model(&models.Notification{}).
		Where("id = ?", notificationID).
		Update("campaign_id", campaignID).Error
}
