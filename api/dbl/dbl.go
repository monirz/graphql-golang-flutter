package dbl

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_USER     = "root"
	DB_PASSWORD = "password"
	DB_NAME     = "gqldemo"
)

var once sync.Once

func Connect() (*sql.DB, error) {
	var db *sql.DB
	var err error

	once.Do(func() {
		db, err = sql.Open("mysql", "root:"+DB_PASSWORD+"@/gqldemo?parseTime=true")
		if err != nil {
			log.Fatal(err, "here")
		}

	})
	log.Println("DB err", err)
	return db, err
}

func LogAndQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Println("query: ", query, args)

	r, err := db.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	return r, err
}

func MustExec(db *sql.DB, query string, args ...interface{}) {
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
}
