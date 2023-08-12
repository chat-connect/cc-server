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

func TestRoomE2E_RoomCreate(t *testing.T) {
	files := []File{
		"sql/room/room_table.sql",
		"sql/room/room_user_table.sql",
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
		body         *parameter.CreateRoom
		expectedCode int
	}{
		{
			email:    "test@example.com",
			password: "test",
			name:     "Successful: Room Create",
			userKey:  "pRxN4QA9bt4p",
			body: &parameter.CreateRoom{
				Name:        "test",
				Explanation: "test",
				Status:      "public",
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

			req, err := http.NewRequest("POST", fmt.Sprintf("%s/room/%s/create_room", os.Getenv("TEST_API_URL"), tc.userKey), bytes.NewBuffer(jsonData))
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
					Items: &output.CreateRoom{},
				}
				expect := &response.Success{
					Types: "room_create",
					Status: 200,
					Items: &output.CreateRoom{
						RoomKey:     "test",
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

				if roomCreate, ok := actual.Items.(*output.CreateRoom); ok {
					roomCreate.RoomKey = "test"
				}
				
				assert.Equal(t, expect, actual)
			}
		})
	}
}

func TestRoomE2E_RoomJoin(t *testing.T) {
	files := []File{
		"sql/room/room_table.sql",
		"sql/room/room_insert.sql",
		"sql/room/room_user_table.sql",
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
		roomKey      string
		body         *parameter.CreateRoom
		expectedCode int
	}{
		{
			email:    "test@example.com",
			password: "test",
			name:     "Successful: Room Join",
			userKey:  "pRxN4QA9bt4p",
			roomKey:  "pRxN4QA9bt4ppRxN4QA9",
			body: &parameter.CreateRoom{
				Name:        "test",
				Explanation: "test",
				Status:      "public",
			},
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", fmt.Sprintf("%s/room/%s/join_room/%s", os.Getenv("TEST_API_URL"), tc.userKey, tc.roomKey), nil)
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
					Items: &output.JoinRoom{},
				}
				expect := &response.Success{
					Types: "room_join",
					Status: 200,
					Items: &output.JoinRoom{
						RoomUserKey: "test",
						Status:      "online",
						Message:     "room join completed",
					},
				}

				err = json.NewDecoder(resp.Body).Decode(actual)
				if err != nil {
					t.Fatalf("Failed to parse response: %v", err)
				}

				if roomJoin, ok := actual.Items.(*output.JoinRoom); ok {
					roomJoin.RoomUserKey = "test"
				}
				
				assert.Equal(t, expect, actual)
			}
		})
	}
}

func TestRoomE2E_RoomOut(t *testing.T) {
	files := []File{
		"sql/room/room_table.sql",
		"sql/room/room_insert.sql",
		"sql/room/room_user_table.sql",
		"sql/room/user_table.sql",
		"sql/room/user_insert.sql",
	}

	db := LoadTestSql(files...)
	defer ClearTestSql(db)

	testCases := []struct {
		name         string
		email        string // 認証
		password     string // 認証
		roomKey      string
		userKey      string
		expectedCode int
	}{
		{
			name:         "Successful: Room Out",
			email:        "test@example.com",
			password:     "test",
			roomKey:  "pRxN4QA9bt4p",
			userKey:  "pRxN4QA9bt4p",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/room/%s/out_room/%s", os.Getenv("TEST_API_URL"), tc.userKey, tc.roomKey), nil)
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
					Items: &output.OutRoom{},
				}
				expect := &response.Success{
					Types: "room_out",
					Status: 200,
					Items: &output.OutRoom{
						Message:  "room out completed",					
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
