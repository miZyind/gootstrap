package v1

import "github.com/gin-gonic/gin"

// Ping ...
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
