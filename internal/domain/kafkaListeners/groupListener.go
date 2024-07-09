package kafkaListeners

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"schedule/internal/domain/dto"
	"schedule/internal/domain/services"
	kafka2 "schedule/kafka"
)

type GroupListener struct {
	Consumer *kafka2.Consumer
	Service  *services.GroupService
}

func NewGroupListener(consumer *kafka2.Consumer, service *services.GroupService) *GroupListener {
	return &GroupListener{Consumer: consumer, Service: service}
}

func (listener *GroupListener) Listen(topic string) {
	err := listener.Consumer.Consume(topic, func(message *sarama.ConsumerMessage) {
		var kafkaMessage kafka2.Message
		if err := json.Unmarshal(message.Value, &kafkaMessage); err != nil {
			log.Printf("Error unmarshaling message: %v", err)

		}

		var command dto.GroupAction
		err := mapToCommand(kafkaMessage.Data, &command)
		if err != nil {
			return
		}

		switch kafkaMessage.Action {
		case "AddPersonToGroup":
			err := listener.Service.AddPersonToGroup(&command)
			if err != nil {
				return
			}

		case "RemovePersonFromGroup":
			err := listener.Service.RemovePersonFromGroup(&command)
			if err != nil {
				return
			}

		case "AddClassToGroup":
			err := listener.Service.AddClassToGroup(&command)
			if err != nil {
				return
			}

		case "RemoveClassFromGroup":
			err := listener.Service.RemoveClassFromGroup(&command)
			if err != nil {
				return
			}

		default:
			log.Printf("Unknown action: %s", kafkaMessage.Action)
		}
	})
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}
}

func mapToCommand(data interface{}, result interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, result)
}
