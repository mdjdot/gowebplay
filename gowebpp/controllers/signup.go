package controllers

import (
	"gowebpp/confs"
	"gowebpp/models"
	"gowebpp/sessions"
	"gowebpp/utils"
	"html/template"
	"net/http"
	"time"
)

// HandleSignup 注册控制器
func HandleSignup(w http.ResponseWriter, r *http.Request) {
	doc := &models.Request{URL: r.URL, Time: time.Now().Unix()}
	err := doc.Insert()
	if err != nil {
		confs.Logger.Println("记录请求到mongodb错误：" + err.Error())
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		user := &models.User{
			Name:     template.HTMLEscapeString(r.PostForm.Get("name")),
			Password: template.HTMLEscapeString(r.PostForm.Get("pwd")),
		}

		if user.Name == "" || user.Password == "" {
			utils.ProcessClientError(w, r)
			return
		}

		user.Password = utils.MD5(user.Password)
		if user.IsExist() {
			w.Write([]byte("用户名已存在，请重新注册"))
			return
		}
		id, err := user.Add()
		if err != nil || id < 1 {
			utils.ProcessServerError(w, err)
			return
		}
		token := utils.MD5(user.Name)
		// sessions.Sessions[token] = time.NewTimer(20 * time.Second)
		err = sessions.Add(token, 20)
		if err != nil {
			utils.ProcessServerError(w, err)
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "token", Value: token})
		w.Header().Add("Location", "/")
		w.WriteHeader(http.StatusFound)
	}
}
