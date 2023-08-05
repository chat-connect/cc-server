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

func TestChannelkDao_Insert(t *testing.T) {
    testCases := []struct {
        name             string
        mockParam        *model.Channel
        mockRowsAffected int64
        mockLastInsertID int64
        mockError        error
        expectedChannel  *model.Channel
        expectedError    error
    }{
        {
            name: "Successful",
            mockParam: &model.Channel{
				ChannelKey:  "test_key",
                RoomKey:     "test_key",
                Explanation: "test_explanation",
                Type:        "type",
            },
            mockRowsAffected: 1,
            mockLastInsertID: 1,
            mockError:        nil,
            expectedChannel: &model.Channel{
                ID:          1,
				ChannelKey:  "test_key",
                RoomKey:     "test_key",
                Explanation: "test_explanation",
                Type:        "type",
				CreatedAt:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
            },
            expectedError: nil,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            db, mock, _ := sqlmock.New()
            defer db.Close()

            gormDB, _ := gorm.Open("mysql", db)
            repo := dao.NewChannelDao(gormDB)
            mock.ExpectBegin()
            mock.ExpectExec("INSERT").
                WillReturnResult(sqlmock.NewResult(tc.mockLastInsertID, tc.mockRowsAffected)).
                WillReturnError(tc.mockError)
            mock.ExpectCommit()

            channel, err := repo.Insert(tc.mockParam, gormDB)
            if tc.expectedError != nil {
                assert.EqualError(t, err, tc.expectedError.Error())
            } else {
                assert.NoError(t, err)
            }

            channel.CreatedAt = time.Time{}
            channel.UpdatedAt = time.Time{}

            assert.Equal(t, tc.expectedChannel, channel)
        })
    }
}
