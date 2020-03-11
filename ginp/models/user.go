package models

import "ginp/tables"

// User 用户类型
type User struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

// ToUser tables.User 转换为 models.User
func ToUser(user tables.User) User {
	return User{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
