package dao_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/game-connect/gc-server/infra/dao"
	"github.com/game-connect/gc-server/domain/model"
)

func TestRoomDao_FindByRoomKey(t *testing.T) {
	testCases := []struct {
		name           string
		mockRows       *sqlmock.Rows
		mockError      error
		expectedRoom   *model.Room
		expectedError  error
	}{
		{
			name: "Successful: Room found",
			mockRows: sqlmock.NewRows([]string{"id", "room_key", "name", "description", "image_path", "user_count", "status", "created_at", "updated_at"}).
				AddRow(1, "test_key", "test", "test", "/", 1, "public", time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
			mockError: nil,
			expectedRoom: &model.Room{
				ID:          1,
				RoomKey:     "test_key",
				Name:        "test",
                Description: "test",
                ImagePath:   "/",
                UserCount:   1,
				Status:      "public",
				CreatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError: nil,
		},
		{
			name: "Successful: Room not found",
			mockRows: sqlmock.NewRows([]string{"id", "room_key", "name", "description", "image_path", "user_count", "status", "created_at", "updated_at"}),
			mockError: nil,
			expectedRoom: (*model.Room)(nil),
			expectedError: gorm.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			gormDB, _ := gorm.Open("mysql", db)
			repo := dao.NewRoomDao(gormDB)
			mock.ExpectQuery("SELECT").WithArgs("test_key").WillReturnRows(tc.mockRows).WillReturnError(tc.mockError)

			user, err := repo.FindByRoomKey("test_key")
			assert.Equal(t, tc.expectedRoom, user)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestRoomDao_Insert(t *testing.T) {
    testCases := []struct {
        name             string
        mockParam        *model.Room
        mockRowsAffected int64
        mockLastInsertID int64
        mockError        error
        expectedRoom     *model.Room
        expectedError    error
    }{
        {
            name: "Successful",
            mockParam: &model.Room{
                RoomKey:     "test_key",
                UserKey:     "test_key",
                Name:        "test",
				Description: "test",
				ImagePath:   "/test",
				UserCount:   1,
				Status:      "public",
            },
            mockRowsAffected: 1,
            mockLastInsertID: 1,
            mockError:        nil,
            expectedRoom: &model.Room{
                ID:          1,
                RoomKey:     "test_key",
                UserKey:     "test_key",
                Name:        "test",
				Description: "test",
				ImagePath:   "/test",
				UserCount:   1,
				Status:      "public",
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
            repo := dao.NewRoomDao(gormDB)
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

            assert.Equal(t, tc.expectedRoom, user)
        })
    }
}

func TestRoomDao_DeleteByRoomKey(t *testing.T) {
    testCases := []struct {
        name             string
        roomKey          string
        mockRowsAffected int64
        mockError        error
        expectedError    error
    }{
        {
            name:            "Successful",
            roomKey:         "test_key",
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
            repo := dao.NewRoomDao(gormDB)
            mock.ExpectBegin()
            mock.ExpectExec("DELETE").WithArgs(tc.roomKey).
                WillReturnResult(sqlmock.NewResult(tc.mockRowsAffected, tc.mockRowsAffected)).
                WillReturnError(tc.mockError)
            mock.ExpectCommit()

            err := repo.DeleteByRoomKey(tc.roomKey, gormDB)
            assert.Equal(t, tc.expectedError, err)
        })
    }
}
