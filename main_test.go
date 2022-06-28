package main

import (
	"net/http"
	"testing"
)

func TestGetAllUsers(t *testing.T) {

	_, err := http.NewRequest("GET", "/users/uses", nil)
	if err != nil {
		t.Fatal(err)
	}

}
