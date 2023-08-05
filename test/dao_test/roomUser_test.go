package dao_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/chat-connect/cc-server/infra/dao"
	"github.com/chat-connect/cc-server/domain/model"
)

func TestRoomDao_ListByUserKey(t *testing.T) {
	testCases := []struct {
		name           string
		mockRows       *sqlmock.Rows
		mockError      error
		expectedRoomUsers  *model.RoomUsers
		expectedError  error
	}{
		{
			name: "Successful: RoomUsers found",
			mockRows: sqlmock.NewRows([]string{"id", "room_user_key",  "room_key", "user_key",  "host", "status", "created_at", "updated_at"}).
				AddRow(1, "test_key", "test_key", "test_key", 0, "online", time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
			mockError: nil,
			expectedRoomUsers: &model.RoomUsers{
                {
                    ID:          1,
					RoomUserKey: "test_key",
                    RoomKey:     "test_key",
                    UserKey:     "test_key",
					Host:        false,
					Status:      "online",
                    CreatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
                    UpdatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),                    
                },
            },
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			gormDB, _ := gorm.Open("mysql", db)
			repo := dao.NewRoomUserDao(gormDB)
			mock.ExpectQuery("SELECT").WithArgs("test_key").WillReturnRows(tc.mockRows).WillReturnError(tc.mockError)

			chat, err := repo.ListByUserKey("test_key")
			assert.Equal(t, tc.expectedRoomUsers, chat)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestRoomUserDao_Insert(t *testing.T) {
    testCases := []struct {
        name             string
        mockParam        *model.RoomUser
        mockRowsAffected int64
        mockLastInsertID int64
        mockError        error
        expectedRoomUser *model.RoomUser
        expectedError    error
    }{
        {
            name: "Successful",
            mockParam: &model.RoomUser{
                RoomUserKey: "test_key",
                RoomKey:     "test_key",
                UserKey:     "test_key",
                Host:        false,
				Status:      "online",
            },
            mockRowsAffected: 1,
            mockLastInsertID: 1,
            mockError:        nil,
            expectedRoomUser: &model.RoomUser{
                ID:          1,
                RoomUserKey: "test_key",
                RoomKey:     "test_key",
                UserKey:     "test_key",
                Host:        false,
				Status:      "online",
				CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
            },
            expectedError: nil,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            db, mock, _ := sqlmock.New()
            defer db.Close()

            gormDB, _ := gorm.Open("mysql", db)
            repo := dao.NewRoomUserDao(gormDB)
            mock.ExpectBegin()
            mock.ExpectExec("INSERT").
                WillReturnResult(sqlmock.NewResult(tc.mockLastInsertID, tc.mockRowsAffected)).
                WillReturnError(tc.mockError)
            mock.ExpectCommit()

            user, err := repo.Insert(tc.mockParam, gormDB)
            if tc.expectedError != nil {
                assert.EqualError(t, err, tc.expectedError.Error())
            } else {
                assert.NoError(t, err)
            }

            user.CreatedAt = time.Time{}
            user.UpdatedAt = time.Time{}

            assert.Equal(t, tc.expectedRoomUser, user)
        })
    }
}

func TestRoomUserDao_DeleteByRoomKeyAndUserKey(t *testing.T) {
    testCases := []struct {
        name            string
        roomKey         string
        userKey         string
        mockRowsAffected int64
        mockError       error
        expectedError   error
    }{
        {
            name:            "Successful",
            roomKey:         "test_key",
            userKey:         "test_key",
            mockRowsAffected: 1,
            mockError:        nil,
            expectedError:    nil,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            db, mock, _ := sqlmock.New()
            defer db.Close()

            gormDB, _ := gorm.Open("mysql", db)
            dao := dao.NewRoomUserDao(gormDB)
            mock.ExpectBegin()
            mock.ExpectExec("DELETE").WithArgs(tc.roomKey, tc.userKey).
                WillReturnResult(sqlmock.NewResult(tc.mockRowsAffected, tc.mockRowsAffected)).
                WillReturnError(tc.mockError)
            mock.ExpectCommit()

            err := dao.DeleteByRoomKeyAndUserKey(tc.roomKey, tc.userKey, gormDB)
            assert.Equal(t, tc.expectedError, err)
        })
    }
}
