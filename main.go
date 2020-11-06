package main

import (
	"net/http"

	"github.com/mizyind/gootstrap/configs"
	"github.com/mizyind/gootstrap/routers"
	"github.com/mizyind/gootstrap/utils/logger"
)

func main() {
	server := &http.Server{
		Addr:         configs.Server.Addr,
		Handler:      routers.Init(configs.Server.Mode),
		ReadTimeout:  configs.Server.ReadTimeout,
		WriteTimeout: configs.Server.WriteTimeout,
	}

	// TODO: Configurable Protocol
	logger.Bootstrapper("Application successfully started on", "http://"+server.Addr)

	server.ListenAndServe()
}
