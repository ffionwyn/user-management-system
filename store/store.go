package store

import (
	"errors"
	"fmt"
	"log"
)

var PersonStorage = make(map[string]Person)

type Person struct {
	FirstName  string
	SecondName string
	DOB        string
	Email string
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