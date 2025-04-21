package router

import (
	"net/http"
	"github.com/amreshpro/go-backend/2-static-file-serve/internal/handler"
)


func AppRouter() http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		handler.HTMLFileSender(w,r,"./web/index.html")
	})
	mux.HandleFunc("/health",handler.HealthCheck)
	return mux
}