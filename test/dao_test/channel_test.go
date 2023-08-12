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

func TestChannelDao_ListByRoomKey(t *testing.T) {
	testCases := []struct {
		name           string
		mockRows       *sqlmock.Rows
		mockError      error
		expectedChats  *model.Channels
		expectedError  error
	}{
		{
			name: "Successful: Channels found",
			mockRows: sqlmock.NewRows([]string{"id", "channel_key",  "room_key", "name",  "explanation", "type", "created_at", "updated_at"}).
				AddRow(1, "test_key", "test_key", "test_name", "test_explanation", "text", time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
			mockError: nil,
			expectedChats: &model.Channels{
                {
                    ID:          1,
                    ChannelKey:  "test_key",
                    RoomKey:     "test_key",
                    Name:        "test_name",
                    Explanation: "test_explanation",
                    Type:        "text",
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
			repo := dao.NewChannelDao(gormDB)
			mock.ExpectQuery("SELECT").WithArgs("test_key").WillReturnRows(tc.mockRows).WillReturnError(tc.mockError)

			channel, err := repo.ListByRoomKey("test_key")
			assert.Equal(t, tc.expectedChats, channel)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

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

func TestChannelDao_DeleteByChannelKey(t *testing.T) {
    testCases := []struct {
        name             string
        channelKey       string
        mockRowsAffected int64
        mockError        error
        expectedError    error
    }{
        {
            name:            "Successful",
            channelKey:      "test_key",
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
            repo := dao.NewChannelDao(gormDB)
            mock.ExpectBegin()
            mock.ExpectExec("DELETE").WithArgs(tc.channelKey).
                WillReturnResult(sqlmock.NewResult(tc.mockRowsAffected, tc.mockRowsAffected)).
                WillReturnError(tc.mockError)
            mock.ExpectCommit()

            err := repo.DeleteByChannelKey(tc.channelKey, gormDB)
            assert.Equal(t, tc.expectedError, err)
        })
    }
}
