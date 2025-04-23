package main

import (
	"net/http"
	"github.com/amreshpro/go-backend/3-expense-tracker/internal/router"
	"github.com/amreshpro/go-backend/3-expense-tracker/pkg/logger"
)

func main() {
	
logger.InitLogger("development")

server := &http.Server{
	Addr: ":8080",
	Handler: router.AppRouter(),
	}


	logger.Log.Infow("Server is started")
	if err:= server.ListenAndServe(); err!=nil{
		logger.Log.Panicw("Failed to start error: ",err)
	}


}