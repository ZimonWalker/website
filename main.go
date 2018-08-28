// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/segmentio/ksuid"
)

//Page is blablaba
type Page struct {
	Title string
	Body  string
}

//LoginCre type
type LoginCre struct {
	Username     string `json:"Username"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	LoopUsername bool   `json:"LoopUsername"`
	LoginFlag    bool   `json:"LoginFlag"`
}

var gp = Page{
	Title: "",
	Body:  "",
}

var templates = template.Must(template.ParseFiles(
	"templates/home.html",
	"templates/register.html",
	"templates/staff.html",
	"templates/staff2.html",
	"templates/staff3.html",
	"templates/staff4.html",
	"templates/hr.html",
	"templates/hr2.html",
	"templates/hr3.html",
	"templates/hr4.html",
	"templates/hr5.html",
))

func main() {

	//myJSONFunc()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", Home)
	http.HandleFunc("/register/", Register)
	http.HandleFunc("/staff/", Staff)
	http.HandleFunc("/staff2/", Staff2)
	http.HandleFunc("/staff3/", Staff3)
	http.HandleFunc("/staff4/", Staff4)
	http.HandleFunc("/hr/", Hr)
	http.HandleFunc("/hr2/", Hr2)
	http.HandleFunc("/hr3/", Hr3)
	http.HandleFunc("/hr4/", Hr4)
	http.HandleFunc("/hr5/", Hr5)
	http.HandleFunc("/logout/", Logout)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func render(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl, p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

func genXid() string {
	id := ksuid.New()
	// fmt.Printf("github.com/segmentio/ksuid:  %s\n", id.String())
	return id.String()
}

//Home handler for / renders the home.html
func Home(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[len("/"):] //s for submit type
	p := &Page{
		Title: "Default",
	}

	if s == "login" {
		l := &LoginCre{
			Username:     r.FormValue("username"),
			Password:     r.FormValue("password"),
			Role:         "",
			LoopUsername: false,
			LoginFlag:    false,
		}
		myJSONFunc(l)
		if l.LoginFlag == true {
			fmt.Println("Success Login")
			if l.Role == "staff" {
				gp = Page{Title: "loggedStaff", Body: l.Username}
				http.Redirect(w, r, "/staff/", http.StatusFound)
				return
			}
			// resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})
			gp = Page{Title: "loggedHr", Body: l.Username}
			http.Redirect(w, r, "/hr/", http.StatusFound)
			return

		}
		fmt.Println("Failed Login")
	} else if s == "signup" {
		p := &StaffPage{
			Username:     r.FormValue("Username"),
			Password:     r.FormValue("Password"),
			Email:        r.FormValue("Email"),
			FullName:     r.FormValue("FullName"),
			Gender:       r.FormValue("Gender"),
			IC:           r.FormValue("IC"),
			Phone:        r.FormValue("Phone"),
			Role:         "staff",
			LeaveBalance: 20,
		}

		db := "./database/register/" + r.FormValue("Username") + ".json"

		b, err := json.Marshal(p)
		if err != nil {
			log.Fatalln(err)
		}
		if err = ioutil.WriteFile(db, b, 0644); err != nil {
			log.Fatalln(err)
		}
	}

	render(w, "home.html", p)
}

//myJSONFunc for parsing json file
func myJSONFunc(l *LoginCre) {

	// Read from file
	//declare db path

	db := "./database/login/" + l.Username + ".json"

	content, err := ioutil.ReadFile(db)
	if err != nil {
		// log.Fatalln(err)
		l.LoginFlag = false
		return
	}

	cre := &LoginCre{}

	// Parsing/Unmarshalling JSON encoding/json
	if err = json.Unmarshal(content, &cre); err != nil {
		log.Fatalln(err)
		// panic(err)
	}

	if cre.Password == l.Password {
		l.Role = cre.Role
		l.LoginFlag = true
	}

	// parseMap(m, l)
}

func parseMap(aMap map[string]interface{}, l *LoginCre) {

	for key, val := range aMap {
		if key == l.Username || l.LoopUsername == true {
			l.LoopUsername = true
			switch concreteVal := val.(type) {
			case map[string]interface{}:
				// fmt.Println("map case 1", key)
				parseMap(val.(map[string]interface{}), l)
			case []interface{}:
				// fmt.Println("map case 2", key)
				parseArray(val.([]interface{}), l)
			default:
				// fmt.Println(key, ":", concreteVal)
				// fmt.Println("map default", key, " : ", concreteVal)
				if l.LoopUsername == true && key == "password" && concreteVal != l.Password {
					l.LoopUsername = false
					return
				}
				if l.LoopUsername == true && key == "password" && concreteVal == l.Password {
					l.LoginFlag = true
				}
				if l.LoginFlag == true && key == "role" {
					l.Role = val.(string)
				}
			}
			if l.LoginFlag == true && l.Role != "" {
				return
			}
		}
	}
}

func parseArray(anArray []interface{}, l *LoginCre) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			// fmt.Println("array case 1 Index:", i)
			parseMap(val.(map[string]interface{}), l)
		case []interface{}:
			// fmt.Println("array case 2 Index:", i)
			parseArray(val.([]interface{}), l)
		default:
			fmt.Println("array default Index", i, ":", concreteVal)
			// if concreteVal == p {
			// 	fmt.Println("Pass")
			// }
		}
	}
}

// Logout func
func Logout(w http.ResponseWriter, r *http.Request) {

	gp = Page{Title: "", Body: ""}

	http.Redirect(w, r, "/", http.StatusFound)
	return
}

// Register func
func Register(w http.ResponseWriter, r *http.Request) {

	p := &Page{
		Title: "Default",
	}
	render(w, "register.html", p)
}
