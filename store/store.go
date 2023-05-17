package store

import (
	"errors"
	"fmt"
	"log"
)

type Person struct {
	email string
	dob string
}

func newPerson(email string, dob string) Person {
	p := Person{
		email: email,
		dob: dob,
	}
	return p
}

func AddToStorage(name string, email string, dob string) error {
	validationErr := validateInput(name, email, dob)
	if validationErr != nil {
		return validationErr
	}
	fmt.Println("Hello " + name + " ")
	p := newPerson(email, dob)
	personStorage[name] = p
	log.Println("Added to storage successful")
	return nil
}

var personStorage = make(map[string]Person)

func validateInput(name string, email string, dob string) error {
	if name == "" {
		return errors.New("missing name parameter")
	}
	if email == "" {
		return errors.New("missing email parameter")
	}
	if dob == "" {
		return errors.New("missing dob parameter")
	}
	return nil
}

func GetPerson(name string) (string, string, error) {
	Person, found := personStorage[name]
	if !found {
		return "", "", fmt.Errorf("person does not exist")
	}
	return Person.email, Person.dob, nil
}