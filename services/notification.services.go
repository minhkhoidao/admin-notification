package services

import (
	"admin-backend/models"
	"admin-backend/repository"
	"context"
	"time"
)

type NotificationService interface {
	CreateNotification(ctx context.Context, notification *models.Notification) error
	GetAllNotification(ctx context.Context, page int, pageSize int) ([]models.Notification, int64, error)
}

type notificationService struct {
	notificationRepo repository.NotificationRepository
	contextTimeout   time.Duration
}

func NewNotificationService(notificationRepo repository.NotificationRepository, contextTimeout time.Duration) NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
		contextTimeout:   contextTimeout,
	}
}

func (s *notificationService) CreateNotification(ctx context.Context, notification *models.Notification) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	return s.notificationRepo.Create(ctx, notification)
}

func (s *notificationService) GetAllNotification(ctx context.Context, page int, pageSize int) ([]models.Notification, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	return s.notificationRepo.GetAllNotification(ctx, page, pageSize)
}
