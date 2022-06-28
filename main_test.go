package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"user-management/controllers"
)

func TestGetAllUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	defer req.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetAllUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
