package web

import (
	"fmt"
	"net/http"

	"github.com/kverko/kasiafot/sessions"
)

var sessionsManager *sessions.Manager

func init() {
	sessionsManager = sessions.NewManager("sessid", 0)
	setRouters()
	fmt.Println("starting server on port 8888")
	http.ListenAndServe(":8888", nil)
}
