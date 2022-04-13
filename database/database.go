package database

import (
	"database/sql"
	"fmt"
	"log"
	"tubesppb-be/config"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {

	// load config
	host := config.LoadConfig("DB_HOST")
	port := config.LoadConfig("DB_PORT")
	user := config.LoadConfig("DB_USER")
	password := config.LoadConfig("DB_PASS")
	name := config.LoadConfig("DB_NAME")

	// attempt connection
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, password, name)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
