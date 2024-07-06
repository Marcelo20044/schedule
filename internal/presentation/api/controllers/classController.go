package controllers

import (
	"encoding/json"
	"net/http"
	"schedule/internal/domain/dto"
	"schedule/internal/domain/services"
	"schedule/internal/presentation/utils"
	"strconv"
)

var GeClassById = func(w http.ResponseWriter, r *http.Request, classService *services.ClassService) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Response(w, "Invalid class ID", http.StatusBadRequest)
		return
	}

	class, err := classService.GetClassById(id)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, class, http.StatusOK)
}

var GetAllClassesByPerson = func(w http.ResponseWriter, r *http.Request, classService *services.ClassService) {
	personIdStr := r.URL.Query().Get("personId")
	personId, err := strconv.Atoi(personIdStr)
	if err != nil {
		utils.Response(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	classes, err := classService.GetAllClassesByPerson(personId)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, classes, http.StatusOK)
}

var CreateClass = func(w http.ResponseWriter, r *http.Request, classService *services.ClassService) {
	var createClassDto dto.CreateClassDto
	err := json.NewDecoder(r.Body).Decode(&createClassDto)
	if err != nil {
		utils.Response(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	class, err := classService.CreateClass(&createClassDto)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, class, http.StatusOK)
}

var UpdateClass = func(w http.ResponseWriter, r *http.Request, classService *services.ClassService) {
	var classDto dto.ClassDto
	err := json.NewDecoder(r.Body).Decode(&classDto)
	if err != nil {
		utils.Response(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = classService.UpdateClass(&classDto)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, "Class updated successfully", http.StatusOK)
}

var DeleteClass = func(w http.ResponseWriter, r *http.Request, classService *services.ClassService) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Response(w, "Invalid class ID", http.StatusBadRequest)
		return
	}

	err = classService.DeleteClass(id)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, "Class deleted successfully", http.StatusOK)
}

var SignUp = func(w http.ResponseWriter, r *http.Request, classService *services.ClassService) {
	var signUpDto dto.SignUpDto
	err := json.NewDecoder(r.Body).Decode(&signUpDto)
	if err != nil {
		utils.Response(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = classService.SignUp(signUpDto)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, "Signed up successfully", http.StatusOK)
}

var SignOut = func(w http.ResponseWriter, r *http.Request, classService *services.ClassService) {
	var signOutDto dto.SignUpDto
	err := json.NewDecoder(r.Body).Decode(&signOutDto)
	if err != nil {
		utils.Response(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = classService.SignOut(signOutDto)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.Response(w, "Signed out successfully", http.StatusOK)
}
