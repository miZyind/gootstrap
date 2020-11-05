package v1

import "github.com/gin-gonic/gin"

type User struct{}

func (*User) BindRoutes(g *gin.RouterGroup) string {
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "users"})
	})

	return g.BasePath()
}
