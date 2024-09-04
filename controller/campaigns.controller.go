package controller

import (
	"admin-backend/models"
	"admin-backend/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	campaignService     *services.CampaignService
	notificationService *services.NotificationService
}

func NewCampaignHandler(campaignService *services.CampaignService, notificationService *services.NotificationService) *CampaignHandler {
	return &CampaignHandler{
		campaignService:     campaignService,
		notificationService: notificationService,
	}
}

// Create a new campaign and associate selected notifications
func (h *CampaignHandler) CreateCampaign(c *gin.Context) {
	var req struct {
		Name            string    `json:"name"`
		CampaignType    string    `json:"campaign_type"`
		EventType       string    `json:"event_type"`
		StartAt         time.Time `json:"start_at"`
		EndAt           time.Time `json:"end_at"`
		Priority        int       `json:"priority"`
		NotificationIDs []uint    `json:"notification_ids"` // List of previously created notification IDs
	}

	// Parse the incoming request JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new Campaign struct
	campaign := &models.Campaign{
		Name:         req.Name,
		CampaignType: models.CampaignType(req.CampaignType),
		Status:       models.CampaignQueue,
		StartAt:      &req.StartAt,
		EndAt:        &req.EndAt,
		Priority:     req.Priority,
		EventType:    req.EventType,
	}

	// Save the campaign using the service
	if err := h.campaignService.CreateCampaign(c.Request.Context(), campaign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Associate selected notifications with the created campaign
	for _, notificationID := range req.NotificationIDs {
		if err := h.notificationService.AssociateNotificationWithCampaign(c.Request.Context(), notificationID, campaign.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to associate notification with campaign"})
			return
		}
	}

	// Fetch the campaign with its associated notifications
	campaignWithNotifications, err := h.campaignService.GetCampaignWithNotifications(c.Request.Context(), campaign.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load notifications for the campaign"})
		return
	}

	// Return the created campaign including associated notifications
	c.JSON(http.StatusOK, campaignWithNotifications)
}

func (h *CampaignHandler) GetCampaignWithNotifications(c *gin.Context) {
	// Parse campaign ID from URL parameter
	campaignIDStr := c.Param("id")
	campaignID, err := strconv.ParseUint(campaignIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	// Call the service to get the campaign with notifications
	campaign, err := h.campaignService.GetCampaignWithNotifications(c.Request.Context(), uint(campaignID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, campaign)
}

func (h *CampaignHandler) GetAllCampaigns(c *gin.Context) {
	// Fetch all campaigns with notifications
	campaigns, err := h.campaignService.GetAllCampaigns(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of campaigns as JSON
	c.JSON(http.StatusOK, campaigns)
}
