package v2

import "github.com/gin-gonic/gin"

func addTodoRoutes(root *gin.RouterGroup) {
	group := root.Group("todos")

	group.GET("/", getTodos)
}

// @security JWT
// @summary Get todos
// @200 Success
// @400 There is no todo
// @500 Unknown error
func getTodos(c *gin.Context) {
	c.JSON(200, gin.H{"message": "todos"})
}
