package controller

import (
	"admin-backend/models"
	"admin-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	NotiService services.NotificationService
}

func (ns *NotificationController) CreateNotification(c *gin.Context) {
	var request models.NotificationRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	notification := models.Notification{
		Template:     request.Template,
		CampaignType: request.CampaignType,
		Content:      request.Content,
	}

	err = ns.NotiService.CreateNotification(c, &notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}

	c.JSON(200, notification)
}

func (ns *NotificationController) GetAllNotification(c *gin.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	notifications, total, err := ns.NotiService.GetAllNotification(c, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	response := models.NotificationResponse{
		Data:  notifications,
		Total: total,
	}

	c.JSON(http.StatusOK, response)
}
