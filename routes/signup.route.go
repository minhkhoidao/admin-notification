package routes

import (
	"admin-backend/controller"
	"admin-backend/repository"
	"admin-backend/services"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SignupController struct {
	services services.UserService
}

func NewSignupRoute(timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	controller := controller.SignupController{
		UserService: services.NewService(ur, timeout),
	}
	group.POST("/signup", controller.SignupController)
}
