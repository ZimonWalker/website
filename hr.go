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

// HRList struct
type HRList struct {
	Username  string      `json:"Username"`
	StaffPage []StaffPage `json:"StaffPage"`
}

// HRLeave struct
type HRLeave struct {
	Username   string       `json:"Username"`
	StaffLeave []StaffLeave `json:"StaffLeave"`
}

var hp = &HrPage{}

//Hr func
func Hr(w http.ResponseWriter, r *http.Request) {

	if gp.Title != "loggedHr" || gp.Title == "" {
		http.Redirect(w, r, "/", http.StatusFound)

		return
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
	if gp.Title != "loggedHr" || gp.Title == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	files, err := ioutil.ReadDir("./database/staff")
	if err != nil {
		log.Fatalln(err)
	}

	i := 0
	var hlist = &HRList{Username: gp.Body}

	for _, f := range files {
		// fmt.Println(f.Name(), i)
		var miniSL StaffPage
		miniSL.Index = i
		db := "./database/staff/" + f.Name()
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

		hlist.StaffPage = append(hlist.StaffPage, miniSL)

		i++
	}

	renderHR2(w, "hr2.html", hlist)
}

func renderHR2(w http.ResponseWriter, tmpl string, p *HRList) {
	err := templates.ExecuteTemplate(w, tmpl, p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

// Hr3 func
func Hr3(w http.ResponseWriter, r *http.Request) {
	if gp.Title != "loggedHr" || gp.Title == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	files, err := ioutil.ReadDir("./database/leave")
	if err != nil {
		log.Fatalln(err)
	}

	i := 0
	var hl = &HRLeave{Username: gp.Body}

	for _, f := range files {
		// fmt.Println(f.Name(), i)
		var miniSL StaffLeave
		miniSL.Index = i
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

	renderHR3(w, "hr3.html", hl)
}

func renderHR3(w http.ResponseWriter, tmpl string, p *HRLeave) {
	err := templates.ExecuteTemplate(w, tmpl, p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
