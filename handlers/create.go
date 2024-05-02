package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gui-sc/api-go/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Println("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": fmt.Sprintf("Error inserting todo: %v", err),
		}
	} else {
		resp = map[string]any{
			"error":   false,
			"message": fmt.Sprintf("Todo inserted successfully: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
