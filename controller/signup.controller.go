package controller

import (
	"admin-backend/models"
	"admin-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	UserService services.UserService
}

var secret = "secret"
var accessTokenExpiry = 1
var refreshTokenExpiry = 24

func (us *SignupController) SignupController(c *gin.Context) {
	var request models.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	_, err = us.UserService.GetUserByEmail(c, request.Email)

	if err == nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	err = us.UserService.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := us.UserService.CreateAccessToken(&user, secret, accessTokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	refreshToken, err := us.UserService.CreateRefreshToken(&user, secret, refreshTokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	signupResponse := models.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
