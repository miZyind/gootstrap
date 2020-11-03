package v1

import "github.com/gin-gonic/gin"

// TodoRouter ...
type TodoRouter struct{}

// Bind ...
func (*TodoRouter) Bind(group *gin.RouterGroup) string {
	group = group.Group("todos")

	group.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "todos"})
	})

	return group.BasePath()
}
