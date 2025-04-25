package handler

import (
	"net/http"

	"github.com/amreshpro/go-backend/logger/pkg/logger"
)

func About(w http.ResponseWriter, r *http.Request) {
	logger.Log.Infow("GET request received", "path", r.URL.Path)
	w.Write([]byte("hello about"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	logger.Log.Infow("GET request received", "path", r.URL.Path)
	w.Write([]byte("hello hello"))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	logger.Log.Infow("GET request received", "path", r.URL.Path)
	w.Write([]byte("health is ok"))
}

func Contact(w http.ResponseWriter, r *http.Request) {
	logger.Log.Infow("GET request received", "path", r.URL.Path)
	w.Write([]byte("hello contact"))
}


