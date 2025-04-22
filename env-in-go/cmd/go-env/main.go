package main

import (
	"log"
	"net/http"
	"time"

	"github.com/amreshpro/go-backend/env-in-go/config"
	"github.com/amreshpro/go-backend/env-in-go/router"
)

func main() {
	// load all env
	config.LoadEnv()
	println("port: ", config.GetEnv("PORT", "8085"))
	println("AppName: ",config.GetEnv("APP_NAME","AppMockName"))
	//server
	server := &http.Server{
		Addr:         config.GetEnvPort("PORT", ":8085"), // yaha :portnumber yaani colon lagana must h 
		Handler:      router.AppRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// listen
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("failed to start sever: ", err)
	}

}
