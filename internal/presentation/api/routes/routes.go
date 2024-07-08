package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"schedule/internal/domain/services"
	"schedule/internal/presentation/api/controllers"
)

func SetupRoutes(router *mux.Router, classService *services.ClassService) {
	router.HandleFunc("/classes", func(w http.ResponseWriter, r *http.Request) {
		controllers.GeClassById(w, r, classService)
	}).Methods("GET")

	router.HandleFunc("/persons/classes", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllClassesByPerson(w, r, classService)
	}).Methods("GET")

	router.HandleFunc("/classes", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateClass(w, r, classService)
	}).Methods("POST")

	router.HandleFunc("/classes", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateClass(w, r, classService)
	}).Methods("PUT")

	router.HandleFunc("/classes", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteClass(w, r, classService)
	}).Methods("DELETE")

}
