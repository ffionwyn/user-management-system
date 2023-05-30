package store

import (
	"testing"
)

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
