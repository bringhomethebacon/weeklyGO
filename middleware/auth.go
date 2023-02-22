package middleware

import (
	"net/http"

	"weekly-report/model"

	"github.com/gin-gonic/gin"
)

func AuthStudentCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userRole, err := model.AnalyseToken(auth)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized Authorization",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if userRole == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "身份无效",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user_claims", userRole)
		c.Next()
	}
}

func AuthTeacherCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userRole, err := model.AnalyseToken(auth)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized Authorization",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if userRole == nil && userRole.Role == "teacher" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "身份无效",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user_claims", userRole)
		c.Next()
	}
}
