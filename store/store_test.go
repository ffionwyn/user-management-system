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
		t.Error(`err := AddToStorage(firstName, secondName, email, dob) error when adding to storage, please check inputs`)
	}
	exists := CheckPerson(email)
	if !exists {
		t.Error(`Error: person hasn't been added`)
	}
}

func TestGetPerson(t *testing.T) {
	firstName := "ffion"
	secondName := "griffiths"
	email := "fgriffiths@example.com"
	dob := "05/11/1993"

	// add a person to storage for testing
	err := AddToStorage(firstName, secondName, email, dob)
	if err != nil {
		t.Errorf("Error adding person to storage: %v", err)
	}

	// get the person details
	getFirstName, getSecondName, getDOB, err := GetPerson(email)
	if err != nil {
		t.Errorf("Error getting person details: %v", err)
	}

	// check that the get details match the original person
	if getFirstName != firstName {
		t.Errorf("Expected first name: %s, got: %s", firstName, getFirstName)
	}

	if getSecondName != secondName {
		t.Errorf("Expected second name: %s, got: %s", secondName, getSecondName)
	}

	if getDOB != dob {
		t.Errorf("Expected DOB: %s, got: %s", dob, getDOB)
	}

	// trying getting details for a non-existent person
	nonExistentEmail := "nonexistent@example.com"
	_, _, _, err = GetPerson(nonExistentEmail)
	if err == nil {
		t.Errorf("Expected error for non-existent person, but got nil")
	} else {
		expectedErrorMessage := "Person does not exist"
		if err.Error() != expectedErrorMessage {
			t.Errorf("Expected error message: %s, got: %s", expectedErrorMessage, err.Error())
		}
	}
}


