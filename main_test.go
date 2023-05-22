package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler(t *testing.T) {
	// create a new HTTP request with the selected method and path
	postRequest, _ := http.NewRequest("POST", "/user", nil)
	getRequest, _ := http.NewRequest("GET", "/user", nil)
	patchRequest, _ := http.NewRequest("PATCH", "/user", nil)
	deleteRequest, _ := http.NewRequest("DELETE", "/user", nil)
	invalidMethodRequest, _ := http.NewRequest("PUT", "/user", nil)

// create a writer to capture the HTTP response
	responseRecorder := httptest.NewRecorder()

	// call the UserHandler function with each request and the writer
	UserHandler(responseRecorder, postRequest)
	checkResponseStatus(t, responseRecorder, http.StatusOK)

	UserHandler(responseRecorder, getRequest)
	checkResponseStatus(t, responseRecorder, http.StatusOK)

	UserHandler(responseRecorder, patchRequest)
	checkResponseStatus(t, responseRecorder, http.StatusOK)

	UserHandler(responseRecorder, deleteRequest)
	checkResponseStatus(t, responseRecorder, http.StatusOK)

	UserHandler(responseRecorder, invalidMethodRequest)
	checkResponseStatus(t, responseRecorder, http.StatusMethodNotAllowed)
}

func checkResponseStatus(t *testing.T, recorder *httptest.ResponseRecorder, expectedStatus int) {
	if recorder.Result().StatusCode != expectedStatus {
		t.Errorf("Expected status code %d, but got %d", expectedStatus, recorder.Result().StatusCode)
	}
}
