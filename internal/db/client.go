package db

import (
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var DB *sqlx.DB

func ConnectDB(dsnDb string) (*sqlx.DB, error) {
	var err error
	DB, err = sqlx.Connect("postgres", dsnDb)
	if err != nil {
		log.Fatalln(err)
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("Databese not connect")
		return DB, err
	}
	return DB, nil
}
