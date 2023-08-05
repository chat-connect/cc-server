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

func TestChatDao_ListByRoomKey(t *testing.T) {
	testCases := []struct {
		name           string
		mockRows       *sqlmock.Rows
		mockError      error
		expectedChats  *model.Chats
		expectedError  error
	}{
		{
			name: "Successful: Chats found",
			mockRows: sqlmock.NewRows([]string{"id", "chat_key",  "channel_key", "user_key",  "user_name", "content", "created_at", "updated_at"}).
				AddRow(1, "test_key", "test_key", "test_key", "test_user", "test_content", time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
			mockError: nil,
			expectedChats: &model.Chats{
                {
                    ID:         1,
                    ChatKey:    "test_key",
                    ChannelKey: "test_key",
                    UserKey:    "test_key",
                    UserName:   "test_user",
                    Content:    "test_content",
                    CreatedAt:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
                    UpdatedAt:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),                    
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
			repo := dao.NewChatDao(gormDB)
			mock.ExpectQuery("SELECT").WithArgs("test_key").WillReturnRows(tc.mockRows).WillReturnError(tc.mockError)

			chat, err := repo.ListByRoomKey("test_key")
			assert.Equal(t, tc.expectedChats, chat)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestChatDao_Insert(t *testing.T) {
    testCases := []struct {
        name             string
        mockParam        *model.Chat
        mockRowsAffected int64
        mockLastInsertID int64
        mockError        error
        expectedChat     *model.Chat
        expectedError    error
    }{
        {
            name: "Successful",
            mockParam: &model.Chat{
				ChatKey:    "test_key",
                ChannelKey: "test_key",
                UserKey:    "test_key",
                UserName:   "test_name",
				Content:    "content",
            },
            mockRowsAffected: 1,
            mockLastInsertID: 1,
            mockError:        nil,
            expectedChat: &model.Chat{
                ID:          1,
				ChatKey:    "test_key",
                ChannelKey: "test_key",
                UserKey:    "test_key",
                UserName:   "test_name",
				Content:    "content",
				CreatedAt:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
            },
            expectedError: nil,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            db, mock, _ := sqlmock.New()
            defer db.Close()

            gormDB, _ := gorm.Open("mysql", db)
            repo := dao.NewChatDao(gormDB)
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

            assert.Equal(t, tc.expectedChat, user)
        })
    }
}
