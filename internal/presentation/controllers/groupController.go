package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"schedule/internal/presentation/utils"
	"schedule/kafka"
)

type GroupController struct {
	Producer *kafka.Producer
}

func NewGroupController(producer *kafka.Producer) *GroupController {
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
		utils.Response(w, fmt.Sprintf("Failed to send message to Kafka: %v", err), http.StatusInternalServerError)
		return
	}

	utils.Response(w, "Message sent successfully", http.StatusOK)
}

func (controller *GroupController) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		controller.SendMessageToKafka(w, r)
	}).Methods("POST")
}
