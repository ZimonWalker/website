package main

import (
	"net/http"
)

//Hr func
func Hr(w http.ResponseWriter, r *http.Request) {
	page := &Page{Title: "p.UserName", Body: []byte("not Logged in")}
	render(w, "hr.html", page)
}

// //Staff2 func
// func Staff2(w http.ResponseWriter, r *http.Request) {
// 	page := &Page{Title: "p.UserName", Body: []byte("not Logged in")}
// 	render(w, "staff2.html", page)
// }

// //Staff3 func
// func Staff3(w http.ResponseWriter, r *http.Request) {
// 	page := &Page{Title: "p.UserName", Body: []byte("not Logged in")}
// 	render(w, "staff3.html", page)
// }
