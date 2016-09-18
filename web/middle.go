package web

import (
	"net/http"

	"github.com/kverko/kasiafot/sessions"
)

//MustLogin middleware to check if user is authenticated before using a handler
func MustLogin(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !sessions.SessionsManager.IsLoggedIn(r) {
			http.Redirect(w, r, "/admin/login", 302)
			return
		}
		handler(w, r)
	}
}
