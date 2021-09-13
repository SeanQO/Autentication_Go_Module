package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

var users = []*person{}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// logic part of log in

	uname := strings.Join(r.Form["username"], " ")
	psw := strings.Join(r.Form["password"], " ")

	fmt.Println("Useruname: ", uname)
	fmt.Println("Userpsw: ", psw)

	for _, p := range users {
		fmt.Println("uname: ", (strings.Compare(p.username, uname)))
		fmt.Println("psw: ", (strings.Compare(p.password, psw)))
		if (strings.Compare(p.username, uname))+(strings.Compare(p.password, psw)) == 0 {
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
	r.ParseForm()
	username := strings.Join(r.Form["username"], " ")
	password := strings.Join(r.Form["psw"], " ")
	fname := strings.Join(r.Form["fname"], " ")
	lname := strings.Join(r.Form["lname"], " ")
	birthDate := strings.Join(r.Form["birthDate"], " ")
	p := person{username: username, password: password, fname: fname, lname: lname, birthDate: birthDate}
	users = append(users, &p)

	http.Redirect(w, r, "/login", http.StatusFound)

}

func signIn(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("signIn.html")
	t.Execute(w, nil)
	for _, p := range users {
		fmt.Println("UN: " + p.username + "PS: " + p.password)
	}
	r.ParseForm()

}

func main() {
	users = append(users, &person{username: "admin", password: "admin", fname: "admin", lname: "admin", birthDate: "3/3/3"})
	http.HandleFunc("/login", login)
	http.HandleFunc("/signUp", signUp)
	http.HandleFunc("/signIn", signIn)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
