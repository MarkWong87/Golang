package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("doumi", 666)
		c.Next()
		return
	}
}
