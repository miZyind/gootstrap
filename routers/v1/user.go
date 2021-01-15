package v1

import "github.com/gin-gonic/gin"

func addUserRoutes(v1 *gin.RouterGroup) {
	group := v1.Group("users")

	// @security JWT
	// @summary Get users
	// @200 Success
	// @400 There is no user
	// @500 Unknown error
	group.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "users"})
	})
}
