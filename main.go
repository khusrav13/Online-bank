package main

import (
	"SecondProject/db"
	"SecondProject/pkg/core/services"
	"database/sql"
	"log"
)

func main() {
	DB, err := sql.Open("sqlite3", "bank")
	if err != nil {
		log.Fatal("Please try again:", err)
	}
	err = db.Dbinits(DB)
	if err != nil {
		log.Fatal("Cannot create table: ", err)
	}
	Start(DB)
}

func Start(Db *sql.DB) {
	for {
		services.Authorizationss(Db)
	}
}
