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

}
