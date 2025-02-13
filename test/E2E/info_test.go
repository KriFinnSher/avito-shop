package e2e

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	token := getAuthToken(t)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/info", baseURL), nil)
	if err != nil {
		t.Fatalf("error creating info request: %v", err)
	}
	req.Header.Set("authorization", "Bearer "+token)

	client := &http.Client{}
	infoResponse, err := client.Do(req)
	if err != nil {
		t.Fatalf("error making info request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("error closing body: %v", err)
		}
	}(infoResponse.Body)

	assert.Equal(t, http.StatusOK, infoResponse.StatusCode)
}
