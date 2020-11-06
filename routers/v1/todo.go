package v1

import "github.com/gin-gonic/gin"

type Todo struct{}

func (*Todo) BindRoutes(group *gin.RouterGroup) string {
	group.Any("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "todos"})
	})

	return group.BasePath()
}
