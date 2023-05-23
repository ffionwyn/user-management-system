package store

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

var counter int
var PersonStorage = make(map[string]Person)

type Person struct {
	UserID     string
	FirstName  string
	SecondName string
	DOB        string
}

func newPerson(UserID string, FirstName string, SecondName string, DOB string) Person {
	p := Person{
		UserID:     UserID,
		FirstName:  FirstName,
		SecondName: SecondName,
		DOB:        DOB,
	}
	return p
}

func AddToStorage(FirstName string, SecondName string, Email string, DOB string) error {
	validationErr := validateInput(FirstName, SecondName, Email, DOB)
	if validationErr != nil {
		return validationErr
	}
	UserID := strconv.Itoa(counter)
	p := newPerson(UserID, FirstName, SecondName, DOB)
	PersonStorage[UserID] = p
	counter++
	log.Println("Added to storage successful")
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

func GetPerson(UserID string) (string, string, string, error) {
	Person, found := PersonStorage[UserID]
	if !found {
		return "", "", "", fmt.Errorf("person does not exist")
	}
	return Person.FirstName, Person.SecondName, Person.DOB, nil
}

func UpdatePersonStorage(UserID string, FirstName string, SecondName string, Email string, DOB string) error {
	validationErr := validateInput(FirstName, SecondName, Email, DOB)
	if validationErr != nil {
		return validationErr
	}
	_, ok := PersonStorage[UserID]
	if !ok {
		log.Print("person not in storage - failed to update")
		return fmt.Errorf("person does not exist")
	}
	p := newPerson(UserID, FirstName, SecondName, DOB)
	PersonStorage[UserID] = p
	log.Println("Update person successful")
	return nil
}

func DeletePerson(UserID string) error {
	if _, ok := PersonStorage[UserID]; !ok {
		return errors.New("person (userID) does not exist")
	}
	delete(PersonStorage, UserID)
	log.Println("Delete person successful")
	return nil
}

func CheckPerson(UserID string) bool {
	_, exists := PersonStorage[UserID]
	return exists
}

func GetPersonByID(id string) (Person, error) {
	person, found := PersonStorage[id]
	if !found {
		return Person{}, fmt.Errorf("user not found")
	}
	return person, nil
}