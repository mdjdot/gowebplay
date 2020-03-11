package controls

import (
	"ginp/dbs"
	"ginp/models"
	"ginp/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register 注册控制器
func Register(c *gin.Context) {
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	err := utils.CheckTelephone(c, telephone)
	if err != nil {
		return
	}

	err = utils.CheckPassword(c, password)
	if err != nil {
		return
	}

	if len(name) == 0 {
		name = utils.RandName(10)
	}

	if models.IsTelephoneExist(dbs.GinDB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "用户已存在",
		})
		return
	}

	// 创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务端异常，请稍后重试",
		})
		return
	}

	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
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

// Login 登录控制器
func Login(c *gin.Context) {
	// 获取参数
	// telephone, _ := c.GetQuery("name") //c.PostForm("telephone")
	telephone, _ := c.GetQuery("telephone") //c.PostForm("telephone")
	password, _ := c.GetQuery("password")

	// 数据验证
	err := utils.CheckTelephone(c, telephone)
	if err != nil {
		return
	}

	err = utils.CheckPassword(c, password)
	if err != nil {
		return
	}

	// 判断手机号是否存在
	var user models.User
	dbs.GinDB.Model(&user).Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "手机号或密码不正确",
		})
		return
	}

	// 判断密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "手机号或密码不正确",
		})
		return
	}

	// 发放token
	token, err := utils.ReleaseToken(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务异常，稍后再试",
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"token": gin.H{
			"token": token,
		},
		"msg": "登录成功",
	})
}

// Info 用户信息控制器
func Info(c *gin.Context) {
	user, isExist := c.Get("user")
	if !isExist {
		log.Println("用户不存在")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "用户不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
