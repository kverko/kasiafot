package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func setRouters() {
	http.HandleFunc("/", index)
	http.HandleFunc("/admin/", admin)
	http.HandleFunc("/admin/upload/", upload)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func admin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[:7] == "/admin/" {
		id := r.URL.Path[7:]
		fmt.Println(id)
	}
	if r.Method == "POST" {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("empty")
			if senttok == token {
				fmt.Println("the same token")
			} else {
				senttok = token
			}
		} else {
			http.Redirect(w, r, "/", 200)
		}
		fmt.Println("user: ", r.Form["user"])
		fmt.Println("pass: ", r.Form["pass"])
	} else {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("templates/admin.html")
		t.Execute(w, token)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("templates/upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		dir_path := "./test/"
		err = os.MkdirAll(dir_path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.OpenFile(dir_path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

}
