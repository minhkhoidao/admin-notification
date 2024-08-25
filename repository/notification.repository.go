package repository

import (
	"admin-backend/models"
	"context"

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
	return r.db.WithContext(ctx).Create(notification).Error
}

func (r *notificationRepository) GetAllNotification(ctx context.Context, page int, pageSize int) ([]models.Notification, int64, error) {
	var notifications []models.Notification
	var total int64

	// Count the total number of notifications
	err := r.db.WithContext(ctx).Model(&models.Notification{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Fetch the notifications with pagination
	err = r.db.WithContext(ctx).Limit(pageSize).Offset((page - 1) * pageSize).Find(&notifications).Error
	if err != nil {
		return nil, 0, err
	}

	return notifications, total, nil
}
