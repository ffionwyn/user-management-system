package store

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewPerson(t *testing.T) {
	expectedUserID := "1"
	expectedFirstName := "ffion"
	expectedSecondName := "griffiths"
	expectedDOB := "05/11/1993"

	p := newPerson(expectedUserID, expectedFirstName, expectedSecondName, expectedDOB)
	if p.UserID != expectedUserID {
		t.Errorf("expected UserID %s, but got %s", expectedUserID, p.UserID)
	}
	if p.FirstName != expectedFirstName {
		t.Errorf("expected FirstName %s, but got %s", expectedFirstName, p.FirstName)
	}
	if p.SecondName != expectedSecondName {
		t.Errorf("expected SecondName %s, but got %s", expectedSecondName, p.SecondName)
	}
	if p.DOB != expectedDOB {
		t.Errorf("expected DOB %s, but got %s", expectedDOB, p.DOB)
	}
}

func TestValidateInput(t *testing.T) {
	err := ValidateInput("ffion", "griffiths", "fgriffiths@example.com", "05/11/1993")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	err = ValidateInput("", "griffiths", "fgriffiths@example.com", "05/11/1993")
	expectedError := errors.New("missing name parameter")
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}

	err = ValidateInput("ffion", "", "fgriffiths@example.com", "05/11/1993")
	expectedError = errors.New("missing name parameter")
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}

	err = ValidateInput("ffion", "griffiths", "", "05/11/1993")
	expectedError = errors.New("missing Email parameter")
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}

	err = ValidateInput("John", "Doe", "johndoe@example.com", "")
	expectedError = errors.New("missing DOB parameter")
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}
}

func TestGetPerson(t *testing.T) {
	UserID := "1"
	expectedFirstName := "ffion"
	expectedSecondName := "griffiths"
	expectedDOB := "05/11/1993"
	PersonStorage[UserID] = Person{
		UserID:     UserID,
		FirstName:  expectedFirstName,
		SecondName: expectedSecondName,
		DOB:        expectedDOB,
	}

	firstName, secondName, dob, err := GetPerson(UserID)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if firstName != expectedFirstName || secondName != expectedSecondName || dob != expectedDOB {
		t.Errorf("Expected values: %s, %s, %s; but got: %s, %s, %s",
			expectedFirstName, expectedSecondName, expectedDOB, firstName, secondName, dob)
	}
	UserID = "2"
	firstName, secondName, dob, err = GetPerson(UserID)
	expectedError := fmt.Errorf("person does not exist")
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedError, err)
	}
	if firstName != "" || secondName != "" || dob != "" {
		t.Errorf("Expected empty values, but got: %s, %s, %s", firstName, secondName, dob)
	}
}