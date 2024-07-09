package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"schedule/internal/domain/dto"
	"schedule/internal/domain/services"
	"schedule/internal/presentation/utils"
	"strconv"
)

type ClassController struct {
	service *services.ClassService
}

func NewClassController(classService *services.ClassService) *ClassController {
	return &ClassController{service: classService}
}

func (controller *ClassController) GetClassById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Response(w, "Invalid class ID", http.StatusBadRequest)
		return
	}

	class, err := controller.service.GetClassById(id)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, class, http.StatusOK)
}

func (controller *ClassController) GetAllClassesByPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personIdStr := vars["id"]
	personId, err := strconv.Atoi(personIdStr)
	if err != nil {
		utils.Response(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	classes, err := controller.service.GetAllClassesByPerson(personId)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, classes, http.StatusOK)
}

func (controller *ClassController) CreateClass(w http.ResponseWriter, r *http.Request) {
	var createClassDto dto.CreateClassDto
	err := json.NewDecoder(r.Body).Decode(&createClassDto)
	if err != nil {
		utils.Response(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	class, err := controller.service.CreateClass(&createClassDto)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, class, http.StatusOK)
}

func (controller *ClassController) UpdateClass(w http.ResponseWriter, r *http.Request) {
	var classDto dto.UpdateClassDto
	err := json.NewDecoder(r.Body).Decode(&classDto)
	if err != nil {
		utils.Response(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = controller.service.UpdateClass(&classDto)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, "Class updated successfully", http.StatusOK)
}

func (controller *ClassController) DeleteClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Response(w, "Invalid class ID", http.StatusBadRequest)
		return
	}

	err = controller.service.DeleteClass(id)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, "Class deleted successfully", http.StatusOK)
}

func (controller *ClassController) SetupClassRoutes(router *mux.Router, classService *services.ClassService) {
	router.HandleFunc("/classes/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.GetClassById(w, r)
	}).Methods("GET")

	router.HandleFunc("/persons/{id}/classes", func(w http.ResponseWriter, r *http.Request) {
		controller.GetAllClassesByPerson(w, r)
	}).Methods("GET")

	router.HandleFunc("/classes", func(w http.ResponseWriter, r *http.Request) {
		controller.CreateClass(w, r)
	}).Methods("POST")

	router.HandleFunc("/classes", func(w http.ResponseWriter, r *http.Request) {
		controller.UpdateClass(w, r)
	}).Methods("PUT")

	router.HandleFunc("/classes/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.DeleteClass(w, r)
	}).Methods("DELETE")
}
