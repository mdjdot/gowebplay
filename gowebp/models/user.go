package models

import (
	"gowebp/dbs"
)

// User 用户类型
type User struct {
	ID       int64
	UserName string
	Password string
	Email    string
}

// AddUser 添加用户
func (u *User) AddUser(userName, password, email string) (int64, error) {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	stmt, err := dbs.WebDB.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(userName, password, email)
	if err != nil {
		return 0, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

// ListUser 列出所有用户
func (u *User) ListUser() (*[]User, error) {
	var users []User
	var id int64
	var userName string
	var email string
	sqlStr := "select id,username,email from users"
	stmt, err := dbs.WebDB.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for result.Next() {
		result.Scan(&id, &userName, &email)
		users = append(users, User{ID: id, UserName: userName, Email: email})
	}
	return &users, nil
}

// QueryUser 查询用户
func (u *User) QueryUser(userName string) (*User, error) {
	userp := &User{}
	sqlStr := "select id,username,email from users where username=?"
	stmt, err := dbs.WebDB.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Query(userName)
	if err != nil {
		return nil, err
	}
	for result.Next() {
		result.Scan(&userp.ID, &userp.UserName, &userp.Email)
	}

	return userp, nil
}

// ModifyUser 修改用户信息
func (u *User) ModifyUser(id int64) error {
	return nil
}

// DeleteUser 删除用户
func (u *User) DeleteUser(id int64) error {
	return nil
}
