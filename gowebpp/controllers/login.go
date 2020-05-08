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

// HandleLogin 登录控制器
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	doc := &models.Request{URL: r.URL, Time: time.Now().Unix()}
	err := doc.Insert()
	if err != nil {
		confs.Logger.Println("记录请求到mongodb错误：" + err.Error())
	}
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("./views/login.html")
		if err != nil {
			utils.ProcessServerError(w, err)
			return
		}
		t.Execute(w, nil)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		user := &models.User{
			Name:     template.HTMLEscapeString(r.PostForm.Get("name")),
			Password: template.HTMLEscapeString(r.PostForm.Get("pwd")),
		}
		user.Password = utils.MD5(user.Password)
		user.GetByNameAndPwd()
		if user.ID <= 0 {
			utils.ProcessClientError(w, r)
			return
		}
		token := utils.MD5(user.Name)
		// sessions.Sessions[token] = time.NewTimer(20 * time.Second)
		err := sessions.Add(token, 20)
		if err != nil {
			utils.ProcessClientError(w, r)
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "token", Value: token})
		w.Header().Add("Location", "/")
		w.WriteHeader(http.StatusFound)
	}
}
