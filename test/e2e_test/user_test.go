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

	"github.com/game-connect/gc-server/api/presentation/parameter"
	"github.com/game-connect/gc-server/api/presentation/output"
	"github.com/game-connect/gc-server/api/presentation/response"
)

func TestUserE2E_Register(t *testing.T) {
	files := []File{
		"sql/user/user_table.sql",
	}

	db := LoadTestSql(files...)
	defer ClearTestSql(db)

	testCases := []struct {
		name         string
		body         *parameter.RegisterUser
		expectedCode int
		expectedKey  string
	}{
		{
			name: "Successful: User Register",
			body: &parameter.RegisterUser{
				Name:     "test",
				Email:    "test@example.com",
				Password: "test_password",
			},
			expectedCode: http.StatusOK,
			expectedKey:  "pRxN4QA9bt4p",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.body)
			if err != nil {
				t.Fatalf("JSON encoding error: %v", err)
				return
			}

			req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/register_user", os.Getenv("TEST_API_URL")), bytes.NewBuffer(jsonData))
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
				actual := &response.Success{
					Items: &output.RegisterUser{},
				}
				expect := &response.Success{
					Types: "user_register",
					Status: 200,
					Items: &output.RegisterUser{
						UserKey: "pRxN4QA9bt4p",
						Name:    "test",
						Email:   "test@example.com",
						Message: "user register completed",						
					},
				}

				err = json.NewDecoder(resp.Body).Decode(actual)
				if err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}

				if userRegister, ok := actual.Items.(*output.RegisterUser); ok {
					userRegister.UserKey = "pRxN4QA9bt4p"
				}

				assert.Equal(t, expect, actual)
			}
		})
	}
}

func TestUserE2E_Login(t *testing.T) {
	files := []File{
		"sql/user/user_table.sql",
		"sql/user/user_insert.sql",
	}

	db := LoadTestSql(files...)
	defer ClearTestSql(db)

	testCases := []struct {
		name         string
		body         *parameter.LoginUser
		expectedCode int
	}{
		{
			name: "Successful: User Register",
			body: &parameter.LoginUser{
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
				t.Fatalf("JSON encoding error: %v", err)
			}

			req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/login_user", os.Getenv("TEST_API_URL")), bytes.NewBuffer(jsonData))
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
				actual := &response.Success{
					Items: &output.LoginUser{},
				}
				expect := &response.Success{
					Types: "user_login",
					Status: 200,
					Items: &output.LoginUser{
						UserKey: "pRxN4QA9bt4p",
						Name:    "test",
						Email:   "test@example.com",
						Token:   "test",
						Message: "user login completed",
					},
				}

				err = json.NewDecoder(resp.Body).Decode(actual)
				if err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}

				if userLogin, ok := actual.Items.(*output.LoginUser); ok {
					userLogin.Token = "test"
				}
				
				assert.Equal(t, expect, actual)
			}
		})
	}
}

func TestUserE2E_Check(t *testing.T) {
	files := []File{
		"sql/user/user_table.sql",
		"sql/user/user_insert.sql",
	}

	db := LoadTestSql(files...)
	defer ClearTestSql(db)

	testCases := []struct {
		name         string
		userKey      string
		email        string
		password     string
		expectedCode int
	}{
		{
			name:         "Successful: User Check",
			userKey:      "pRxN4QA9bt4p",
			email:        "test@example.com",
			password:     "test",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/check_user/%s", os.Getenv("TEST_API_URL"), tc.userKey), nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			token := AuthUserLogin(tc.email, tc.password)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
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
				actual := &response.Success{
					Items: &output.CheckUser{},
				}
				expect := &response.Success{
					Types: "user_check",
					Status: 200,
					Items: &output.CheckUser{
						UserKey: tc.userKey,
						Name:    "test",
						Email:   "test@example.com",
						Message: "user check completed",
					},
				}

				err = json.NewDecoder(resp.Body).Decode(actual)
				if err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}
				
				assert.Equal(t, expect, actual)
			}
		})
	}
}

func TestUserE2E_Delete(t *testing.T) {
	files := []File{
		"sql/user/user_table.sql",
		"sql/user/user_insert.sql",
	}

	db := LoadTestSql(files...)
	defer ClearTestSql(db)

	testCases := []struct {
		name         string
		userKey      string
		email        string
		password     string
		expectedCode int
	}{
		{
			name:         "Successful: User Delete",
			userKey:      "pRxN4QA9bt4p",
			email:        "test@example.com",
			password:     "test",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/auth/delete_user/%s", os.Getenv("TEST_API_URL"), tc.userKey), nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			token := AuthUserLogin(tc.email, tc.password)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
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
				actual := &response.Success{
					Items: &output.DeleteUser{},
				}
				expect := &response.Success{
					Types: "user_delete",
					Status: 200,
					Items: &output.DeleteUser{
						Message:  "user delete completed",					
					},
				}

				err = json.NewDecoder(resp.Body).Decode(actual)
				if err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}
				
				assert.Equal(t, expect, actual)
			}
		})
	}
}
