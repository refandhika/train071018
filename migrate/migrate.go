package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "test:test@/?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query1 := "CREATE TABLE `taxdata` (" +
		"`id` INT(64) NOT NULL AUTO_INCREMENT," +
		"`name` VARCHAR(255) NOT NULL," +
		"`code` INT(10) NOT NULL," +
		"`price` FLOAT(64,10) NOT NULL," +
		"`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		"`deleted_at` TIMESTAMP NULL," +
		"PRIMARY KEY (`id`)" +
		")"

	_, err = db.Exec("CREATE DATABASE train071018")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE train071018")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(query1)
	if err != nil {
		panic(err)
	}
}
