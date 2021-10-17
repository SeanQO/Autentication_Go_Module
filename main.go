package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type person struct {
	username  string
	password  string
	fname     string
	lname     string
	birthDate string
}

type Data struct {
	Items []person
}

var users = []person{}

func login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	// logic part of log in
	uname := strings.Join(r.Form["username"], " ")
	psw := strings.Join(r.Form["password"], " ")

	for _, p := range users {
		if (strings.Compare(p.username, uname))+(strings.Compare(p.password, psw)) == 0 {
			uname = ""
			psw = ""
			http.Redirect(w, r, "/signIn", http.StatusFound)
		} else {
			fmt.Println("INVALID LOGIN")
		}
	}

	t, _ := template.ParseFiles("login.html")
	t.Execute(w, nil)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("signUp.html")
	t.Execute(w, nil)
	if r.Method == "POST" {
		r.ParseForm()
		username := strings.Join(r.Form["username"], " ")
		password := strings.Join(r.Form["psw"], " ")
		fname := strings.Join(r.Form["fname"], " ")
		lname := strings.Join(r.Form["lname"], " ")
		birthDate := strings.Join(r.Form["birthDate"], " ")
		p := person{username: username, password: password, fname: fname, lname: lname, birthDate: birthDate}
		users = append(users, p)

		saveData()

		http.Redirect(w, r, "/login", http.StatusFound)

	}

}

func signIn(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("signIn.html")
	t.Execute(w, nil)
	r.ParseForm()

}

func readData() {
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	users = updateUsers(data)

}

func updateUsers(data [][]string) []person {
	var usersList []person
	for i, line := range data {
		if i > 0 {
			var rec person
			for j, field := range line {
				if j == 0 {
					rec.username = field
				} else if j == 1 {
					rec.password = field
				} else if j == 2 {
					rec.fname = field
				} else if j == 2 {
					rec.lname = field
				} else if j == 2 {
					rec.birthDate = field
				}

			}
			usersList = append(usersList, rec)
		}
	}
	return usersList
}

func saveData() {
	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// initialize csv writer
	writer := csv.NewWriter(file)

	defer writer.Flush()

	data := [][]string{}

	for _, p := range users {
		d := []string{p.username, p.password, p.fname, p.lname, p.birthDate}

		data = append(data, d)

	}
	// write all rows at once
	writer.WriteAll(data)

}

func main() {
	readData()
	http.HandleFunc("/login", login)
	http.HandleFunc("/signUp", signUp)
	http.HandleFunc("/signIn", signIn)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
