package routes

import (
	"admin-backend/middlewares"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(timeout time.Duration, db *gorm.DB) *gin.Engine {
	router := gin.Default()
	publicRouter := router.Group("/")
	NewSignupRoute(timeout, publicRouter, db)
	NewLoginRoute(timeout, publicRouter, db)
	NewRefreshTokenRoute(timeout, publicRouter, db)

	protectedRouter := router.Group("/api")
	protectedRouter.Use(middlewares.JwtAuthMiddleware("secret"))
	NewNotificationRoute(timeout, protectedRouter, db)
	NewCampaignRoute(timeout, protectedRouter, db)
	return router
}
