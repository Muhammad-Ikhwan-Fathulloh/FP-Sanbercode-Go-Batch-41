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

func GetEventById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("event_id"))
	eventResponse, err := repository.GetEventById(database.DbConnection, id)

	if err != nil {
		result := helper.BuildResponse(false, "Get Data Event By Id Failed", err)
		c.JSON(http.StatusOK, result)
	} else {
		result := helper.BuildResponse(true, "Get Data Event By Id Success", eventResponse)
		c.JSON(http.StatusOK, result)
	}
}

func GetAllEventByCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("event_category_id"))
	eventResponse, err := repository.GetAllEventByCategory(database.DbConnection, id)

	if err != nil {
		result := helper.BuildResponse(false, "Get Data Event By Category Failed", err)
		c.JSON(http.StatusOK, result)
	} else {
		result := helper.BuildResponse(true, "Get Data Event By Category Success", eventResponse)
		c.JSON(http.StatusOK, result)
	}
}

func GetAllEventByCommunity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("event_community_id"))
	eventResponse, err := repository.GetAllEventByCommunity(database.DbConnection, id)

	if err != nil {
		result := helper.BuildResponse(false, "Get Data Event By Community Failed", err)
		c.JSON(http.StatusOK, result)
	} else {
		result := helper.BuildResponse(true, "Get Data Event By Community Success", eventResponse)
		c.JSON(http.StatusOK, result)
	}
}

func GetAllEvent(c *gin.Context) {
	if c.Request.Method == "GET" {

		events, err := repository.GetAllEvent(database.DbConnection)
		if err != nil {
			result := helper.BuildResponse(false, "Get Data Event Failed", err)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Get Data Event Success", events)
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

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func UpdateEvent(c *gin.Context) {
	if c.Request.Method == "PUT" {
		var event entity.Event

		id, _ := strconv.Atoi(c.Param("event_id"))

		err := c.ShouldBindJSON(&event)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		event.EventId = id

		err = repository.UpdateEvent(database.DbConnection, event)

		if err != nil {
			result := helper.BuildResponse(false, "Update Data Event Failed", err.Error())
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Update Data Event Success", event)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func DeleteEvent(c *gin.Context) {
	if c.Request.Method == "DELETE" {
		var event entity.Event

		id, _ := strconv.Atoi(c.Param("event_id"))

		event.EventId = id

		err := repository.DeleteEvent(database.DbConnection, event)

		if err != nil {
			result := helper.BuildResponse(false, "Delete Data Event Failed", err.Error())
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Delete Data Event Success", nil)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}
