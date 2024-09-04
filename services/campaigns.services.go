package services

import (
	"admin-backend/models"
	"admin-backend/repository"
	"context"
	"time"
)

type CampaignService struct {
	repo    *repository.CampaignRepository
	timeout *time.Duration
}

func NewCampaignService(repo *repository.CampaignRepository, timeout *time.Duration) *CampaignService {
	return &CampaignService{repo: repo, timeout: timeout}
}

// Get all campaigns with notifications
func (s *CampaignService) GetAllCampaigns(ctx context.Context) ([]models.Campaign, error) {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.GetAllCampaigns(ctx)
}

func (s *CampaignService) GetCampaignsForProcessing(ctx context.Context) ([]models.Campaign, error) {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.GetCampaignsForProcessing(ctx)
}

func (s *CampaignService) CompleteCampaign(ctx context.Context, campaign *models.Campaign) error {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.UpdateCampaignStatus(ctx, campaign, models.CampaignStopped)
}

func (s *CampaignService) CreateCampaign(ctx context.Context, campaign *models.Campaign) error {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.CreateCampaign(ctx, campaign)
}

func (s *CampaignService) GetCampaignWithNotifications(ctx context.Context, campaignID uint) (*models.Campaign, error) {
	ctx, cancel := context.WithTimeout(ctx, *s.timeout)
	defer cancel()
	return s.repo.GetCampaignWithNotifications(ctx, campaignID)
}
