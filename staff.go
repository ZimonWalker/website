package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// StaffPage struct
type StaffPage struct {
	Index        int      `json:"Index"`
	Username     string   `json:"Username"`
	Password     string   `json:"Password"`
	Email        string   `json:"Email"`
	FullName     string   `json:"FullName"`
	Gender       string   `json:"Gender"`
	IC           string   `json:"IC"`
	Phone        string   `json:"Phone"`
	LeaveBalance int      `json:"LeaveBalance"`
	LeaveID      []string `json:"LeaveID"`
	Role         string   `json:"Role"`
}

// StaffLeave struct
type StaffLeave struct {
	Index        int    `json:"Index"`
	ID           string `json:"ID"`
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

	getPath := r.URL.Path[len("/staff2/"):]
	// fmt.Println(getPath)

	if getPath == "updateLeave" {

		ApplyDate := r.FormValue("ApplyDate")
		ByEmail := r.FormValue("Email")
		ByFullName := r.FormValue("FullName")
		ByName := r.FormValue("Username")
		EndDate := r.FormValue("end_date")
		ID := genXid()
		LeaveBalance := r.FormValue("LeaveBalance")
		LeaveType := r.FormValue("leave_type")
		NumDays := r.FormValue("NumDays")
		Remark := r.FormValue("remark")
		StartDate := r.FormValue("start_date")
		Status := "Pending"

		db := "./database/leave/" + ID + ".json"

		var LeaveBalance64 int64
		var NumDays64 int64

		if i, err := strconv.ParseInt(LeaveBalance, 10, 64); err == nil {
			LeaveBalance64 = i
		}
		if i, err := strconv.ParseInt(NumDays, 10, 64); err == nil {
			NumDays64 = i
		}

		// Creating the maps for JSON
		m := StaffLeave{
			ApplyDate:    ApplyDate,
			ByEmail:      ByEmail,
			ByFullName:   ByFullName,
			ByName:       ByName,
			EndDate:      EndDate,
			ID:           ID,
			LeaveBalance: LeaveBalance64,
			LeaveType:    LeaveType,
			NumDays:      NumDays64,
			Remark:       Remark,
			StartDate:    StartDate,
			Status:       Status,
		}

		// fmt.Println(m)

		b, err := json.Marshal(m)
		if err != nil {
			log.Fatalln(err)
		}
		if err = ioutil.WriteFile(db, b, 0644); err != nil {
			log.Fatalln(err)
		}

		sp.LeaveID = append(sp.LeaveID, ID)

		db = "./database/staff/" + ByName + ".json"

		b, err = json.Marshal(sp)
		if err != nil {
			log.Fatalln(err)
		}
		if err = ioutil.WriteFile(db, b, 0644); err != nil {
			log.Fatalln(err)
		}

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

	i := 0
	// s := strconv.FormatInt(sp.LeaveBalance, 64)
	s := strconv.Itoa(sp.LeaveBalance)
	var hl = &HRLeave{Username: s}

	for _, id := range sp.LeaveID {
		var miniSL StaffLeave
		i++
		db := "./database/leave/" + id + ".json"
		content, err := ioutil.ReadFile(db)
		if err != nil {
			log.Fatalln(err)
		}

		if err = json.Unmarshal(content, &miniSL); err != nil {
			log.Fatalln(err)
			// panic(err)
		}
		miniSL.Index = i

		hl.StaffLeave = append(hl.StaffLeave, miniSL)
	}

	renderStaff3(w, "staff3.html", hl)
}

func renderStaff3(w http.ResponseWriter, tmpl string, p *HRLeave) {
	err := templates.ExecuteTemplate(w, tmpl, p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
