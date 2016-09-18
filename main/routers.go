package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func init() {
	setRouters()
}

func setRouters() {
	http.HandleFunc("/", home)
	http.HandleFunc("/admin", MustLogin(admin))
	http.HandleFunc("/admin/login", login)
	http.HandleFunc("/admin/logout", logout)
	http.HandleFunc("/admin/list-tags", MustLogin(list_tags))
}

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../templates/home.html")
	t.Execute(w, nil)
}

func admin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../templates/admin.html")
	t.Execute(w, nil)
}

func list_tags(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../templates/list-tags.html")
	t.Execute(w, nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	sid, err := globalSesMan.SessionID(r)
	if err != nil {
		fmt.Println("logout: couldn't retrieve current session id")
	}
	globalSesMan.RemoveSession(sid)
	globalSesMan.DelSessionCookie(w, r)
	http.Redirect(w, r, "/admin/login", 302)
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("../templates/login.html")
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
		err = globalSesMan.SessionStart(w, r)
		if err != nil {
			log.Fatal("login router: couldn't start session")
		}
		http.Redirect(w, r, "/admin", 302)
	}
}
