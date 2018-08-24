package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Person To make JSON file, the format should be CAPITAL LETTER
type Person struct {
	UserName string
	Password string
}

func exists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			return false, nil
		}
		// exist but got other error
		return true, err
	}
	return true, nil
}

// func (p *Person) deletePerson() {

// 	db := "./database/person/" + p.UserName + ".json"

// 	if err := os.Remove(db); err != nil {
// 		log.Fatalln(err)
// 	}
// }

func userHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/user/"):]
	username := r.FormValue("username")
	psw := r.FormValue("psw")
	p := Person{
		UserName: username,
		Password: psw,
	}
	page := &Page{Title: p.UserName, Body: []byte("not Logged in")}
	db := "./database/person/" + p.UserName + ".json"
	exs, err := exists(db)
	if err != nil {
		log.Fatalln(err)
	}
	switch title {
	case "signup":
		if !exs {

			var b []byte
			b, err = json.Marshal(p)
			if err != nil {
				log.Fatalln(err)
			}
			if err = ioutil.WriteFile(db, b, 0644); err != nil {
				log.Fatalln(err)
			}

		} else {
			http.Redirect(w, r, "/user/login", http.StatusFound)
		}

	case "login":
		if !exs {
			http.Redirect(w, r, "/user/signup", http.StatusFound)

		} else {
			var replys Person
			content, err := ioutil.ReadFile(db)
			if err != nil {
				log.Fatalln(err)
			}
			if err = json.Unmarshal(content, &replys); err != nil {
				log.Fatalln(err)
			}

			if p.UserName == replys.UserName && p.Password == replys.Password {
				page = &Page{Title: p.UserName, Body: []byte("is NOW Logged in")}
			} else {
				log.Printf("Not equal")
			}
		}

	}

	renderTemplateUser(w, "user", page)
}

func renderTemplateUser(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandlerUser(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fn(w, r)
	}
}
