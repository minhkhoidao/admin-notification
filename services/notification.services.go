package services

import (
	"admin-backend/models"
	"admin-backend/repository"
	"context"
	"time"
)

type NotificationService struct {
	repo    *repository.NotificationRepository
	timeout *time.Duration
}

func NewNotificationService(repo *repository.NotificationRepository, timeout *time.Duration) *NotificationService {
	return &NotificationService{
		repo:    repo,
		timeout: timeout,
	}
}

func (s *NotificationService) GetAllNotifications(ctx context.Context) ([]models.Notification, error) {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.GetAllNotifications(ctx)
}

func (s *NotificationService) CreateNotification(ctx context.Context, notification *models.Notification) error {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.CreateNotification(ctx, notification)
}

// Get all notifications by campaign ID
func (s *NotificationService) GetNotificationsByCampaignID(ctx context.Context, campaignID uint) ([]models.Notification, error) {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.GetNotificationsByCampaignID(ctx, campaignID)
}

// Get all pending notifications by campaign ID
func (s *NotificationService) GetPendingNotificationsByCampaignID(ctx context.Context, campaignID uint) ([]models.Notification, error) {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.GetPendingNotificationsByCampaignID(ctx, campaignID)
}

// Update notification status
func (s *NotificationService) UpdateNotificationStatus(ctx context.Context, notification *models.Notification, status models.NotificationStatus) error {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.UpdateNotificationStatus(ctx, notification, status)
}

// Get a notification by ID
func (s *NotificationService) GetNotificationByID(ctx context.Context, notificationID uint) (*models.Notification, error) {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.GetNotificationByID(ctx, notificationID)
}

// Associate a notification with a campaign
func (s *NotificationService) AssociateNotificationWithCampaign(ctx context.Context, notificationID uint, campaignID uint) error {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.AssociateWithCampaign(ctx, notificationID, campaignID)
}
