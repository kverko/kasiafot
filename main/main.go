package main

import (
	"fmt"
	"net/http"

	"github.com/kverko/kasiafot/sessions"
)

var globalSesMan *sessions.Manager

func main() {
	globalSesMan = sessions.NewManager("sessid", 0)
	fmt.Println("starting server on port 8888")
	http.ListenAndServe(":8888", nil)
}
