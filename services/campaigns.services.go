package services

import (
	"admin-backend/models"
	"admin-backend/repository"
	"context"
	"time"
)

type CampaignsService interface {
	CreateNewCampaign(ctx context.Context, campaigns *models.Campaigns) error
	GetAllCampaigns(ctx context.Context, page int, pageSize int) ([]models.Campaigns, int64, error)
}

type campaignService struct {
	cr           repository.CampaignRepository
	timeDuration time.Duration
}

func NewCampaignService(cr repository.CampaignRepository, timeDuration time.Duration) *campaignService {
	return &campaignService{
		cr:           cr,
		timeDuration: timeDuration,
	}
}

func (s *campaignService) CreateNewCampaign(ctx context.Context, campaigns *models.Campaigns) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeDuration)
	defer cancel()
	return s.cr.CreateNewCampaign(ctx, campaigns)
}

func (s *campaignService) GetAllCampaigns(ctx context.Context, page int, pageSize int) ([]models.Campaigns, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeDuration)
	defer cancel()
	return s.cr.GetAllCampaign(ctx, page, pageSize)
}
