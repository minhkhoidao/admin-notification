package controller

import (
	"admin-backend/models"
	"admin-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	notificationService *services.NotificationService
}

func NewNotificationHandler(notificationService *services.NotificationService) *NotificationHandler {
	return &NotificationHandler{notificationService: notificationService}
}

// Create a new notification for a specific campaign
func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var req struct {
		Content     string `json:"content"`
		CampaignID  *uint  `json:"campaign_id,omitempty"` // CampaignID must be provided
		Template    string `json:"template"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification := &models.Notification{
		Content:     req.Content,
		CampaignID:  req.CampaignID, // Associate notification with campaign
		Status:      models.NotificationPending,
		Template:    req.Template,
		Description: req.Description,
	}

	if err := h.notificationService.CreateNotification(c.Request.Context(), notification); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notification)
}

func (h *NotificationHandler) GetNotificationsByCampaignID(c *gin.Context) {
	// Parse campaign ID from URL parameter
	campaignIDStr := c.Param("id")
	campaignID, err := strconv.ParseUint(campaignIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	notifications, err := h.notificationService.GetNotificationsByCampaignID(c.Request.Context(), uint(campaignID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

func (h *NotificationHandler) GetAllNotifications(c *gin.Context) {
	// Fetch all notifications with campaigns
	notifications, err := h.notificationService.GetAllNotifications(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of notifications as JSON
	c.JSON(http.StatusOK, notifications)
}
