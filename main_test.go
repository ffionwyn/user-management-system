package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"user-management/store"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type Person struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	DOB        string `json:"dob"`
	Email      string `json:"email"`
}

// create a new Gin router and register a GET route with "/users/:id" and assign the getUser function as the handler.
// create a new instance of the response recorder, which is used to capture the response.
// use "/users/123" as the URL for the nil HTTP request. The router handles the request and generates the response.
// assert that the response status code captured in the variable is equal to the HTTP status code (404 in this case).
// declare responseBody to store the JSON response body.
// attempt to parse the response body captured as JSON and store it in responseBody. If an error occurs, it will be assigned to the err variable.
// check if there was an error during JSON parsing and log it if so.
// declare expected as a map[string]interface{} with the expected JSON response body.
// assert that responseBody matches the expected response body. This ensures that the actual response body matches the expected response body, indicating that the "getUser" function returned the correct response for the tests.

func TestGetUser(t *testing.T) {
	router := gin.Default()
	router.GET("/users/:id", getUser)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users/123", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		log.Fatal(err)
	}

	expected := map[string]interface{}{
		"message": "User not found",
	}
	assert.Equal(t, expected, responseBody)
}

// create a new Gin router and register a GET route with "/users" and assign the getAllUsers function as the handler.
// create a new instance of the response recorder, which is used to capture the response.
// set up mock data in the personStorage map (store package), which is a collection of users.
// create a responseRecorder to capture the response from the handler.
// send a new GET request to "/users" using NewRequest.
// capture the response using router.ServeHTTP.
// assert the response status code (http.StatusOK in this case).
// declare the variable users as a slice of store.Person to store the response body.
// attempt to unmarshal the response body and assert that there is no error during this process.
// define the expected list of users as a slice and assert that the parsed users match the expected ones.
func TestGetAllUsers(t *testing.T) {
	router := gin.Default()
	router.GET("/users", getAllUsers)

	store.PersonStorage = map[string]store.Person{
		"1": {FirstName: "ffion", SecondName: "griffiths", DOB: "05/11/1993", Email: "ffiongriffiths@example.com"},
		"2": {FirstName: "minnie", SecondName: "griffiths", DOB: "19/10/2018", Email: "minniegriffiths@example.com"},
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []store.Person
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)

	expected := []store.Person{
		{FirstName: "ffion", SecondName: "griffiths", DOB: "05/11/1993", Email: "ffiongriffiths@example.com"},
		{FirstName: "minnie", SecondName: "griffiths", DOB: "19/10/2018", Email: "minniegriffiths@example.com"},
	}
	assert.Equal(t, expected, users)
}

// create a new Gin router, define a route for the POST request with the path "/users", and assign the postUser function as the handler for this route.
// initialize the personStorage map, which represents the storage for user data.
// create a request with a JSON payload. This creates an HTTP request with the POST method, the URL "/users", and the provided JSON payload.
// create a response recorder, which captures the response generated by the router.
// serve the request and capture the response.
// assert that the response status code is http.StatusCreated (201 in this case), indicating that the user was successfully created.
// attempt to parse the response body as JSON and store it in the responsePerson variable. Ensure there are no errors during the unmarshaling process.
// assert that the personStorage map contains a user with the key "1", showing that the user was stored successfully.
// assert that the stored person matches the expected person data. Check that all details of a "person" match the provided values.
func TestPostUser(t *testing.T) {
	router := gin.Default()
	router.POST("/users", postUser)

	store.PersonStorage = make(map[string]store.Person)

	requestBody := `{
		"firstName": "ffion",
		"secondName": "griffiths",
		"dob": "05/11/1993",
		"email": "ffiongriffiths@example.com"
	}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var responsePerson store.Person
	err := json.Unmarshal(w.Body.Bytes(), &responsePerson)
	assert.NoError(t, err)

	assert.Contains(t, store.PersonStorage, "1")

	assert.Equal(t, store.Person{
		FirstName:  "ffion",
		SecondName: "griffiths",
		DOB:        "05/11/1993",
		Email:      "ffiongriffiths@example.com",
	}, responsePerson)
}

func TestUpdateUser(t *testing.T) {
	router := gin.Default()
	router.PUT("/users/:id", updateUser)

	store.PersonStorage = make(map[string]store.Person)
	store.PersonStorage["1"] = store.Person{
		FirstName:  "ffion",
		SecondName: "griffiths",
		DOB:        "05/11/1993",
		Email:      "ffiongriffiths@example.com",
	}

	requestBody := `{
		"firstName": "updatedFirstName",
		"secondName": "updatedSecondName",
		"dob": "01/01/2000",
		"email": "updated@example.com"
	}`
	req, _ := http.NewRequest("PUT", "/users/1", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var responsePerson store.Person
	err := json.Unmarshal(w.Body.Bytes(), &responsePerson)
	assert.NoError(t, err)

	updatedUser, ok := store.PersonStorage["1"]
	assert.True(t, ok)
	assert.Equal(t, "updatedFirstName", updatedUser.FirstName)
	assert.Equal(t, "updatedSecondName", updatedUser.SecondName)
	assert.Equal(t, "01/01/2000", updatedUser.DOB)
	assert.Equal(t, "updated@example.com", updatedUser.Email)
}

