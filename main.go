// main.go
package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi, this is index")
}

func admin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi, this is admin")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/admin", admin)

	http.ListenAndServe(":8000", nil)
}
