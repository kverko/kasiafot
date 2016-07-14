package main

import "net/http"

func MustLogin(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !sesMan.IsLoggedIn(r) {
			http.Redirect(w, r, "/admin/login", 302)
			return
		}
		handler(w, r)
	}
}
