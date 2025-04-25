package main

import (
	"net/http"
	"github.com/amreshpro/go-backend/3-expense-tracker/pkg/config"
	"github.com/amreshpro/go-backend/3-expense-tracker/pkg/logger"
	"github.com/amreshpro/go-backend/3-expense-tracker/internal/router"
)

func main() {
	config.LoadEnv()
	logger.InitLogger("development")
	defer logger.SyncLogger() // ✅ flush logs safely

	port := config.GetEnvPort("PORT", "3000")
	logger.Log.Infow("Starting server...", "port", port)
	jwtKeySec := config.GetEnv("JWT_SECRET_KEY", "")
	logger.Log.Infow("JWT SECRET KEY...", "JWTKEY: ", jwtKeySec)

	server := &http.Server{
		Addr:    port,
		Handler: router.AppRouter(),
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Log.Panicw("Server failed to start", "error", err)
	}
}
