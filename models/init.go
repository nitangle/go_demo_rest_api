package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	fmt.Print("here!!!\n")
	//capture connection properties
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		DBName: "book_manager",
	}
	//Get a database handle.
	var err error
	db, err = sql.Open("mysql",cfg.FormatDSN())
	if err!= nil{
		fmt.Print("inside1\n")
		log.Fatal(err)
		os.Exit(1)
	}
	pingErr := db.Ping()
	if pingErr != nil{
		fmt.Print("inside2\n")
		log.Fatal(pingErr)
		os.Exit(1)
	}
}

func GetDB() *sql.DB{
	return db
}