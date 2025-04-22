package router

import (
	"github.com/amreshpro/go-backend/env-in-go/internal/handler"
	"net/http"
)

func AppRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.TestApi)
	mux.HandleFunc("/health", handler.HealthCheck)
	return mux
}
