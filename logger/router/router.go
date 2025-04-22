package router

import (
	"net/http"

	"github.com/amreshpro/go-backend/logger/internal/handler"
)


func AppRouter() http.Handler{

	mux := http.NewServeMux()

	mux.HandleFunc("/",handler.Hello)
	mux.HandleFunc("/about",handler.About)
	mux.HandleFunc("/contact",handler.Contact)
	mux.HandleFunc("/health",handler.HealthCheck)


	return mux
}