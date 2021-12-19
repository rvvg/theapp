package main

import (
	"io"
	"net/http"
	"testing"
)

// TestCalc send http request to the app and check if the response is correct
func TestCalc(t *testing.T) {
	// Send http request to the app
	res, err := http.Get("http://localhost:3000/?n=2")

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	defer res.Body.Close()

	// Check if the response is correct
	if res.StatusCode != 200 {
		t.Errorf("Status code is not 200, got: %d", res.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// Check if the response body is correct
	if string(body) != "4" {
		t.Errorf("Response body is not 4, got: %s", string(body))
	}
}

func TestBlacklisted(t *testing.T) {
	// Send http request to the app
	res, err := http.Get("http://localhost:3000/blacklisted")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	defer res.Body.Close()

	// Check if the response is correct
	if res.StatusCode != 444 {
		t.Errorf("Status code is not 444, got: %d", res.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// Check if the response body is correct
	if string(body) != "Blacklisted" {
		t.Errorf("Response body is not Blacklisted, got: %s", string(body))
	}
}
