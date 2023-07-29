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
	"github.com/chat-connect/cc-server/api/presentation/response"
)

func TestRoomE2E_RoomCreate(t *testing.T) {
	files := []File{
		"sql/room/room_table.sql",
		"sql/room/user_table.sql",
		"sql/room/user_insert.sql",
	}

	db := LoadTestSql(files...)
	defer ClearTestSql(db)

	testCases := []struct {
		email        string // 認証用
		password     string // 認証用
		name         string
		userKey      string
		body         *parameter.RoomCreate
		expectedCode int
	}{
		{
			email:    "test@example.com",
			password: "test",
			name:     "Successful: Room Create",
			userKey:  "pRxN4QA9bt4p",
			body: &parameter.RoomCreate{
				Name:        "test",
				Explanation: "test",
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

			req, err := http.NewRequest("POST", fmt.Sprintf("%s/user/%s/room_create", os.Getenv("TEST_API_URL"), tc.userKey), bytes.NewBuffer(jsonData))
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
					Items: &output.RoomCreate{},
				}
				expect := &response.Success{
					Types: "room_create",
					Status: 200,
					Items: &output.RoomCreate{
						RoomKey:     "test",
						UserKey:     tc.userKey,
						Name:        "test",
						Explanation: "test",
						ImagePath:   "",
						UserCount:   1,
						Status:      "public",
						Message:     "room create completed",
					},
				}

				err = json.NewDecoder(resp.Body).Decode(actual)
				if err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}

				if roomCreate, ok := actual.Items.(*output.RoomCreate); ok {
					roomCreate.RoomKey = "test"
				}
				
				assert.Equal(t, expect, actual)
			}
		})
	}
}
