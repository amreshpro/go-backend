package main

import (
	"log"
	"net/http"
	"time"
	"github.com/amreshpro/go-backend/2-static-file-serve/router"
)




func main() {

server:= &http.Server{
	Addr: ":8080",
	ReadTimeout: 5*time.Second,
	WriteTimeout: 10*time.Second,
	IdleTimeout: 120*time.Second,
	Handler: router.AppRouter(),
}

log.Print("Server started")
if err := server.ListenAndServe(); err !=nil{
	log.Fatal("server failed to start: ",err)
}

	
}