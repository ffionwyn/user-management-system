package main

import (
	"fmt"
	"log"
	"net/http"
	"user-management/store"

	"github.com/bmizerany/pat"
)

func main() {
	m := pat.New()
	m.Get("/user/:id", http.HandlerFunc(UserHandler))
	m.Post("/new-user", http.HandlerFunc(newUserHandler))

	http.Handle("/", m)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}	



func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUserByID(w, r)
	case "PATCH":
		updatePerson(w, r)
	case "DELETE":
		DeletePerson(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry, only POST/GET/PATCH/DELETE methods are supported.")
	}
}

func newUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		addUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry, only POST/GET/PATCH/DELETE methods are supported.")
	}
}

func addUser(w http.ResponseWriter, r *http.Request) {
	inputtedFirstName := r.URL.Query().Get("firstName")
	inputtedSecondName := r.URL.Query().Get("secondName")
	inputtedEmail := r.URL.Query().Get("email")
	inputtedDOB := r.URL.Query().Get("dob")
	err := store.AddToStorage(inputtedFirstName, inputtedSecondName, inputtedEmail, inputtedDOB)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	person, err := store.GetPersonByID(id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	response := fmt.Sprintf("firstName: %s, secondName: %s, dob: %s",
		person.FirstName, person.SecondName, person.DOB)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, response)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	inputtedUserID := r.URL.Query().Get(":id")
	inputtedFirstName := r.URL.Query().Get("firstName")
	inputtedSecondName := r.URL.Query().Get("secondName")
	inputtedEmail := r.URL.Query().Get("email")
	inputtedDOB := r.URL.Query().Get("dob")
	err := store.UpdatePersonStorage(inputtedUserID, inputtedFirstName, inputtedSecondName, inputtedEmail, inputtedDOB)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	inputtedUserID := r.URL.Query().Get(":id")
	err := store.DeletePerson(inputtedUserID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
