package controls

import (
	"ginp/dbs"
	"ginp/models"
	"ginp/tables"
	"ginp/utils"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register 注册控制器
func Register(c *gin.Context) {
	// name := c.PostForm("name")
	// telephone := c.PostForm("telephone")
	// password := c.PostForm("password")

	var requestUser tables.User
	c.Bind(&requestUser)

	err := utils.CheckTelephone(c, requestUser.Telephone)
	if err != nil {
		return
	}

	err = utils.CheckPassword(c, requestUser.Password)
	if err != nil {
		return
	}

	if len(requestUser.Name) == 0 {
		requestUser.Name = utils.RandName(10)
	}

	if tables.IsTelephoneExist(dbs.GinDB, requestUser.Telephone) {
		models.ResponseClientError(c, nil, "用户已存在")
		return
	}

	// 创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		models.ResponseServerError(c, nil, "服务端异常，请稍后重试")
		return
	}

	requestUser.Password = string(hashedPassword)

	dbs.GinDB.Create(&requestUser)

	models.ResponseSuccess(c, nil, "注册成功")
}

// Login 登录控制器
func Login(c *gin.Context) {
	// 获取参数
	// telephone, _ := c.GetQuery("name") //c.PostForm("telephone")
	// telephone, _ := c.GetQuery("telephone") //c.PostForm("telephone")
	// password, _ := c.GetQuery("password")

	var loginUser tables.User
	c.BindQuery(&loginUser)
	// c.Bind(&loginUser)

	// 数据验证
	err := utils.CheckTelephone(c, loginUser.Telephone)
	if err != nil {
		return
	}

	err = utils.CheckPassword(c, loginUser.Password)
	if err != nil {
		return
	}

	// 判断手机号是否存在
	var user tables.User
	dbs.GinDB.Model(&user).Where("telephone=?", loginUser.Telephone).First(&user)
	if user.ID == 0 {
		models.ResponseClientError(c, nil, "手机号或密码不正确")
		return
	}

	// 判断密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		models.ResponseClientError(c, nil, "手机号或密码不正确")
		return
	}

	// 发放token
	token, err := utils.ReleaseToken(user)
	if err != nil {
		log.Println(err)
		models.ResponseServerError(c, nil, "服务异常，稍后再试")
		return
	}

	// 返回结果
	models.ResponseSuccess(c, gin.H{"token": token}, "登录成功")
}

// Info 用户信息控制器
func Info(c *gin.Context) {
	user, isExist := c.Get("user")
	if !isExist {
		log.Println("用户不存在")
		models.ResponseClientError(c, nil, "用户不存在")
		return
	}
	userT, ok := user.(tables.User)
	if !ok {
		log.Println("用户信息格式有误")
		models.ResponseServerError(c, nil, "服务异常，稍后重试")
		return
	}
	userM := models.ToUser(userT)
	models.ResponseSuccess(c, gin.H{"user": userM}, "")
}
