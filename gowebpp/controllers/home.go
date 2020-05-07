package controllers

import (
	"gowebpp/confs"
	"gowebpp/models"
	"gowebpp/sessions"
	"gowebpp/utils"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// HandleHome 主页控制器
func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ckUser, err := r.Cookie("token")
		if err != nil || sessions.Sessions[ckUser.Value] == nil {
			w.Header().Set("Location", "/login")
			w.WriteHeader(http.StatusFound)
			confs.Logger.Printf("请求没有cookie，转到登录页，错误：%v\n", err)
			return
		}
		t, err := template.ParseFiles("./views/home.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("服务异常，请稍后再试"))
			confs.Logger.Printf("主页请求错误：%v\n", err)
			return
		}
		// t.Execute(w, template.HTML("<script>alert('登录成功')</script>"))
		t.Execute(w, template.HTMLEscapeString("<script>alert('登录成功')</script>"))
		return
	}
	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(21 * 1024 * 1024)
		if err != nil {
			confs.Logger.Printf("上传文件出错，错误：%v\n", err)
			utils.ProcessClientError(w, r)
			return
		}
		f, header, err := r.FormFile("file")
		if err != nil {
			utils.ProcessServerError(w, err)
			return
		}
		defer f.Close()

		fileName := ""
		if strings.Contains(header.Filename, "\\") {
			_, fileName = filepath.Split(header.Filename)
		} else {
			_, fileName = path.Split(header.Filename)
		}

		file := &models.File{
			Name:     fileName,
			Size:     header.Size,
			Location: "./temp/" + fileName,
		}

		dest, err := os.Create(file.Location)
		if err != nil {
			utils.ProcessServerError(w, err)
			return
		}
		defer dest.Close()
		_, err = io.Copy(dest, f)
		if err != nil {
			utils.ProcessServerError(w, err)
			return
		}

		hash, err := utils.MD5File(dest)
		if err != nil {
			utils.ProcessServerError(w, err)
			return
		}
		file.Hash = hash
		id, err := file.Add()
		if err != nil || id < 1 {
			utils.ProcessServerError(w, err)
			return
		}
		w.Write([]byte("文件上传成功"))
		return
	}
}
