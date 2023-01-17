package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/entity"
	"database/sql"
	_ "fmt"
	_ "log"
	"time"
)

func GetAllUser(db *sql.DB) []entity.User {
	var results = []entity.User{}
	sql := `SELECT * FROM users`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = entity.User{}

		err = rows.Scan(user.UserId, user.UserFullname, user.UserUsername, user.UserPassword, user.UserEmail)
		if err != nil {
			panic(err)
		}

		results = append(results, user)
	}

	return results
}

func InsertUser(db *sql.DB, user entity.User) (err error) {
	sql := `INSERT INTO users (user_id, user_fullname, user_username, user_password, user_email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	errs := db.QueryRow(sql, user.UserId, user.UserFullname, user.UserUsername, user.UserPassword, user.UserEmail, time.Now(), time.Now())

	return errs.Err()
}

func UpdateUser(db *sql.DB, user entity.User) (err error) {
	sql := `UPDATE users SET user_fullname = $1, user_username = $2, user_password = $3, user_email = $4, updated_at = $5 WHERE user_id = $6`

	errs := db.QueryRow(sql, user.UserFullname, user.UserUsername, user.UserPassword, user.UserEmail, time.Now(), user.UserId)

	return errs.Err()
}

func DeleteUser(db *sql.DB, user entity.User) (err error) {
	sql := `DELETE FROM users WHERE user_id = $1`

	errs := db.QueryRow(sql, user.UserId)

	return errs.Err()
}
