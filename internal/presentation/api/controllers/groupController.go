package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"schedule/internal/kafka"
	"schedule/internal/presentation/utils"
)

type GroupController struct {
	Producer *kafka.Producer
}

func NewGroupController() *GroupController {
	producer, err := kafka.NewProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	return &GroupController{Producer: producer}
}

func (controller *GroupController) SendMessageToKafka(w http.ResponseWriter, r *http.Request) {
	var message kafka.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		utils.Response(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = controller.Producer.SendMessage("groups", message)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		utils.Response(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	utils.Response(w, "Message sent successfully", http.StatusOK)
}

func (controller *GroupController) SetupGroupRoutes(router *mux.Router) {
	router.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		controller.SendMessageToKafka(w, r)
	}).Methods("POST")
}
