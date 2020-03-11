package utils

import (
	"errors"
	"ginp/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RandName 生成随机名字
func RandName(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	newName := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		newName[i] = letters[rand.Intn(len(letters))]
	}
	return string(newName)
}

// CheckTelephone 检查手机号是否合法
func CheckTelephone(c *gin.Context, telephone string) error {
	if len(telephone) != 11 {
		models.Response(c, http.StatusUnprocessableEntity, nil, "手机号必须为11位")
		return errors.New("手机号不合法")
	}
	return nil
}

// CheckPassword 检查密码是否合法
func CheckPassword(c *gin.Context, password string) error {
	if len(password) < 6 {
		models.Response(c, http.StatusUnprocessableEntity, nil, "密码不能少于6位")
		return errors.New("密码不合法")
	}
	return nil
}
