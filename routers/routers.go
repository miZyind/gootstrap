package routers

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"

	v1 "github.com/mizyind/gootstrap/routers/v1"
	"github.com/mizyind/gootstrap/utils/logger"
)

type Router interface {
	BindRoutes(g *gin.RouterGroup) string
}

type RouterInfo struct {
	path   string
	router Router
}

var (
	engine    *gin.Engine
	routersV1 = []RouterInfo{
		{path: "todos", router: &v1.Todo{}},
		{path: "users", router: &v1.User{}},
	}
)

func initRouters(group *gin.RouterGroup, routerInfo []RouterInfo) {
	for _, info := range routerInfo {
		name := reflect.TypeOf(info.router).Elem().Name()
		path := info.router.BindRoutes(group.Group(info.path))

		logger.InitRouter(name, path)

		for _, route := range engine.Routes() {
			if strings.Contains(route.Path, path) {
				logger.BindRoute(route.Method, route.Path)
			}
		}
	}
}

func Init(mode string) *gin.Engine {
	gin.SetMode(mode)

	engine = gin.New()

	engine.Use(gin.Recovery())

	api := engine.Group("api")

	initRouters(api.Group("v1"), routersV1)

	return engine
}
