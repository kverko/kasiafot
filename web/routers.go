package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	templatePath = "./templates"
)

func setRouters() {
	http.HandleFunc("/", home)
	http.HandleFunc("/admin", MustLogin(admin))
	http.HandleFunc("/admin/login", login)
	http.HandleFunc("/admin/logout", MustLogin(logout))
	http.HandleFunc("/admin/list-tags", MustLogin(listTags))
}

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(filepath.Join(templatePath, "home.html"))
	t.Execute(w, nil)
}

func admin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(filepath.Join(templatePath, "admin.html"))
	t.Execute(w, nil)
}

func listTags(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(filepath.Join(templatePath, "list-tags.html"))
	t.Execute(w, nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	sid, err := sessionsManager.SessionID(r)
	if err != nil {
		fmt.Println("logout: couldn't retrieve current session id")
	}
	sessionsManager.RemoveSession(sid)
	sessionsManager.DelSessionCookie(w, r)
	http.Redirect(w, r, "/admin/login", 302)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles(filepath.Join(templatePath, "login.html"))
		t.Execute(w, nil)
	}

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			panic("POST: cannot parse form")
		}
		if r.Form.Get("username") != "admin" || r.Form.Get("password") != "pass" {
			http.Redirect(w, r, "/admin/login", 302)
			return
		}
		err = sessionsManager.SessionStart(w, r)
		if err != nil {
			log.Fatal("login router: couldn't start session")
		}
		http.Redirect(w, r, "/admin", 302)
	}
}
