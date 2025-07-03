package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yemiwebby/feedback-service/handlers"
)

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server Ok")
	})

	http.HandleFunc("/feedback", handlers.FeedbackHandler)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
