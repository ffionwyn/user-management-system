package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// test case 1 – getting an existing user by making a get request (/user/1)
// test case 2 – getting a none existing user (/user/3)
// create a new request, create new response recorder, create a gin test and then assign the response recorder to it.
// call the getUser function with the test context, verify the expected results and then assert the response status code and response body
func TestGetUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "/user/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	getUser(c)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{
		"FirstName": "ffion",
		"SecondName": "griffiths",
		"DOB": "05/11/1993",
		"Email": "ffiongriffiths@example.com"
	}`
	assert.JSONEq(t, expectedBody, w.Body.String())

	req, _ = http.NewRequest("GET", "/user/3", nil)
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request = req

	getUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	expectedError := `{"message":"user not found"}`
	assert.JSONEq(t, expectedError, w.Body.String())
}

