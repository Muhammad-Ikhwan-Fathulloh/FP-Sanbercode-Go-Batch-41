package controller

import (
	"Final-Project-Sanbercode-Go-Batch-41/auth"
	"Final-Project-Sanbercode-Go-Batch-41/entity"
	"Final-Project-Sanbercode-Go-Batch-41/repository"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"user_email"`
	Password string `json:"user_password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user entity.User
	var db *sql.DB
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	err := db.QueryRow(
		`SELECT user_id, user_email, user_password FROM users WHERE user_email = $1`, request.Email).Scan(
		&user.UserId, &user.UserEmail, &user.UserPassword,
	)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password.",
		})
		panic(err)
	}

	credentialError := repository.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.UserEmail, user.UserUsername)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
