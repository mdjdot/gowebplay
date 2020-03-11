package middlewares

import (
	"ginp/dbs"
	"ginp/models"
	"ginp/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("authorization")

		// 验证token 格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "权限不足",
			})
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		// 验证通过，获取claims中的userID
		userID := claims.UserID
		var user models.User
		dbs.GinDB.First(&user, userID)
		// dbs.GinDB.Model(&User{}).First(&user, userID)

		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		// 如果用户存在，将user 信息写入上下文
		c.Set("user", user)

		c.Next()
	}
}
