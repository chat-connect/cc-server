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

func TestRoomRepository_Insert(t *testing.T) {
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
                UserKey:     "test_key",
                Name:        "test",
				Explanation: "test",
				ImagePath:   "/test",
				UserCount:   0,
				Status:      "public",
            },
            mockRowsAffected: 1,
            mockLastInsertID: 1,
            mockError:        nil,
            expectedRoom: &model.Room{
                ID:          1,
                UserKey:     "test_key",
                Name:        "test",
				Explanation: "test",
				ImagePath:   "/test",
				UserCount:   0,
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
            repo := dao.NewRoomRepository(gormDB)
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
