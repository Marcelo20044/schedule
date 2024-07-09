package services

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"schedule/internal/domain/abstraction"
	"schedule/internal/domain/dto"
	"schedule/internal/kafka"
)

type GroupService struct {
	Repository abstraction.GroupRepositoryInterface
	Consumer   *kafka.Consumer
}

func NewGroupService(repository abstraction.GroupRepositoryInterface, brokers []string) *GroupService {
	consumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
	return &GroupService{Repository: repository, Consumer: consumer}
}

func (service *GroupService) StartConsuming() {
	err := service.Consumer.Consume("groups", func(message *sarama.ConsumerMessage) {
		var kafkaMessage kafka.Message
		if err := json.Unmarshal(message.Value, &kafkaMessage); err != nil {
			log.Printf("Error unmarshaling message: %v", err)

		}

		switch kafkaMessage.Action {
		case "AddPersonToGroup":
			var command dto.AddPersonToGroup
			if err := mapToStruct(kafkaMessage.Data, &command); err == nil {
				err := service.AddPersonToGroup(&command)
				if err != nil {
					return
				}
			}

		case "RemovePersonFromGroup":
			var command dto.RemovePersonFromGroup
			if err := mapToStruct(kafkaMessage.Data, &command); err == nil {
				err := service.RemovePersonFromGroup(&command)
				if err != nil {
					return
				}
			}

		case "AddClassToGroup":
			var command dto.AddClassToGroup
			if err := mapToStruct(kafkaMessage.Data, &command); err == nil {
				err := service.AddClassToGroup(&command)
				if err != nil {
					return
				}
			}

		case "RemoveClassFromGroup":
			var command dto.RemoveClassFromGroup
			if err := mapToStruct(kafkaMessage.Data, &command); err == nil {
				err := service.RemoveClassFromGroup(&command)
				if err != nil {
					return
				}
			}

		default:
			log.Printf("Unknown action: %s", kafkaMessage.Action)
		}
	})
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

}

func (service *GroupService) AddPersonToGroup(command *dto.AddPersonToGroup) error {
	return service.Repository.AddPersonToGroup(command.PersonId, command.GroupId)
}

func (service *GroupService) RemovePersonFromGroup(command *dto.RemovePersonFromGroup) error {
	return service.Repository.RemovePersonFromGroup(command.PersonId, command.GroupId)
}

func (service *GroupService) AddClassToGroup(command *dto.AddClassToGroup) error {
	return service.Repository.AddClassToGroup(command.ClassId, command.GroupId)
}

func (service *GroupService) RemoveClassFromGroup(command *dto.RemoveClassFromGroup) error {
	return service.Repository.RemoveClassFromGroup(command.ClassId, command.GroupId)
}

func mapToStruct(data interface{}, result interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, result)
}
