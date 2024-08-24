package routes

import (
	"admin-backend/controller"
	"admin-backend/repository"
	"admin-backend/services"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRefreshTokenRoute(timeout time.Duration, r *gin.RouterGroup, db *gorm.DB) {
	refreshTokenController := controller.RefreshTokenController{
		UserService: services.NewService(repository.NewUserRepository(db), timeout),
	}
	r.POST("/refresh_token", refreshTokenController.RefreshTokenController)
}
