package store

import (
	"errors"
	"fmt"
	"log"
)

type Person struct {
	firstName string
	secondName string
	dob string
}

func newPerson(firstName string, secondName string, dob string) Person {
	p := Person{
		firstName: firstName,
		secondName: secondName,
		dob: dob,
	}
	return p
}

func AddToStorage(firstName string, secondName string, email string, dob string) error {
	validationErr := validateInput(firstName, secondName, email, dob)
	if validationErr != nil {
		return validationErr
	}
	fmt.Println("Hello " + firstName + " ")
	p := newPerson(firstName, secondName, dob)
	personStorage[email] = p
	log.Println("Added to storage successful")
	return nil
}

var personStorage = make(map[string]Person)

func validateInput(firstName string, secondName string, email string, dob string) error {
	if firstName == "" {
		return errors.New("missing name parameter")
	}
	if secondName == "" {
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

func GetPerson(email string) (string, string, string, error) {
	Person, found := personStorage[email]
	if !found {
		return "", "", "", fmt.Errorf("person does not exist")
	}
	return Person.firstName, Person.secondName, Person.dob, nil
}