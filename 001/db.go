package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connectToDB() *sql.DB {

	
	db, err := sql.Open("mysql", "soham123:soham123@tcp(127.0.0.1:3306)/go_dns")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}else{
		log.Printf("Connected to DB")
	}
	_,err = db.Exec("USE "+"go_dns")
	if err != nil {
		panic(err)
	}
	return db

}

func insertIntoDB(db *sql.DB, name string, ip string) {
	_, err := db.Exec("INSERT INTO user (id,name, ip) VALUES (?, ?, ?)", 1,name, ip)
	if err != nil {
		log.Fatal(err)
	}
}