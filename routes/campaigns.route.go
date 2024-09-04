package routes

import (
	"admin-backend/controller"
	"admin-backend/repository"
	"admin-backend/services"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCampaignRoute(timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	campaignRepo := repository.NewCampaignRepository(db)
	notiRepo := repository.NewNotificationRepository(db)

	campaignService := services.NewCampaignService(campaignRepo, &timeout)
	notificationService := services.NewNotificationService(notiRepo, &timeout)

	campaignHandler := controller.NewCampaignHandler(campaignService, notificationService)

	group.POST("/campaigns", campaignHandler.CreateCampaign)
	group.GET("/campaigns", campaignHandler.GetAllCampaigns)
}
