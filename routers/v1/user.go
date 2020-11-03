package v1

import "github.com/gin-gonic/gin"

// UserRouter ...
type UserRouter struct{}

// Bind ...
func (*UserRouter) Bind(group *gin.RouterGroup) string {
	group = group.Group("users")

	group.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "users"})
	})

	return group.BasePath()
}
