package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// HrPage struct
type HrPage struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
	FullName string `json:"FullName"`
	Gender   string `json:"Gender"`
	IC       string `json:"IC"`
	Phone    string `json:"Phone"`
	Role     string `json:"Role"`
}

// HRLeave struct
type HRLeave struct {
	StaffLeave []StaffLeave `json:"staffLeave"`
}

var hp = &HrPage{}
var hl = &HRLeave{}

//Hr func
func Hr(w http.ResponseWriter, r *http.Request) {

	if gp.Title != "loggedHr" {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	u := gp.Body

	db := "./database/hr/" + u + ".json"
	var content []byte
	content, err := ioutil.ReadFile(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Parsing/Unmarshalling JSON encoding/json
	if err = json.Unmarshal(content, &hp); err != nil {
		log.Fatalln(err)
		// panic(err)
	}

	renderHR(w, "hr.html", hp)
}

func renderHR(w http.ResponseWriter, tmpl string, p *HrPage) {
	err := templates.ExecuteTemplate(w, tmpl, p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

//Hr2 func
func Hr2(w http.ResponseWriter, r *http.Request) {
	if gp.Title != "loggedHr" {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	files, err := ioutil.ReadDir("./database/leave")
	if err != nil {
		log.Fatalln(err)
	}

	i := 0

	for _, f := range files {
		// fmt.Println(f.Name(), i)
		var miniSL StaffLeave
		db := "./database/leave/" + f.Name()
		var content []byte
		content, err := ioutil.ReadFile(db)
		if err != nil {
			log.Fatalln(err)
		}

		// Parsing/Unmarshalling JSON encoding/json
		if err = json.Unmarshal(content, &miniSL); err != nil {
			log.Fatalln(err)
			// panic(err)
		}

		hl.StaffLeave = append(hl.StaffLeave, miniSL)

		i++
	}

	renderHR2(w, "hr2.html", hl)
}

func renderHR2(w http.ResponseWriter, tmpl string, p *HRLeave) {
	err := templates.ExecuteTemplate(w, tmpl, p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

// //Staff3 func
// func Staff3(w http.ResponseWriter, r *http.Request) {
// 	page := &Page{Title: "p.UserName", Body: []byte("not Logged in")}
// 	render(w, "staff3.html", page)
// }
