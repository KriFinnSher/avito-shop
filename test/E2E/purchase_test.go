package e2e

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestPurchaseItem(t *testing.T) {
	token := getAuthToken(t)

	itemName := "t-shirt"

	purchaseRequestBody := `{"item": "` + itemName + `"}`

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/buy/%s", baseURL, itemName), bytes.NewBufferString(purchaseRequestBody))
	if err != nil {
		t.Fatalf("Error creating purchase request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	purchaseResponse, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error making purchase request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("error closing body: %v", err)
		}
	}(purchaseResponse.Body)

	assert.Equal(t, http.StatusOK, purchaseResponse.StatusCode)
}
