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

func TestUpdatePersonStorage(t *testing.T) {
	firstName := "ffion"
	secondName := "griffiths"
	email := "fgriffiths@example.com"
	dob := "05/11/1993"

	// add a person to storage for testing
	err := AddToStorage(firstName, secondName, email, dob)
	if err != nil {
		t.Errorf("Error adding person to storage: %v", err)
	}

	// update the person details
	newFirstName := "minnie"
	newSecondName := "griffiths"
	newDOB := "18/11/2018"
	err = UpdatePersonStorage(newFirstName, newSecondName, email, newDOB)
	if err != nil {
		t.Errorf("Error updating person details: %v", err)
	}

	// get the updated person details
	getFirstName, getSecondName, getDOB, err := GetPerson(email)
	if err != nil {
		t.Errorf("Error getting person details: %v", err)
	}

	// check that the details match the updated values
	if getFirstName != newFirstName {
		t.Errorf("Expected first name: %s, got: %s", newFirstName, getFirstName)
	}

	if getSecondName != newSecondName {
		t.Errorf("Expected second name: %s, got: %s", newSecondName, getSecondName)
	}

	if getDOB != newDOB {
		t.Errorf("Expected DOB: %s, got: %s", newDOB, getDOB)
	}

	// try updating details for a non-existent person
	nonExistentEmail := "nonexistent@example.com"
	err = UpdatePersonStorage(newFirstName, newSecondName, nonExistentEmail, newDOB)
	if err == nil {
		t.Errorf("Expected error for non-existent person, but got nil")
	} else {
		expectedErrorMessage := "person does not exist"
		if err.Error() != expectedErrorMessage {
			t.Errorf("Expected error message: %s, got: %s", expectedErrorMessage, err.Error())
		}
	}
}

func TestDeletePerson(t *testing.T) {
	firstName := "ffion"
	secondName := "griffiths"
	email := "fgriffiths@example.com"
	dob := "05/11/1993"

	// add a person to storage for testing
	err := AddToStorage(firstName, secondName, email, dob)
	if err != nil {
		t.Errorf("Error adding person to storage: %v", err)
	}

	// delete the person from storage
	err = DeletePerson(email)
	if err != nil {
		t.Errorf("Error deleting person: %v", err)
	}

	// try getting the deleted person's details
	_, _, _, err = GetPerson(email)
	if err == nil {
		t.Errorf("Expected error for non-existent person, but got nil")
	} else {
		expectedErrorMessage := "person does not exist"
		if err.Error() != expectedErrorMessage {
			t.Errorf("Expected error message: %s, got: %s", expectedErrorMessage, err.Error())
		}
	}
}
