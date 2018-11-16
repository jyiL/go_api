package database

import (
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"fmt"
)

var PostgresDb *gorm.DB

func init() {
	var err error
    PostgresDb, _ = gorm.Open("postgres", "host=localhost user=postgres dbname=jhpm port=5432 sslmode=disable password=12345678")

    if err != nil {
		fmt.Printf("postgres connect error %v", err)
		log.Fatal(err.Error())
	}

	if PostgresDb.Error != nil {
		fmt.Printf("database error %v", PostgresDb.Error)
		log.Fatal(PostgresDb.Error)
	}
	
	fmt.Printf("postgres %v", PostgresDb)
}