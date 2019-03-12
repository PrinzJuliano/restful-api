package routes

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"Ping": "Pong",
		"Status": "WOW",
	})
}
