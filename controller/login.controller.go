package controller

import (
	"admin-backend/models"
	"admin-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	UserServer services.UserService
}

func (lc *LoginController) LoginController(c *gin.Context) {
	var request models.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	user, err := lc.UserServer.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid email or password"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.UserServer.CreateAccessToken(user, secret, refreshTokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	refreshToken, err := lc.UserServer.CreateRefreshToken(user, secret, refreshTokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)

}
