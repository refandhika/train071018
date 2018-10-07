package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+
		":"+os.Getenv("MYSQL_PASS")+
		"@"+os.Getenv("MYSQL_PROTOCOL")+
		"("+os.Getenv("MYSQL_ADDRESS")+
		":"+os.Getenv("MYSQL_PORT")+
		")/"+os.Getenv("MYSQL_DB")+
		"?charset=utf8&parseTime=True&loc=Local")
	log.Print("mysql", os.Getenv("MYSQL_USER")+
		":"+os.Getenv("MYSQL_PASS")+
		"@"+os.Getenv("MYSQL_PROTOCOL")+
		"("+os.Getenv("MYSQL_ADDRESS")+
		":"+os.Getenv("MYSQL_PORT")+
		")/"+os.Getenv("MYSQL_DB")+
		"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Panic(err)
	}

	return db
}
