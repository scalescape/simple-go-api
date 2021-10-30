package main

import (
	"log"
	"net/http"
	"os"

	"github.com/devdinu/simple-api/config"
	"github.com/devdinu/simple-api/logger"
	"github.com/gorilla/handlers"
)

func main() {
	appCfg := config.MustLoad()
	router, err := server(appCfg)
	if err != nil {
		logger.Fatalf("[Main] error creating server: %v", err)
	}

	addr := config.AppAddress()
	logger.Infof("[Main] listening on address %s", addr)
	lr := handlers.LoggingHandler(os.Stdout, router)
	err = http.ListenAndServe(addr, lr)
	if err != nil {
		log.Fatalf("[Main] error listening for rerquests on port: %s err: %v", addr, err)
	}
}
