package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type JsonResponseMessage struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		jsonResponse := JsonResponseMessage{
			Message: "Hello from json type response",
			Status:  "OK",
		}

		json.NewEncoder(w).Encode(jsonResponse)

	})

	mux.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello text response"))
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Print("Server Started")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
