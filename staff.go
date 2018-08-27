package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// StaffPage struct
type StaffPage struct {
	Index        int    `json:"Index"`
	Username     string `json:"Username"`
	Password     string `json:"Password"`
	Email        string `json:"Email"`
	FullName     string `json:"FullName"`
	Gender       string `json:"Gender"`
	IC           string `json:"IC"`
	Phone        string `json:"Phone"`
	LeaveBalance int64  `json:"LeaveBalance"`
	LeaveID      string `json:"LeaveID"`
	Role         string `json:"Role"`
}

// StaffLeave struct
type StaffLeave struct {
	Index        int    `json:"Index"`
	ByName       string `json:"ByName"`
	ByFullName   string `json:"ByFullName"`
	ByEmail      string `json:"ByEmail"`
	LeaveBalance int64  `json:"LeaveBalance"`
	LeaveType    string `json:"LeaveType"`
	ApplyDate    string `json:"ApplyDate"`
	StartDate    string `json:"StartDate"`
	EndDate      string `json:"EndDate"`
	NumDays      int64  `json:"NumDays"`
	Remark       string `json:"Remark"`
	Status       string `json:"Status"`
}

var sp = &StaffPage{}
var sl = &StaffLeave{}

//Staff func
func Staff(w http.ResponseWriter, r *http.Request) {
	if gp.Title != "loggedStaff" || gp.Title == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	u := gp.Body

	db := "./database/staff/" + u + ".json"
	var content []byte
	content, err := ioutil.ReadFile(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Parsing/Unmarshalling JSON encoding/json
	if err = json.Unmarshal(content, &sp); err != nil {
		log.Fatalln(err)
		// panic(err)
	}

	renderStaff2(w, "staff.html", sp)
}

//Staff2 func
func Staff2(w http.ResponseWriter, r *http.Request) {
	if gp.Title != "loggedStaff" || gp.Title == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	renderStaff2(w, "staff2.html", sp)
}

func renderStaff2(w http.ResponseWriter, tmpl string, p *StaffPage) {
	err := templates.ExecuteTemplate(w, tmpl, p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

//Staff3 func
func Staff3(w http.ResponseWriter, r *http.Request) {
	if gp.Title != "loggedStaff" || gp.Title == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	u := sp.LeaveID

	if u != "" {
		db := "./database/leave/" + u + ".json"
		var content []byte
		content, err := ioutil.ReadFile(db)
		if err != nil {
			log.Fatalln(err)
		}

		// Parsing/Unmarshalling JSON encoding/json
		if err = json.Unmarshal(content, &sl); err != nil {
			log.Fatalln(err)
			// panic(err)
		}
	}

	renderStaff3(w, "staff3.html", sl)
}

func renderStaff3(w http.ResponseWriter, tmpl string, p *StaffLeave) {
	err := templates.ExecuteTemplate(w, tmpl, p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
