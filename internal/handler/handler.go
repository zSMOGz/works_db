package handler

import (
	"encoding/json"
	"net/http"
)

// HelloHandler возвращает JSON с сообщением "Привет, мир!"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Привет, мир!"})
}

// GoodbyeHandler возвращает JSON с сообщением "Пока, мир!"
func GoodbyeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Пока, мир!"})
}
