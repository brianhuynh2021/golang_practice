package data

import (
	"database/sql"
	_"github.com/lib/pq"
	"log"
	"fmt"
)


const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	password = "54321"
	dbname = "lucky_number"
)


func SetupDatabase() *sql.DB{
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil{
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	
	return db
}
