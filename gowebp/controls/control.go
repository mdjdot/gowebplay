package controls

import (
	"encoding/json"
	"gowebp/logger"
	"gowebp/models"
	"html/template"
	"net/http"
)

// InitControls 初始化控制器
func InitControls() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	{
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			t := template.Must(template.ParseFiles("views/index.html"))
			t.Execute(w, nil)
		})

		http.HandleFunc("/ListUser", func(w http.ResponseWriter, r *http.Request) {
			usersp, err := (&models.User{}).ListUser()
			if err != nil {
				logger.Logger.Print(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("服务异常，请稍后重试"))
				return
			}
			usersb, err := json.Marshal(usersp)
			if err != nil {
				logger.Logger.Print(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("服务异常，请稍后重试"))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(usersb)
		})

		http.HandleFunc("/AddUser", func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			user := models.User{UserName: r.FormValue("name"), Password: r.FormValue("password"), Email: r.FormValue("email")}
			lastID, err := user.AddUser(user.UserName, user.Password, user.Email)
			if err != nil {
				logger.Logger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("StatusInternalServerError"))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("添加成功"))
			user.ID = lastID
			logger.Logger.Printf("Add user %v success", user)
		})

		http.HandleFunc("/QueryUser", func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			userName := r.FormValue("name")
			userp, err := (&models.User{}).QueryUser(userName)
			if err != nil {
				logger.Logger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("StatusInternalServerError"))
				return
			}
			w.WriteHeader(http.StatusOK)
			jsonUser, err := json.Marshal(userp)
			if err != nil {
				logger.Logger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("StatusInternalServerError"))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonUser)
		})
	}
}
