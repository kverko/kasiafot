// main.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func admin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("user: ", r.Form["user"])
		fmt.Println("pass: ", r.Form["pass"])
	} else {
		t, _ := template.ParseFiles("templates/admin.html")
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/admin", admin)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
