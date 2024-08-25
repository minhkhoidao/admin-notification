package routes

import (
	"admin-backend/controller"
	"admin-backend/repository"
	"admin-backend/services"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewNotificationRoute(timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	nr := repository.NewNotificationRepository(db)
	controller := controller.NotificationController{
		NotiService: services.NewNotificationService(nr, timeout),
	}
	group.POST("/notification", controller.CreateNotification)
	group.GET("/notification", controller.GetAllNotification)
}
