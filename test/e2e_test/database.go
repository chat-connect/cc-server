package e2e_test

import (
	"io/ioutil"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/chat-connect/cc-server/config/database"
)

type Model interface{}

type File interface{}

func SetupTestDatabase(models ...Model) (db *gorm.DB) {
	db = database.NewGormDB()
	for _, model := range models {
		err := db.AutoMigrate(model).Error
		if err != nil {
			panic(err)
		}
	}

	return db
}

func TeardownTestDatabase(db *gorm.DB, models ...Model) {
	for _, model := range models {
		err := db.DropTableIfExists(model).Error
		if err != nil {
			panic(err)
		}
	}

	db.Close()
}

func LoadTestData(files ...File) {
	db := database.NewSqlDB()
    defer db.Close()
	
	for _, file := range files {
		query, err := ioutil.ReadFile(file.(string))
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(string(query))
	}
}
