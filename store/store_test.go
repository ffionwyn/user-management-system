package store

import (
	"testing"
)

//Test for adding someone to storage.
func TestAddToStorage(t *testing.T) {
	firstName := "ffion"
	secondName := "griffiths"
	email := "fgriffiths@example.com"
	dob := "05/11/1993"
	err := AddToStorage(firstName, secondName, email, dob)
	if err != nil {
		t.Error(`err := AddToStorage(firstName, secondName, email, dob) not equal to nil`)
	}
	exists := CheckPerson(email)
	if !exists {
		t.Error(`error: person hasn't been added`)
	}
}

