package v1

import "github.com/gin-gonic/gin"

// @Info
// @title My API
// @version 1.0.0
// @description This is a sample server celler server.

// @Security
// @name JWT
// @type http
// @description JWT Auth
// @scheme bearer

// InitRouters will bind all defined routers
func InitRouters(api *gin.RouterGroup) {
	root := api.Group("v1")

	addTodoRoutes(root)
	addUserRoutes(root)
}
