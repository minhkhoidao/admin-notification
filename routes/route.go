package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(timeout time.Duration, db *gorm.DB) *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	publicRouter := router.Group("/api")
	NewSignupRoute(timeout, publicRouter, db)
	NewLoginRoute(timeout, publicRouter, db)
	NewRefreshTokenRoute(timeout, publicRouter, db)
	return router
}
