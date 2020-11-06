package routers

import (
	"strings"

	"github.com/gin-gonic/gin"

	v1 "github.com/mizyind/gootstrap/routers/v1"
	"github.com/mizyind/gootstrap/utils/logger"
)

type router struct {
	path     string
	instance interface {
		BindRoutes(g *gin.RouterGroup) string
	}
}

var (
	engine    *gin.Engine
	routersV1 = []router{
		{path: "todos", instance: &v1.Todo{}},
		{path: "users", instance: &v1.User{}},
	}
)

func initRouters(g *gin.RouterGroup, routers []router) {
	for _, r := range routers {
		path := r.instance.BindRoutes(g.Group(r.path))

		logger.Router(r.instance, path)

		for _, route := range engine.Routes() {
			if strings.Contains(route.Path, path) {
				logger.Route(route.Method, route.Path)
			}
		}
	}
}

// Init returns a Gin engine with pre-defined routers
func Init(mode string) *gin.Engine {
	gin.SetMode(mode)

	engine = gin.New()

	engine.Use(gin.Recovery())

	api := engine.Group("api")

	initRouters(api.Group("v1"), routersV1)

	return engine
}
