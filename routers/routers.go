package routers

import (
	"reflect"

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
	}
}

func Init(mode string) *gin.Engine {
	gin.SetMode(mode)

	routers := gin.New()

	routers.Use(gin.Recovery())

	initRouters(routers.Group("api/v1"), routersV1)

	return routers
}
