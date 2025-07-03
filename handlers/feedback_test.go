package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yemiwebby/feedback-service/models"
)

func TestFeedbackHandler(t *testing.T) {
	// Reset shared state
	feedbacks = []models.Feedback{}

	t.Run("GET empty feedback", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/feedback", nil)
		res := httptest.NewRecorder()

		FeedbackHandler(res, req)

		if res.Code != http.StatusOK {
			t.Fatalf("Expected 200, got %d", res.Code)
		}

		var got []models.Feedback
		if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if len(got) != 0 {
			t.Errorf("Expected 0 feedbacks, got %d", len(got))
		}
	})

	t.Run("POST valid feedback", func(t *testing.T) {
		payload := models.Feedback{Name: "User", Message: "Clean API!"}
		body, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/feedback", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		FeedbackHandler(res, req)

		if res.Code != http.StatusCreated {
			t.Fatalf("Expect 201 created, got %d", res.Code)
		}

		var got models.Feedback
		if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if got.Name != payload.Name || got.Message != payload.Message {
			t.Errorf("Mismatch in response: got %+v, expected %+v", got, payload)
		}
	})

	t.Run("POST invalid feedback (missing fields)", func(t *testing.T) {
		payload := models.Feedback{Name: "", Message: ""}
		body, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPost, "/feedback", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		FeedbackHandler(res, req)

		if res.Code != http.StatusBadRequest {
			t.Fatalf("Expected 400 Bad Request, got %d", res.Code)
		}
	})

	t.Run("GET after feedback added", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/feedback", nil)
		res := httptest.NewRecorder()

		FeedbackHandler(res, req)

		var got []models.Feedback
		if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if len(got) != 1 {
			t.Errorf("Expected 1 feedback, got %d", len(got))
		}
	})
}
