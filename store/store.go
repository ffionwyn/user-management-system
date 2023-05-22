package store

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

var counter int 
var users []Person
var personStorage = make(map[string]Person)
type Person struct {
	ID int
	FirstName string
	SecondName string
	DOB string
}

func newPerson(FirstName string, SecondName string, DOB string) Person {
	p := Person{
		FirstName: FirstName,
		SecondName: SecondName,
		DOB: DOB,
	}
	return p
}

func AddToStorage(FirstName string, SecondName string, Email string, DOB string) error {
	validationErr := validateInput(FirstName, SecondName, Email, DOB)
	if validationErr != nil {
		return validationErr
	}
	fmt.Println("Hello " + FirstName + " ")
	p := newPerson(FirstName, SecondName, DOB)
	personStorage[Email] = p
	log.Println("Added to storage successful")
	counter++
	return nil
}

func validateInput(FirstName string, SecondName string, Email string, DOB string) error {
	if FirstName == "" {
		return errors.New("missing name parameter")
	}
	if SecondName == "" {
		return errors.New("missing name parameter")
	}
	if Email == "" {
		return errors.New("missing Email parameter")
	}
	if DOB == "" {
		return errors.New("missing DOB parameter")
	}
	return nil
}

func GetPerson(Email string) (string, string, string, error) {
	Person, found := personStorage[Email]
	if !found {
		return "", "", "", fmt.Errorf("person does not exist")
	}
	return Person.FirstName, Person.SecondName, Person.DOB, nil
}

func UpdatePersonStorage(FirstName string, SecondName string, Email string, DOB string) error {
	validationErr := validateInput(FirstName, SecondName, Email, DOB)
	if validationErr != nil {
		return validationErr
	}
	_, ok := personStorage[Email]
	if !ok {
		log.Print("person not in storage - failed to update")
		return fmt.Errorf("person does not exist")
	}
	p := newPerson(FirstName, SecondName, DOB)
	personStorage[Email] = p
	log.Println("Update person successful")
	return nil
}

func DeletePerson(Email string) error {
	if _, ok := personStorage[Email]; !ok {
		return errors.New("person (Email) does not exist")
	}
	delete(personStorage, Email)
	log.Println("Delete person successful")
	return nil
}

func CheckPerson(Email string) bool {
	_, exists := personStorage[Email]
	return exists
}

func GetPersonByID(id string) (Person, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return Person{}, fmt.Errorf("invalid user ID")
	}

	if userID <= 0 || userID > counter {
		return Person{}, fmt.Errorf("user not found")
	}
	person := users[userID-1]

	return person, nil
}