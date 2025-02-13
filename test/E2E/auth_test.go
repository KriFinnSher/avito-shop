package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080"

func TestAuth(t *testing.T) {
	authRequestBody := `{
		"username": "testuser",
		"password": "testpassword"
	}`

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/auth", baseURL), bytes.NewBufferString(authRequestBody))
	if err != nil {
		t.Fatalf("error creating auth request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	authResponse, err := client.Do(req)
	if err != nil {
		t.Fatalf("error making auth request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("error closing body: %v", err)
		}
	}(authResponse.Body)

	assert.Equal(t, http.StatusOK, authResponse.StatusCode)

	authBody, err := io.ReadAll(authResponse.Body)
	if err != nil {
		t.Fatalf("error reading auth response body: %v", err)
	}

	var authData map[string]interface{}
	err = json.Unmarshal(authBody, &authData)
	if err != nil {
		t.Fatalf("error unmarshaling auth response: %v", err)
	}

	_, ok := authData["token"].(string)
	if !ok {
		t.Fatalf("error parsing token from auth response")
	}
}

func getAuthToken(t *testing.T) string {
	authRequestBody := `{
		"username": "testuser",
		"password": "testpassword"
	}`

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/auth", baseURL), bytes.NewBufferString(authRequestBody))
	if err != nil {
		t.Fatalf("error creating auth request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	authResponse, err := client.Do(req)
	if err != nil {
		t.Fatalf("error making auth request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("error closing body: %v", err)
		}
	}(authResponse.Body)

	assert.Equal(t, http.StatusOK, authResponse.StatusCode)

	authBody, err := io.ReadAll(authResponse.Body)
	if err != nil {
		t.Fatalf("error reading auth response body: %v", err)
	}

	var authData map[string]interface{}
	err = json.Unmarshal(authBody, &authData)
	if err != nil {
		t.Fatalf("error unmarshaling auth response: %v", err)
	}

	token, ok := authData["token"].(string)
	if !ok {
		t.Fatalf("error parsing token from auth response")
	}
	return token
}
