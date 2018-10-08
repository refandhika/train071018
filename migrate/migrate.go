package main

import (
	"database/sql"
	"log"
	"net"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetGateway() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Panic(err)
	}
	var ip net.IP
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Panic(err)
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
		}
	}
	split := strings.Split(ip.String(), ".")
	split[len(split)-1] = "1"
	return strings.Join(split, ".")
}

func main() {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+
		":"+os.Getenv("MYSQL_PASS")+
		"@"+os.Getenv("MYSQL_PROTOCOL")+
		"("+GetGateway()+
		":"+os.Getenv("MYSQL_PORT")+
		")/"+os.Getenv("MYSQL_DB")+
		"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(5)
	defer db.Close()

	query1 := "CREATE TABLE IF NOT EXISTS `taxdata` (" +
		"`id` INT(64) NOT NULL AUTO_INCREMENT," +
		"`name` VARCHAR(255) NOT NULL," +
		"`code` INT(10) NOT NULL," +
		"`price` FLOAT(64,10) NOT NULL," +
		"`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		"`deleted_at` TIMESTAMP NULL," +
		"PRIMARY KEY (`id`)" +
		")"

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + os.Getenv("MYSQL_DB"))
	if err != nil {
		log.Panic(err)
	} else {
		_, err = db.Exec("USE " + os.Getenv("MYSQL_DB"))
		if err != nil {
			log.Panic(err)
		} else {
			_, err = db.Exec(query1)
			if err != nil {
				log.Panic(err)
			}
		}
	}

	log.Print("Migrate END!")
	return

}
