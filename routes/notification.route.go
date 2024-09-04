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
	notiRepo := repository.NewNotificationRepository(db)
	notificationService := services.NewNotificationService(notiRepo, &timeout)

	notificationHandler := controller.NewNotificationHandler(notificationService)

	group.POST("/notification", notificationHandler.CreateNotification)
	group.GET("/notification", notificationHandler.GetAllNotifications)
}
