package e2e_test

import (
	"testing"
	
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/chat-connect/cc-server/config/database"
)

type Model interface{}

func SetupTestDatabase(t *testing.T, models ...Model) (db *gorm.DB) {
	db = database.NewDB()
	for _, model := range models {
		err := db.AutoMigrate(model).Error
		if err != nil {
			t.Fatalf("Failed to set up table: %v", err)
		}
	}

	return db
}

func TeardownTestDatabase(t *testing.T, db *gorm.DB, models ...Model) {
	for _, model := range models {
		err := db.DropTableIfExists(model).Error
		if err != nil {
			t.Fatalf("Failed to drop table: %v", err)
		}
	}
	db.Close()
}
