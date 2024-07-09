package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"schedule/internal/domain/dto"
	"schedule/internal/domain/services"
	"schedule/internal/presentation/utils"
	"schedule/middleware"
	"strconv"
)

type ClassController struct {
	classService *services.ClassService
	userService  *services.UserService
}

func NewClassController(classService *services.ClassService, userService *services.UserService) *ClassController {
	return &ClassController{classService: classService, userService: userService}
}

func (controller *ClassController) GetClassById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Response(w, "Invalid class ID", http.StatusBadRequest)
		return
	}

	class, err := controller.classService.GetClassById(id)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, class, http.StatusOK)
}

func (controller *ClassController) GetAllClassesByPerson(w http.ResponseWriter, r *http.Request) {
	personIdStr := r.URL.Query().Get("personId")
	personId, err := strconv.Atoi(personIdStr)
	if err != nil {
		utils.Response(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	username := r.Context().Value("username").(string)
	roles := r.Context().Value("roles").([]string)

	if !controller.isAuthorized(username, personId, roles) {
		utils.Response(w, "Недостаточно прав", http.StatusUnauthorized)
		return
	}

	classes, err := controller.classService.GetAllClassesByPerson(personId)
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

	class, err := controller.classService.CreateClass(&createClassDto)
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

	err = controller.classService.UpdateClass(&classDto)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, "Class updated successfully", http.StatusOK)
}

func (controller *ClassController) DeleteClass(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Response(w, "Invalid class ID", http.StatusBadRequest)
		return
	}

	err = controller.classService.DeleteClass(id)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, "Class deleted successfully", http.StatusOK)
}

func (controller *ClassController) isAuthorized(username string, personId int, roles []string) bool {
	user, err := controller.userService.GetUserByUsername(username)

	if user == nil || err != nil {
		return false
	}

	if user.Id == personId {
		return true
	}

	for _, role := range roles {
		if role == "ROLE_ADMIN" {
			return true
		}
	}
	return false
}

func (controller *ClassController) SetupRoutes(router *mux.Router) {
	router.Handle("/classes", middleware.JwtAuth(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			controller.GetClassById(w, r)
		},
	))).Methods("GET")

	router.Handle("/persons/classes", middleware.JwtAuth(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			controller.GetAllClassesByPerson(w, r)
		},
	))).Methods("GET")

	router.Handle("/classes", middleware.JwtAuth(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			controller.CreateClass(w, r)
		},
	))).Methods("POST")

	router.Handle("/classes", middleware.JwtAuth(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			controller.UpdateClass(w, r)
		},
	))).Methods("PUT")

	router.Handle("/classes", middleware.JwtAuth(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			controller.DeleteClass(w, r)
		},
	))).Methods("DELETE")
}
