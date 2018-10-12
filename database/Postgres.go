package database

import (
    "database/sql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var PostgresDb *sql.DB

func init() {
	var err error
    PostgresDb, _ := sql.Open("postgres", "host=10.0.0.119 user=postgres dbname=jhpm port=5432 sslmode=disable password=123456")

    if err != nil {
		log.Fatal(err.Error())
	}
	err = PostgresDb.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}