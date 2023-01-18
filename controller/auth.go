package controller

import (
	"FP-Sanbercode-Go-Batch-41/auth"
	"FP-Sanbercode-Go-Batch-41/database"
	"FP-Sanbercode-Go-Batch-41/helper"
	"FP-Sanbercode-Go-Batch-41/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"user_email"`
	Password string `json:"user_password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := repository.GetUserByEmail(database.DbConnection, request.Email)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	credentialError := repository.CheckPasswordHash(request.Password, userResponse.UserPassword)
	if credentialError == false {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(userResponse.UserEmail, userResponse.UserUsername)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	result := helper.BuildResponse(true, "Login User Success", gin.H{
		"fullname": userResponse.UserFullname,
		"username": userResponse.UserUsername,
		"email":    userResponse.UserEmail,
		"token":    tokenString,
	})
	context.JSON(http.StatusOK, result)
}
