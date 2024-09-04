package services

import (
	"admin-backend/mqtt"
	"admin-backend/repository"
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type SchedulerService struct {
	notificationService *repository.NotificationRepository
	mqttService         *mqtt.MQTTService
}

func NewSchedulerService(notificationService *repository.NotificationRepository, mqttService *mqtt.MQTTService) *SchedulerService {
	return &SchedulerService{
		notificationService: notificationService,
		mqttService:         mqttService,
	}
}

// Start the cron job to check for notifications to send every minute
func (s *SchedulerService) StartCronJob() {
	c := cron.New()

	// Run the job every minute
	c.AddFunc("@every 1m", func() {
		fmt.Println("Checking for pending notifications...")
		s.SendPendingNotifications(context.Background())
	})

	// Start the cron scheduler
	c.Start()
}

// SendPendingNotifications finds notifications that need to be sent
func (s *SchedulerService) SendPendingNotifications(ctx context.Context) {
	// Get pending notifications (e.g., status = "Pending")
	notifications, err := s.notificationService.GetPendingNotifications(ctx)
	if err != nil {
		fmt.Printf("Error fetching pending notifications: %v\n", err)
		return
	}

	for _, notification := range notifications {
		// Simulate sending notification (replace with actual sending logic)
		fmt.Printf("Sending notification: %s\n", notification.Content)

		// Mark notification as sent
		notification.SentAt = &time.Time{}
		if err := s.notificationService.MarkAsSent(ctx, &notification); err != nil {
			fmt.Printf("Failed to mark notification as sent: %v\n", err)
			continue
		}

		// Send MQTT event
		message := fmt.Sprintf("Notification sent: %s", notification.Content)
		if err := s.mqttService.Publish(message); err != nil {
			fmt.Printf("Failed to publish MQTT message: %v\n", err)
		} else {
			fmt.Printf("Published MQTT message: %s\n", message)
		}
	}
}
