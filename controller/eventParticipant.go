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

func GeteventParticipantById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("event_participant_id"))
	eventParticipantResponse, err := repository.GeteventParticipantById(database.DbConnection, id)

	if err != nil {
		result := helper.BuildResponse(false, "Get Data Event Participant By Id Failed", err)
		c.JSON(http.StatusOK, result)
	} else {
		result := helper.BuildResponse(true, "Get Data Event Participant By Id Success", gin.H{
			"data":   eventParticipantResponse,
			"ticket": "https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=" + eventParticipantResponse.EventParticipantName + eventParticipantResponse.EventParticipantEmail + eventParticipantResponse.EventParticipantCommunity + eventParticipantResponse.EventParticipantStatus,
		})
		c.JSON(http.StatusOK, result)
	}
}

func GetAllEventParticipant(c *gin.Context) {
	if c.Request.Method == "GET" {

		eventParticipants, err := repository.GetAllEventParticipant(database.DbConnection)
		if err != nil {
			result := helper.BuildResponse(false, "Get Data Event Participant Failed", err)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Get Data Event Participant Success", eventParticipants)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func InsertEventParticipant(c *gin.Context) {
	if c.Request.Method == "POST" {
		var eventParticipant entity.EventParticipant

		err := c.ShouldBindJSON(&eventParticipant)
		if err != nil {
			panic(err)
		}

		err = repository.InsertEventParticipant(database.DbConnection, eventParticipant)
		if err != nil {
			result := helper.BuildResponse(false, "Create Data Event Participant Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Create Data Event Participant Success", eventParticipant)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func UpdateEventParticipant(c *gin.Context) {
	if c.Request.Method == "PUT" {
		var eventParticipant entity.EventParticipant

		id, _ := strconv.Atoi(c.Param("event_participant_id"))

		err := c.ShouldBindJSON(&eventParticipant)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		eventParticipant.EventParticipantId = id

		err = repository.UpdateEventParticipant(database.DbConnection, eventParticipant)

		if err != nil {
			result := helper.BuildResponse(false, "Update Data Event Participant Failed", err.Error())
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Update Data Event Participant Success", eventParticipant)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func DeleteEventParticipant(c *gin.Context) {
	if c.Request.Method == "DELETE" {
		var eventParticipant entity.EventParticipant

		id, _ := strconv.Atoi(c.Param("event_participant_id"))

		eventParticipant.EventParticipantId = id

		err := repository.DeleteEventParticipant(database.DbConnection, eventParticipant)

		if err != nil {
			result := helper.BuildResponse(false, "Delete Data Event Participant Failed", err.Error())
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Delete Data Event Participant Success", nil)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}
