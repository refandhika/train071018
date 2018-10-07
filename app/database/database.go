package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	db, err := sql.Open("mysql", "test:test@/train071018?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic(err)
	}

	return db
}
