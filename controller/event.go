package controller

import (
	"FP-Sanbercode-Go-Batch-41/database"
	"FP-Sanbercode-Go-Batch-41/entity"
	"FP-Sanbercode-Go-Batch-41/helper"
	"FP-Sanbercode-Go-Batch-41/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllEvent(c *gin.Context) {
	if c.Request.Method == "GET" {

		event := repository.GetAllEvent(database.DbConnection)
		if event != nil {
			result := helper.BuildResponse(false, "Get Data Event Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Get Data Event Success", event)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func InsertEvent(c *gin.Context) {
	if c.Request.Method == "POST" {
		var event entity.Event

		err := c.ShouldBindJSON(&event)
		if err != nil {
			panic(err)
		}

		err = repository.InsertEvent(database.DbConnection, event)
		if err != nil {
			result := helper.BuildResponse(false, "Create Data Event Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Create Data Event Success", event)
			c.JSON(http.StatusOK, result)
		}
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func UpdateEvent(c *gin.Context) {
	if c.Request.Method == "PUT" {
		var event entity.Event

		eventId, _ := strconv.Atoi(c.Param("event_id"))

		err := c.ShouldBindJSON(&event)
		if err != nil {
			panic(err)
		}

		event.EventId = int64(eventId)

		err = repository.UpdateEvent(database.DbConnection, event)

		if err != nil {
			result := helper.BuildResponse(false, "Update Data Event Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Update Data Event Success", event)
			c.JSON(http.StatusOK, result)
		}
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func DeleteEvent(c *gin.Context) {
	if c.Request.Method == "DELETE" {
		var event entity.Event

		eventId, _ := strconv.Atoi(c.Param("event_id"))

		err := c.ShouldBindJSON(&event)
		if err != nil {
			panic(err)
		}

		event.EventId = int64(eventId)

		err = repository.DeleteEvent(database.DbConnection, event)

		if err != nil {
			result := helper.BuildResponse(false, "Delete Data Event Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Delete Data Event Success", event)
			c.JSON(http.StatusOK, result)
		}
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}
