package main

import "testing"

func TestReadUserInputValid(t *testing.T) {
	_, err := readUserInput(LoadTestArgs{requestsNumber: 10, workersNumber: 5, targetURL: "https://www.google.com"})

	if err != nil {
		t.Error("Expected valid user input")
	}
}

func TestReadUserInputEmptyURL(t *testing.T) {
	_, err := readUserInput(LoadTestArgs{requestsNumber: 10, workersNumber: 5, targetURL: ""})

	if err == nil {
		t.Error("Expected no URL supplied error")
	}
}
