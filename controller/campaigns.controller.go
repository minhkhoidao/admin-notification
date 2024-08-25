package controller

import (
	"admin-backend/models"
	"admin-backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CampaignsController struct {
	CampaignsService services.CampaignsService
}

func (ct *CampaignsController) CreateCampaignController(c *gin.Context) {
	log.Println("test")

	var request models.CampaignsRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, models.ErrorResponse{Message: err.Error()})
		return
	}
	campaign := models.Campaigns{
		Template:       request.Template,
		Type:           request.Type,
		Event:          request.Event,
		Content:        request.Content,
		SaveTemplate:   request.SaveTemplate,
		NotificationID: request.NotificationID,
	}

	err = ct.CampaignsService.CreateNewCampaign(c, &campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaign)
}
