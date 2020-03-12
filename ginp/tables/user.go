package tables

import "github.com/jinzhu/gorm"

// User 定义用户类型
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);notn ull" form:"name"`
	Telephone string `gorm:"type：varchar(11);not null unique" form:"telephone"`
	Password  string `gorm:"size:250;not null" form:"password"`
}

// IsTelephoneExist 判断用户手机号是否已存在
func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
