package v1

import "github.com/gin-gonic/gin"

// Todo represents a virtual router struct
type Todo struct{}

// BindRoutes registers routes to a router group and returns base path
func (*Todo) BindRoutes(group *gin.RouterGroup) string {
	group.Any("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "todos"})
	})

	return group.BasePath()
}
