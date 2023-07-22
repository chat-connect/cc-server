package e2e_test

import (
	"fmt"
	"os"
	"io/ioutil"
	"database/sql"
)

type File interface{}

func NewTestDB() (db *sql.DB) {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	name := os.Getenv("MYSQL_DATABASE")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, name))
	if err != nil {
		panic(err)
	}

	return db
}

func LoadTestSql(files ...File) (db *sql.DB) {
	db = NewTestDB()
	
	for _, file := range files {
		query, err := ioutil.ReadFile(file.(string))
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(string(query))
	}

	return db
}

func ClearTestSql(db *sql.DB) {
    _, err := db.Exec("SHOW TABLES")
    if err != nil {
        panic(err.Error())
    }

    rows, err := db.Query("SHOW TABLES")
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    var tableName string
    for rows.Next() {
        if err := rows.Scan(&tableName); err != nil {
            panic(err.Error())
        }
        _, err = db.Exec("DROP TABLE " + tableName)
        if err != nil {
            panic(err.Error())
        }
    }

    db.Close()
}
