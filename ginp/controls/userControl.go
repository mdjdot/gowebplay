package controls

import (
	"ginp/dbs"
	"ginp/models"
	"ginp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 注册控制器
func Register(c *gin.Context) {
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "手机号必须为11位",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "密码不能少于6位",
		})
		return
	}

	if len(name) == 0 {
		name = utils.RandName(10)
	}

	if models.IsTelephoneExist(dbs.GinDB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "用户名已存在",
		})
		return
	}

	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	dbs.GinDB.Create(&newUser)

	c.JSON(http.StatusOK, gin.H{
		"code":      http.StatusOK,
		"msg":       "注册成功",
		"name":      name,
		"telephone": telephone,
		"password":  password,
	})
}
