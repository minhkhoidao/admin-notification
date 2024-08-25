package controller

import (
	"admin-backend/models"
	"admin-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CampaignsController struct {
	CampaignsService services.CampaignsService
}

func (ct *CampaignsController) CreateNewCampaign(c *gin.Context) {
	var request models.CampaignsRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, models.ErrorResponse{Message: err.Error()})
	}
}
