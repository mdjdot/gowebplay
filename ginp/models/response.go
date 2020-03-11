package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 封装响应
func Response(c *gin.Context, httpStatus int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{"code": httpStatus, "data": data, "msg": msg})
}

// ResponseSuccess 封装成功的响应
func ResponseSuccess(c *gin.Context, data gin.H, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": data, "msg": msg})
}

// ResponseClientError 封装客户端错误的响应
func ResponseClientError(c *gin.Context, data gin.H, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": data, "msg": msg})
}

// ResponseServerError 封装服务端错误的响应
func ResponseServerError(c *gin.Context, data gin.H, msg string) {
	c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "data": data, "msg": msg})
}
