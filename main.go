// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"log"
	"net/http"
)

//Page is blablaba
type Page struct {
	Title string
	Body  []byte
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", Home)
	http.HandleFunc("/user/", User)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func render(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl, p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

//Home handler for / renders the home.html
func Home(w http.ResponseWriter, r *http.Request) {

	p := &Page{
		Title: "Default",
	}

	render(w, "home.html", p)
}

var templates = template.Must(template.ParseFiles("templates/home.html", "templates/user.html"))

func makeHandlerUser(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fn(w, r)

	}

}
