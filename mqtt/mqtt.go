package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTService struct {
	client mqtt.Client
	topic  string
}

func NewMQTTService(broker string, topic string) *MQTTService {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Connected to MQTT broker")

	return &MQTTService{
		client: client,
		topic:  topic,
	}
}

func (s *MQTTService) Publish(message string) error {
	token := s.client.Publish(s.topic, 0, false, message)
	token.Wait()
	return token.Error()
}
