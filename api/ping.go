package api

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	// c.da
	c.JSON(200, gin.H{"message": "pong"})
}