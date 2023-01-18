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

func GetCommunityById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("community_id"))
	communityResponse, err := repository.GetCommunityById(database.DbConnection, id)

	if err != nil {
		result := helper.BuildResponse(false, "Get Data Community By Id Failed", err)
		c.JSON(http.StatusOK, result)
	} else {
		result := helper.BuildResponse(true, "Get Data Community By Id Success", communityResponse)
		c.JSON(http.StatusOK, result)
	}
}

func GetAllCommunity(c *gin.Context) {
	if c.Request.Method == "GET" {

		communities, err := repository.GetAllCommunity(database.DbConnection)
		if err != nil {
			result := helper.BuildResponse(false, "Get Data Community Failed", err)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Get Data Community Success", communities)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func InsertCommunity(c *gin.Context) {
	if c.Request.Method == "POST" {
		var community entity.Community

		err := c.ShouldBindJSON(&community)
		if err != nil {
			panic(err)
		}

		err = repository.InsertCommunity(database.DbConnection, community)
		if err != nil {
			result := helper.BuildResponse(false, "Create Data Community Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Create Data Community Success", community)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func UpdateCommunity(c *gin.Context) {
	if c.Request.Method == "PUT" {
		var community entity.Community

		id, _ := strconv.Atoi(c.Param("community_id"))

		err := c.ShouldBindJSON(&community)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		community.CommunityId = id

		err = repository.UpdateCommunity(database.DbConnection, community)

		if err != nil {
			result := helper.BuildResponse(false, "Update Data Community Failed", err.Error())
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Update Data Community Success", community)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func DeleteCommunity(c *gin.Context) {
	if c.Request.Method == "DELETE" {
		var community entity.Community

		id, _ := strconv.Atoi(c.Param("community_id"))

		community.CommunityId = id

		err := repository.DeleteCommunity(database.DbConnection, community)

		if err != nil {
			result := helper.BuildResponse(false, "Delete Data Community Failed", err.Error())
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Delete Data Community Success", nil)
			c.JSON(http.StatusOK, result)
		}

		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}
