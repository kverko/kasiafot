// main.go
package main

import (
	"log"
	"net/http"
)

var senttok string
var globalSessions *Manager

func init() {
	globalSessions := NewManager("memory", "sessid", 1800)
}

func main() {
	senttok = ""
	setRouters()
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
