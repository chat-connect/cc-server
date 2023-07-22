package e2e_test

import (
	"fmt"
	"os"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/chat-connect/cc-server/api/presentation/parameter"
	"github.com/chat-connect/cc-server/api/presentation/output"
)

func AuthUserLogin(email string, password string) (token string) {
	jsonData, err := json.Marshal(&parameter.UserLogin{
		Email:     email,
		Password:  password,
	})
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/user_login", os.Getenv("TEST_API_URL")), bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Failed to create request:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
	}
	defer resp.Body.Close()

	actual := &output.UserLogin{}
	err = json.NewDecoder(resp.Body).Decode(actual)
	if err != nil {
		fmt.Println("Failed to parse response:", err)
	}

	token = actual.Token
	
	return token
}