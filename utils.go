package utils

import "github.com/gin-gonic/gin"

func RespondWithError(c *gin.Context, code int, msg string, err error) {
	if err != nil {
		c.JSON(code, gin.H{
			"error":   msg,
			"details": err.Error(),
		})
	} else {
		c.JSON(code, gin.H{
			"error": msg,
		})
	}
}
