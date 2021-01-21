package v1

import "github.com/gin-gonic/gin"

func addUserRoutes(root *gin.RouterGroup) {
	group := root.Group("users")

	group.GET("/", getUsers)
}

// @security JWT
// @summary Get users
// @200 Success
// @400 There is no user
// @500 Unknown error
func getUsers(c *gin.Context) {
	c.JSON(200, gin.H{"message": "users"})
}
