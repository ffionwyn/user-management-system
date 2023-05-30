package store

import (
	"testing"
)

// person is created, each field is individually checked against the expected value (using if)
//this is checking if firstName is equal to ffion (the expected name) and repeated for secondName, DOB and email
func TestPerson(t *testing.T) {
	person := Person{
		FirstName:  "ffion",
		SecondName: "griffiths",
		DOB:        "05/11/1993",
		Email:      "ffiongriffiths@example.com",
	}

	if person.FirstName != "ffion" {
		t.Errorf("Expected FirstName to be 'ffion', but got '%s'", person.FirstName)
	}

	if person.SecondName != "griffiths" {
		t.Errorf("Expected SecondName to be 'griffiths', but got '%s'", person.SecondName)
	}

	if person.DOB != "05/11/1993" {
		t.Errorf("Expected DOB to be '05/11/1993', but got '%s'", person.DOB)
	}

	if person.Email != "ffiongriffiths@example.com" {
		t.Errorf("Expected Email to be 'ffiongriffiths@example.com', but got '%s'", person.Email)
	}
}

// test case 1 is deleting an existing person and then checks if they have been deleted from the storage
// test case 2 is deleting a non existing person
func TestDeletePerson(t *testing.T) {
	PersonStorage := map[string]Person{
		"1": {
			FirstName:  "ffion",
			SecondName: "griffiths",
			DOB:        "05/11/1993",
			Email:      "ffiongriffiths@example.com",
		},
		"2": {
			FirstName:  "minnie",
			SecondName: "griffiths",
			DOB:        "19/10/2018",
			Email:      "minniegriffiths@example.com",
		},
	}

	err := DeletePerson("1")
	if err != nil {
		t.Errorf("Expected nil error, but got '%v'", err)
	}

	if _, ok := PersonStorage["1"]; ok {
		t.Error("Expected person with UserID '1' to be deleted, but it still exists")
	}

	err = DeletePerson("3")
	expectedError := "person (userID) does not exist"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error '%s', but got '%v'", expectedError, err)
	}
}


