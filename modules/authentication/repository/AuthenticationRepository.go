package repository

import (
	"database/sql"
	"tubesppb-be/database"
)

func VerifyLogin(username, password string) (*sql.Rows, error) {
	db := database.Connect()
	defer db.Close()

	return db.Query("SELECT u.uid, u.username, u.profile_picture, u.friend_mode FROM users u WHERE u.username=$1 AND u.password=$2",
		username,
		password,
	)
}
