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

func GetAllCommunity(c *gin.Context) {
	if c.Request.Method == "GET" {

		communities := repository.GetAllCommunity(database.DbConnection)
		if communities != nil {
			result := helper.BuildResponse(false, "Get Data Community Failed", nil)
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
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func UpdateCommunity(c *gin.Context) {
	if c.Request.Method == "PUT" {
		var community entity.Community

		communityId, _ := strconv.Atoi(c.Param("community_id"))

		err := c.ShouldBindJSON(&community)
		if err != nil {
			panic(err)
		}

		community.CommunityId = int64(communityId)

		err = repository.UpdateCommunity(database.DbConnection, community)

		if err != nil {
			result := helper.BuildResponse(false, "Update Data Community Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Update Data Community Success", community)
			c.JSON(http.StatusOK, result)
		}
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func DeleteCommunity(c *gin.Context) {
	if c.Request.Method == "DELETE" {
		var community entity.Community

		communityId, _ := strconv.Atoi(c.Param("community_id"))

		err := c.ShouldBindJSON(&community)
		if err != nil {
			panic(err)
		}

		community.CommunityId = int64(communityId)

		err = repository.DeleteCommunity(database.DbConnection, community)

		if err != nil {
			result := helper.BuildResponse(false, "Delete Data Community Failed", nil)
			c.JSON(http.StatusOK, result)
		} else {
			result := helper.BuildResponse(true, "Delete Data Community Success", community)
			c.JSON(http.StatusOK, result)
		}
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	c.AbortWithStatusJSON(http.StatusBadRequest, res)
}
