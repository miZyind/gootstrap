package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mizyind/gootstrap/configs"
	v1 "github.com/mizyind/gootstrap/routers/v1"
	"github.com/mizyind/gootstrap/utils/logger"
	"github.com/mizyind/gootstrap/utils/saunter"
)

func initGinEngine() *gin.Engine {
	gin.SetMode(configs.Server.Mode)

	engine := gin.New()

	v1.InitRouters(engine.Group("api"))

	engine.Use(gin.Recovery())
	engine.GET("/api/v1", saunter.Handler(engine.Routes()))
	engine.StaticFS("/swagger-static", saunter.Static())

	return engine
}

func main() {
	server := &http.Server{
		Addr:         configs.Server.Addr,
		ReadTimeout:  configs.Server.ReadTimeout,
		WriteTimeout: configs.Server.WriteTimeout,
		Handler:      initGinEngine(),
	}

	// TODO: Configurable Protocol
	logger.Bootstrapper("Application successfully started on", "http://"+server.Addr)

	server.ListenAndServe()
}
