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

//User handler for /user renders the user.html
func User(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/user/"):]
	username := r.FormValue("username")
	psw := r.FormValue("psw")
	p := Person{
		UserName: username,
		Password: psw,
	}
	page := &Page{Title: p.UserName, Body: "not Logged in"}
	db := "./database/person/" + p.UserName + ".json"
	exs, err := exists(db)
	if err != nil {
		log.Fatalln(err)
	}
	switch title {
	case "signup":
		if !exs && p.UserName != "" {

			var b []byte
			b, err = json.Marshal(p)
			if err != nil {
				log.Fatalln(err)
			}
			if err = ioutil.WriteFile(db, b, 0644); err != nil {
				log.Fatalln(err)
			}

		} else {
			http.Redirect(w, r, "/user/", http.StatusFound)
		}

	case "login":
		if !exs {
			http.Redirect(w, r, "/user/", http.StatusFound)

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
				page = &Page{Title: p.UserName, Body: "is NOW Logged in"}
			} else {
				log.Printf("Not equal")
			}
		}

	}

	render(w, "user.html", page)
}
