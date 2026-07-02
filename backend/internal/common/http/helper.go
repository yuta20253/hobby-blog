package httphelper

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func getUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return 0, false
	}

	id, ok := userID.(uint)

	if !ok {
		c.JSON(500, gin.H{"error": "invalid user id"})
		return 0, false
	}

	return uint(id), true
}

func getParamID(c *gin.Context, key string) (uint, bool) {
	idStr := c.Param(key)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return 0, false
	}

	return uint(id), true
}
