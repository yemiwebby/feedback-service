package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/yemiwebby/feedback-service/models"
)

var (
	feedbacks []models.Feedback
	mu        sync.Mutex
)

func FeedbackHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var fb models.Feedback
		if err := json.NewDecoder(r.Body).Decode(&fb); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if fb.Name == "" || fb.Message == "" {
			http.Error(w, "Name and message required", http.StatusBadRequest)
			return
		}

		mu.Lock()
		feedbacks = append(feedbacks, fb)
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(fb)

	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(feedbacks)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
