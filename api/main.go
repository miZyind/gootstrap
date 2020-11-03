package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mizyind/gootstrap/api/v1"
)

// Bind ...
func Bind(mode string) *gin.Engine {
	gin.SetMode(mode)

	handler := gin.New()

	handler.Use(gin.Recovery())

	api := handler.Group("/api")

	apiV1 := api.Group("/v1")
	{
		apiV1.GET("/", v1.Ping)
	}

	return handler
}
