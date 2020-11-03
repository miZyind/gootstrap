package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mizyind/gootstrap/configs"
	"github.com/mizyind/gootstrap/routers"
	"github.com/mizyind/gootstrap/utils/logger"
)

var (
	serverConfig   *configs.ServerConfig
	databaseConfig *configs.DatabaseConfig
	handler        *gin.Engine
)

func init() {
	logger.Boot("Starting application...")

	configs.InitViper()
	logger.Boot("Viper initialized")

	serverConfig = configs.InitServerConfig()
	logger.Boot("Server config initialized", serverConfig)

	databaseConfig = configs.InitDatabaseConfig()
	logger.Boot("Database config initialized", databaseConfig)
}

func main() {
	server := &http.Server{
		Addr:         serverConfig.Addr,
		Handler:      routers.Init(serverConfig.Mode),
		ReadTimeout:  serverConfig.ReadTimeout,
		WriteTimeout: serverConfig.WriteTimeout,
	}

	logger.Boot(
		"Application successfully started on",
		fmt.Sprint("http://", server.Addr),
	)

	server.ListenAndServe()
}
