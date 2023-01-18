package repository

import (
	"FP-Sanbercode-Go-Batch-41/entity"
	"database/sql"
	"errors"
	_ "fmt"
	_ "log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserById(db *sql.DB, userId int) (userResponse entity.User, err error) {
	sqlStatement := "SELECT * FROM users WHERE user_id = $1"

	err = db.QueryRow(sqlStatement, userId).Scan(&userResponse.UserId, &userResponse.UserFullname, &userResponse.UserUsername, &userResponse.UserPassword, &userResponse.UserEmail, &userResponse.CreatedAt, &userResponse.UpdatedAt)

	return
}

func GetUserByEmail(db *sql.DB, userEmail string) (userResponse entity.User, err error) {
	sql := "SELECT user_id, user_fullname, user_username, user_password, user_email FROM users WHERE user_email = $1"

	err = db.QueryRow(sql, userEmail).Scan(&userResponse.UserId, &userResponse.UserFullname, &userResponse.UserUsername, &userResponse.UserPassword, &userResponse.UserEmail)

	return
}

func GetAllUser(db *sql.DB) (results []entity.User, err error) {
	sqlStatement := `SELECT * FROM users`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var user = entity.User{}

		err = rows.Scan(&user.UserId, &user.UserFullname, &user.UserUsername, &user.UserPassword, &user.UserEmail, &user.CreatedAt, &user.UpdatedAt)

		results = append(results, user)
	}

	return
}

func InsertUser(db *sql.DB, user entity.User) (err error) {
	sqlStatement := `INSERT INTO users (user_fullname, user_username, user_password, user_email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.UserPassword), 8)
	if err != nil {
		return err
	}

	user.UserPassword = string(bytes)

	errs := db.QueryRow(sqlStatement, &user.UserFullname, &user.UserUsername, &user.UserPassword, &user.UserEmail, time.Now(), time.Now())

	return errs.Err()
}

func UpdateUser(db *sql.DB, user entity.User) (err error) {
	sqlStatement := `UPDATE users SET user_fullname = $2, user_username = $3, user_email = $4, updated_at = $5 WHERE user_id = $1`

	res, err := db.Exec(sqlStatement, user.UserId, user.UserFullname, user.UserUsername, user.UserEmail, time.Now())

	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.New("rows is empty")
	}

	return
}

func DeleteUser(db *sql.DB, user entity.User) (err error) {
	sqlStatement := `DELETE FROM users WHERE user_id = $1`

	res, err := db.Exec(sqlStatement, user.UserId)

	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.New("rows is empty")
	}

	return
}
