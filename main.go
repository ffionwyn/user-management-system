package main

import (
	"fmt"
	"log"
	"net/http"
	"user-management/store"
)

func main() {
	http.HandleFunc("/user", UserHandler)
	http.ListenAndServe(":5000", nil)
}


func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		addUser(w, r)
	case "GET":
		getPerson(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func addUser(w http.ResponseWriter, r *http.Request) {
	inputtedUser := r.URL.Query().Get("name")
	inputtedEmail := r.URL.Query().Get("email")
	inputtedDOB := r.URL.Query().Get("dob")
	err := store.AddToStorage(inputtedUser, inputtedEmail, inputtedDOB)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	email, dob, err := store.GetPerson(name)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	message := fmt.Sprintf("email: %s, dob %s", email, dob)
	w.Write([]byte(message))
	w.WriteHeader(http.StatusOK)
}