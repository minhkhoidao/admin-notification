package routes

import (
	"admin-backend/controller"
	"admin-backend/repository"
	"admin-backend/services"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewLoginRoute(timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	uc := repository.NewUserRepository(db)
	controller := controller.LoginController{
		UserServer: services.NewService(uc, timeout),
	}

	group.POST("/login", controller.LoginController)
}
