package v1

import "github.com/gin-gonic/gin"

type Todo struct{}

func (*Todo) BindRoutes(g *gin.RouterGroup) string {
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "todos"})
	})

	return g.BasePath()
}
