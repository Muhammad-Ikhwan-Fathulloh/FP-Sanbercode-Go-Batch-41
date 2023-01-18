package controller

import (
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"Final-Project-Sanbercode-Go-Batch-41/entity"
	"Final-Project-Sanbercode-Go-Batch-41/helper"
	"Final-Project-Sanbercode-Go-Batch-41/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	if c.Request.Method == "GET" {

		users := repository.GetAllUser(database.DbConnection)
		if users == nil {
			result := helper.BuildResponse(false, "Get Data User Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Get Data User Success", users)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func InsertUser(c *gin.Context) {
	if c.Request.Method == "POST" {
		var user entity.User

		err := c.ShouldBindJSON(&user)
		if err != nil {
			panic(err)
		}

		err = repository.InsertUser(database.DbConnection, user)
		if err != nil {
			result := helper.BuildResponse(false, "Create Data User Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Create Data User Success", user)
			c.JSON(http.StatusOK, result)
		}
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func UpdateUser(c *gin.Context) {
	if c.Request.Method == "PUT" {
		var user entity.User

		userId, _ := strconv.Atoi(c.Param("user_id"))

		err := c.ShouldBindJSON(&user)
		if err != nil {
			panic(err)
		}

		user.UserId = int64(userId)

		err = repository.UpdateUser(database.DbConnection, user)

		if err != nil {
			result := helper.BuildResponse(false, "Update Data User Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Update Data User Success", user)
			c.JSON(http.StatusOK, result)
		}
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func DeleteUser(c *gin.Context) {
	if c.Request.Method == "DELETE" {
		var user entity.User

		userId, _ := strconv.Atoi(c.Param("user_id"))

		err := c.ShouldBindJSON(&user)
		if err != nil {
			panic(err)
		}

		user.UserId = int64(userId)

		err = repository.DeleteUser(database.DbConnection, user)

		if err != nil {
			result := helper.BuildResponse(false, "Delete Data User Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Delete Data User Success", user)
			c.JSON(http.StatusOK, result)
		}
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}
