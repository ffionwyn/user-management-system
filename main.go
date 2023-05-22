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
		email := r.URL.Query().Get("email")
		if email != "" {
			getPerson(w, r)
		} else {
			searchUser(w, r)
		}
	case "PATCH":
		updatePerson(w, r)
	case "DELETE":
		DeletePerson(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
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

func getPerson(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	firstName, secondName, dob, err := store.GetPerson(email)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	message := fmt.Sprintf("firstName: %s, secondName: %s, dob %s", firstName, secondName, dob)
	w.Write([]byte(message))
	w.WriteHeader(http.StatusOK)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	inputtedFirstName := r.URL.Query().Get("firstName")
	inputtedSecondName := r.URL.Query().Get("secondName")
	inputtedEmail := r.URL.Query().Get("email")
	inputtedDOB := r.URL.Query().Get("dob")
	err := store.UpdatePersonStorage(inputtedFirstName, inputtedSecondName, inputtedEmail, inputtedDOB)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	err := store.DeletePerson(email)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func searchUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	searchResult := "result from email search: " + email

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, searchResult)
}
