package routers

import (
	"reflect"

	"github.com/gin-gonic/gin"

	v1 "github.com/mizyind/gootstrap/routers/v1"
	"github.com/mizyind/gootstrap/utils/logger"
)

// Router ...
type Router interface {
	Bind(r *gin.RouterGroup) string
}

var (
	routersV1 = []Router{
		&v1.TodoRouter{},
		&v1.UserRouter{},
	}
)

func bindRoutersToGroup(group *gin.RouterGroup, routers []Router) {
	for _, router := range routers {
		path := router.Bind(group)
		logger.Router(
			reflect.Indirect(reflect.ValueOf(router)).Type().Name(),
			path,
		)
	}
}

// Init ...
func Init(mode string) *gin.Engine {
	gin.SetMode(mode)

	routers := gin.New()

	routers.Use(gin.Recovery())

	bindRoutersToGroup(routers.Group("api/v1"), routersV1)

	return routers
}
