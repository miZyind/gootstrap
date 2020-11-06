package v1

import "github.com/gin-gonic/gin"

// User represents a virtual router struct
type User struct{}

// BindRoutes registers routes to a router group and returns base path
func (*User) BindRoutes(group *gin.RouterGroup) string {
	group.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "users"})
	})

	return group.BasePath()
}
