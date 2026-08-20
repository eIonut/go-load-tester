package main

import "testing"

func TestParseURLValid(t *testing.T) {
	err := parseURL("https://example.com")

	if err != nil {
		t.Error("expected valid URL")
	}
}

func TestParseURL(t *testing.T) {
	err := parseURL("jsonplaceholder.typicode.com/posts")

	if err == nil {
		t.Error("expected URL without scheme to be invalid")
	}
}

func TestParseURLInvalidScheme(t *testing.T) {
	err := parseURL("ftp://example.com")

	if err == nil {
		t.Error("expected error for unsupported scheme")
	}
}
