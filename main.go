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
	fmt.Println("methodLogin:", r.Method) //get request method
	r.ParseForm()
	// logic part of log in

	uname := strings.Join(r.Form["username"], " ")
	psw := strings.Join(r.Form["password"], " ")

	for _, p := range users {
		fmt.Println("huhuh")
		if p.username == uname && p.password == psw {
			fmt.Println("p: " + p.username + " " + p.password)
			fmt.Println("************************************************************")
			http.Redirect(w, r, "/signIn", http.StatusFound)

		}
	}

	t, _ := template.ParseFiles("login.html")
	t.Execute(w, nil)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("methodSignUP:", r.Method)
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
	fmt.Println("methodSignIn:", r.Method)
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
