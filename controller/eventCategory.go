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

func GetAllEventCategory(c *gin.Context) {
	if c.Request.Method == "GET" {

		eventCategory := repository.GetAllEventCategory(database.DbConnection)
		if eventCategory != nil {
			result := helper.BuildResponse(false, "Get Data Event Category Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Get Data Event Category Success", eventCategory)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func InsertEventCategory(c *gin.Context) {
	if c.Request.Method == "POST" {
		var eventCategory entity.EventCategory

		err := c.ShouldBindJSON(&eventCategory)
		if err != nil {
			panic(err)
		}

		err = repository.InsertEventCategory(database.DbConnection, eventCategory)
		if err != nil {
			result := helper.BuildResponse(false, "Create Data Event Category Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Create Data Event Category Success", eventCategory)
			c.JSON(http.StatusOK, result)
		}
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func UpdateEventCategory(c *gin.Context) {
	if c.Request.Method == "PUT" {
		var eventCategory entity.EventCategory

		eventCategoryId, _ := strconv.Atoi(c.Param("event_category_id"))

		err := c.ShouldBindJSON(&eventCategory)
		if err != nil {
			panic(err)
		}

		eventCategory.EventCategoryId = int64(eventCategoryId)

		err = repository.UpdateEventCategory(database.DbConnection, eventCategory)

		if err != nil {
			result := helper.BuildResponse(false, "Update Data Event Category Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Update Data Event Category Success", eventCategory)
			c.JSON(http.StatusOK, result)
		}
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func DeleteEventCategory(c *gin.Context) {
	if c.Request.Method == "DELETE" {
		var eventCategory entity.EventCategory

		eventCategoryId, _ := strconv.Atoi(c.Param("event_category_id"))

		err := c.ShouldBindJSON(&eventCategory)
		if err != nil {
			panic(err)
		}

		eventCategory.EventCategoryId = int64(eventCategoryId)

		err = repository.DeleteEventCategory(database.DbConnection, eventCategory)

		if err != nil {
			result := helper.BuildResponse(false, "Delete Data Event Category Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Delete Data Event Category Success", eventCategory)
			c.JSON(http.StatusOK, result)
		}
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}
