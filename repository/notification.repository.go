package repository

import (
	"admin-backend/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

type NotificationRepository interface {
	Create(ctx context.Context, notification *models.Notification) error
	GetAllNotification(ctx context.Context, page int, pageSize int) ([]models.Notification, int64, error)
}

func NewNotificationRepository(db *gorm.DB) *notificationRepository {
	return &notificationRepository{
		db: db,
	}
}
func (r *notificationRepository) Create(ctx context.Context, notification *models.Notification) error {
	notification.CreatedAt = time.Now()
	notification.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Preload("Campaigns").Create(notification).Error
}

func (r *notificationRepository) GetAllNotification(ctx context.Context, page int, pageSize int) ([]models.Notification, int64, error) {
	var notifications []models.Notification
	var total int64

	// Count total notifications
	if err := r.db.WithContext(ctx).Model(&models.Notification{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Retrieve notifications with pagination and preload associated campaigns
	if err := r.db.WithContext(ctx).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Preload("Campaign").
		Find(&notifications).Error; err != nil {
		return nil, 0, err
	}

	return notifications, total, nil
}
