package main

import (
	"fmt"
	"net/http"

	"github.com/kverko/kasiafot/sessions"
	_ "github.com/kverko/kasiafot/web"
)

//SessionsManager to rule all sessions
var SessionsManager = sessions.NewManager("sessid", 0)

func main() {

	sessions.SessionsManager = SessionsManager

	fmt.Println("starting server on port 8888")
	http.ListenAndServe(":8888", nil)
}
