package main

import (
	"net/http"
	"github.com/amreshpro/go-backend/logger/pkg/logger"
	"github.com/amreshpro/go-backend/logger/router"
)

func main() {
	
// load env if have

// load logger
logger.InitLogger("development")
defer logger.SyncLogger()


server := &http.Server{
	Addr: ":8080",
	Handler: router.AppRouter(),

}

logger.Log.Info("Server has started")
if err := server.ListenAndServe(); err !=nil{
	logger.Log.Error("failed to start server: ",err)
}

}