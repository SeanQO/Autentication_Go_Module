package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type person struct {
	username  string
	fname     string
	lName     string
	birthDate time.Time
}

var users = []*person{}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("methodLogin:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("methodSignUP:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("signUp.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

	}

}

func signIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("methodSignIn:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("signIn.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

	}

}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/signUp", signUp)
	http.HandleFunc("/signIn", signIn)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
