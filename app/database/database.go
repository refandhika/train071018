package database

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

func Init() *sql.DB {
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
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(5)

	return db
}
