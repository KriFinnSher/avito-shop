package e2e

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestSendCoins(t *testing.T) {
	token := getAuthToken(t)

	sendCoinsRequestBody := `{
		"toUser": "krifinnsher",
		"amount": 50
	}`

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/sendCoin", baseURL), bytes.NewBufferString(sendCoinsRequestBody))
	if err != nil {
		t.Fatalf("error creating sendCoins request: %v", err)
	}
	req.Header.Set("authorization", "Bearer "+token)
	req.Header.Set("content-Type", "application/json")

	client := &http.Client{}
	sendCoinsResponse, err := client.Do(req)
	if err != nil {
		t.Fatalf("error making sendCoins request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("error closing body: %v", err)
		}
	}(sendCoinsResponse.Body)

	assert.Equal(t, http.StatusOK, sendCoinsResponse.StatusCode)
}
