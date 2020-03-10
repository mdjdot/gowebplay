package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 定义用户类型
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);notn ull"`
	Telephone string `gorm:"type：varchar(11);not null unique"`
	Password  string `gorm:"size:250;not null"`
}

func main() {
	db := initDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
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
			name = randName(10)
		}

		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": http.StatusUnprocessableEntity,
				"msg":  "用户名已存在",
			})
			return
		}

		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		c.JSON(http.StatusOK, gin.H{
			"code":      http.StatusOK,
			"msg":       "注册成功",
			"name":      name,
			"telephone": telephone,
			"password":  password,
		})

	})
	r.Run(":8080")
}

func randName(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	newName := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		newName[i] = letters[rand.Intn(len(letters))]
	}
	return string(newName)
}

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "dm:dmtest@tcp(127.0.0.1:3306)/gindb?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}

	// db.CreateTable(&User{})
	db.AutoMigrate(&User{})
	return db
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
