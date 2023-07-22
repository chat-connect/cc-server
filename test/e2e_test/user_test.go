package e2e_test

import (
	"fmt"
	"os"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"

	"github.com/chat-connect/cc-server/api/presentation/parameter"
	"github.com/chat-connect/cc-server/api/presentation/output"
	"github.com/chat-connect/cc-server/domain/model"
)

func TestUserRegister(t *testing.T) {
	models := []Model{
		&model.User{},
	}

	db := SetupTestDatabase(models...)
	defer TeardownTestDatabase(db, models...)

	testCases := []struct {
		name         string
		body         *parameter.UserRegister
		expectedCode int
		expectedKey  string
	}{
		{
			name: "Successful User Register",
			body: &parameter.UserRegister{
				Username:  "test",
				Email:     "test@example.com",
				Password:  "test_password",
			},
			expectedCode: http.StatusOK,
			expectedKey:  "pRxN4QA9bt4p",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.body)
			if err != nil {
				fmt.Println("JSON encoding error:", err)
				return
			}

			req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/user_register", os.Getenv("TEST_API_URL")), bytes.NewBuffer(jsonData))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedCode {
				t.Fatalf("Expected status code %v, but got %v", tc.expectedCode, resp.StatusCode)
			}

			if tc.expectedCode == http.StatusOK {
				actual := &output.UserRegister{}
				expect := &output.UserRegister{
					UserKey:  "pRxN4QA9bt4p",
					Username: "test",
					Email:    "test@example.com",
					Message:  "user register completed",
				}

				err = json.NewDecoder(resp.Body).Decode(actual)
				if err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}

				actual.UserKey = "pRxN4QA9bt4p"
				
				assert.Equal(t, expect, actual)
			}
		})
	}
}

func TestUserLogin(t *testing.T) {
	models := []Model{
		&model.User{},
	}

	files := []File{
		"sql/user/user.sql",
	}

	db := SetupTestDatabase(models...)
	LoadTestData(files...)
	defer TeardownTestDatabase(db, models...)

	testCases := []struct {
		name         string
		body         *parameter.UserLogin
		expectedCode int
	}{
		{
			name: "Successful User Register",
			body: &parameter.UserLogin{
				Email:     "test@example.com",
				Password:  "test",
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.body)
			if err != nil {
				fmt.Println("JSON encoding error:", err)
				return
			}

			req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/user_login", os.Getenv("TEST_API_URL")), bytes.NewBuffer(jsonData))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedCode {
				t.Fatalf("Expected status code %v, but got %v", tc.expectedCode, resp.StatusCode)
			}

			if tc.expectedCode == http.StatusOK {
				actual := &output.UserLogin{}
				expect := &output.UserLogin{
					UserKey:  "pRxN4QA9bt4p",
					Username: "test",
					Email:    "test@example.com",
					Message:  "user login completed",
				}

				err = json.NewDecoder(resp.Body).Decode(actual)
				if err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}

				expect.Token = actual.Token
				
				assert.Equal(t, expect, actual)
			}
		})
	}
}
