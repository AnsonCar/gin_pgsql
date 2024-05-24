package user

import "github.com/gin-gonic/gin"

func getting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user",
	})
}

func posting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func putting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func deleting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
