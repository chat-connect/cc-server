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

func TestUserRepository_FindByEmail(t *testing.T) {
	testCases := []struct {
		name           string
		mockRows       *sqlmock.Rows
		mockError      error
		expectedUser   *model.User
		expectedError  error
	}{
		{
			name: "Successful: User found",
			mockRows: sqlmock.NewRows([]string{"id", "user_key", "name", "email", "password", "token", "status", "created_at", "updated_at"}).
				AddRow(1, "test_key", "test", "test@example.com", "test_password", "test_token", "login",
					time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
			mockError: nil,
			expectedUser: &model.User{
				ID:        1,
				UserKey:   "test_key",
				Name:      "test",
				Email:     "test@example.com",
				Password:  "test_password",
				Token:     "test_token",
				Status:    "login",
				CreatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError: nil,
		},
		{
			name: "Successful: User not found",
			mockRows: sqlmock.NewRows([]string{"id", "user_key", "name", "email", "password", "token", "status", "created_at", "updated_at"}),
			mockError: nil,
			expectedUser: &model.User{
				ID:        0,
				UserKey:   "",
				Name:      "",
				Email:     "",
				Password:  "",
				Token:     "",
				Status:    "",
				CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError: gorm.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			gormDB, _ := gorm.Open("mysql", db)
			repo := dao.NewUserRepository(gormDB)
			mock.ExpectQuery("SELECT").WithArgs("test@example.com").WillReturnRows(tc.mockRows).WillReturnError(tc.mockError)

			// run test
			user, err := repo.FindByEmail("test@example.com")
			assert.Equal(t, tc.expectedUser, user)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestUserRepository_FindByUserKey(t *testing.T) {
	testCases := []struct {
		name           string
		mockRows       *sqlmock.Rows
		mockError      error
		expectedUser   *model.User
		expectedError  error
	}{
		{
			name: "Successful: User found",
			mockRows: sqlmock.NewRows([]string{"id", "user_key", "name", "email", "password", "token", "status", "created_at", "updated_at"}).
				AddRow(1, "test_key", "test", "test@example.com", "test_password", "test_token", "login",
					time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)),
			mockError: nil,
			expectedUser: &model.User{
				ID:        1,
				UserKey:   "test_key",
				Name:      "test",
				Email:     "test@example.com",
				Password:  "test_password",
				Token:     "test_token",
				Status:    "login",
				CreatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError: nil,
		},
		{
			name: "Successful: User not found",
			mockRows: sqlmock.NewRows([]string{"id", "user_key", "name", "email", "password", "token", "status", "created_at", "updated_at"}),
			mockError: nil,
			expectedUser: &model.User{
				ID:        0,
				UserKey:   "",
				Name:      "",
				Email:     "",
				Password:  "",
				Token:     "",
				Status:    "",
				CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError: gorm.ErrRecordNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			gormDB, _ := gorm.Open("mysql", db)
			repo := dao.NewUserRepository(gormDB)
			mock.ExpectQuery("SELECT").WithArgs("test_key").WillReturnRows(tc.mockRows).WillReturnError(tc.mockError)

			user, err := repo.FindByUserKey("test_key")
			assert.Equal(t, tc.expectedUser, user)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestUserRepository_CountByStatus(t *testing.T) {
	testCases := []struct {
		name          string
		mockResult    int64
		mockError     error
		status        string
		expectedCount int64
		expectedError error
	}{
		{
			name: "Successful",
			mockResult:    5,
			mockError:     nil,
			status:        "login",
			expectedCount: 5,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			gormDB, _ := gorm.Open("mysql", db)
			repo := dao.NewUserRepository(gormDB)
			mock.ExpectQuery("SELECT").WithArgs(tc.status).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(tc.mockResult)).WillReturnError(tc.mockError)

			count, err := repo.CountByStatus(tc.status)
			assert.Equal(t, tc.expectedCount, count)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestUserRepository_Insert(t *testing.T) {
    testCases := []struct {
        name            string
        mockParam       *model.User
        mockRowsAffected int64
        mockLastInsertID int64
        mockError       error
        expectedUser    *model.User
        expectedError   error
    }{
        {
            name: "Successful",
            mockParam: &model.User{
                UserKey:  "test_key",
                Name:     "test",
                Email:    "test@example.com",
                Password: "test_password",
                Token:    "test_token",
                Status:   "login",
            },
            mockRowsAffected: 1,
            mockLastInsertID: 1,
            mockError:        nil,
            expectedUser: &model.User{
                ID:        1,
                UserKey:   "test_key",
                Name:      "test",
                Email:     "test@example.com",
                Password:  "test_password",
                Token:     "test_token",
                Status:    "login",
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
            repo := dao.NewUserRepository(gormDB)
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

            assert.Equal(t, tc.expectedUser, user)
        })
    }
}

func TestUserRepository_Update(t *testing.T) {
    testCases := []struct {
        name            string
        mockParam       *model.User
        mockRowsAffected int64
        mockLastUpdateID int64
        mockError       error
        expectedUser    *model.User
        expectedError   error
    }{
        {
            name: "Successful",
            mockParam: &model.User{
                UserKey:  "test_key",
                Name:     "test",
                Email:    "test@example.com",
                Password: "test_password",
                Token:    "test_token",
                Status:   "login",
            },
            mockRowsAffected: 1,
            mockLastUpdateID: 1,
            mockError:        nil,
            expectedUser: &model.User{
                ID:        0,
                UserKey:   "test_key",
                Name:      "test",
                Email:     "test@example.com",
                Password:  "test_password",
                Token:     "test_token",
                Status:    "login",
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
            repo := dao.NewUserRepository(gormDB)
            mock.ExpectBegin()
			mock.ExpectExec("UPDATE").
                WillReturnResult(sqlmock.NewResult(tc.mockLastUpdateID, tc.mockRowsAffected)).
				WillReturnError(tc.mockError)
            mock.ExpectCommit()

            user, err := repo.Update(tc.mockParam, gormDB)
            if tc.expectedError != nil {
                assert.EqualError(t, err, tc.expectedError.Error())
            } else {
                assert.NoError(t, err)
            }

			user.ID = 0
            user.CreatedAt = time.Time{}
            user.UpdatedAt = time.Time{}

            assert.Equal(t, tc.expectedUser, user)
        })
    }
}

func TestUserRepository_DeleteByUserKey(t *testing.T) {
    testCases := []struct {
        name            string
        userKey         string
        mockRowsAffected int64
        mockError       error
        expectedError   error
    }{
        {
            name:            "Successful",
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
            repo := dao.NewUserRepository(gormDB)
            mock.ExpectBegin()
            mock.ExpectExec("DELETE").WithArgs(tc.userKey).
                WillReturnResult(sqlmock.NewResult(tc.mockRowsAffected, tc.mockRowsAffected)).
                WillReturnError(tc.mockError)
            mock.ExpectCommit()

            err := repo.DeleteByUserKey(tc.userKey, gormDB)
            assert.Equal(t, tc.expectedError, err)
        })
    }
}
