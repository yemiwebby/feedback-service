package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yemiwebby/feedback-service/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server Ok")
	})

	mux.HandleFunc("/feedback", handlers.FeedbackHandler)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
