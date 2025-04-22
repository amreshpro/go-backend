package router

import (
	"net/http"

	"github.com/amreshpro/go-backend/2-static-file-serve/internal/handler"
)

func AppRouter() http.Handler {
	mux := http.NewServeMux()
	// Serve static files from ./web directory
	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/health", handler.HealthCheck)
	return mux
}
