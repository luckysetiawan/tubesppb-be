package repository

import (
	"database/sql"
	"tubesppb-be/database"
)

func VerifyLogin(username, password string) *sql.Row {
	db := database.Connect()
	defer db.Close()

	return db.QueryRow("SELECT * FROM users e WHERE e.username=$1 AND e.password=$2",
		username,
		password,
	)
}
