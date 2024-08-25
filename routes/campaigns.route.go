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
	cr := repository.NewCampaignRepository(db)
	controller := controller.CampaignsController{
		CampaignsService: services.NewCampaignService(cr, timeout),
	}

	group.POST("/campaigns", controller.CreateCampaignController)
}
