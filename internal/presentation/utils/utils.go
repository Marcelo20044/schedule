package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				w.WriteHeader(500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Response(w http.ResponseWriter, message interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(message)
}
