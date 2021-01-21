package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mizyind/gootstrap/configs"
	v1 "github.com/mizyind/gootstrap/routers/v1"
	v2 "github.com/mizyind/gootstrap/routers/v2"
	"github.com/mizyind/gootstrap/utils/logger"
	"github.com/mizyind/saunter"
)

func main() {
	gin.SetMode(configs.Server.Mode)

	engine := gin.New()
	basePath := "/api"
	root := engine.Group(basePath)

	v1.InitRouters(root)
	v2.InitRouters(root)

	saunter.Initialize(basePath, engine.Routes())

	engine.Use(gin.Recovery())
	engine.GET("/api/v1", saunter.Handler())
	engine.GET("/api/v2", saunter.Handler())
	engine.StaticFS("/swagger-static", saunter.Static())

	server := &http.Server{
		Addr:         configs.Server.Addr,
		ReadTimeout:  configs.Server.ReadTimeout,
		WriteTimeout: configs.Server.WriteTimeout,
		Handler:      engine,
	}

	// TODO: Configurable Protocol
	logger.Bootstrapper("Application successfully started on", "http://"+server.Addr)

	server.ListenAndServe()
}
