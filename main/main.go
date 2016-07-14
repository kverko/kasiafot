package main

import (
	"fmt"
	"net/http"

	"github.com/kverko/kasiafot/sessions"
)

var sesMan *sessions.Manager

func main() {
	sesMan = &sessions.Manager{
		CookieName: "sessid",
		Lifetime:   0,
		Sessions:   make(map[string]sessions.Session, 0),
	}
	setHttpHandlers()
	fmt.Println("starting server on port 8888")
	http.ListenAndServe(":8888", nil)
}
