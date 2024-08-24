package controller

import (
	"admin-backend/models"
	"admin-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	UserService services.UserService
}

func (rc *RefreshTokenController) RefreshTokenController(c *gin.Context) {
	var request models.RefreshTokenRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := rc.UserService.ExtractIDFromToken(request.RefreshToken, secret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "User not found"})
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Invalid user ID"})
		return
	}
	user, err := rc.UserService.GetUserByID(c, intID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "User not found"})
	}

	accessToken, err := rc.UserService.CreateAccessToken(user, secret, accessTokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := rc.UserService.CreateRefreshToken(user, secret, refreshTokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	refreshTokenResponse := models.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(200, refreshTokenResponse)
}
