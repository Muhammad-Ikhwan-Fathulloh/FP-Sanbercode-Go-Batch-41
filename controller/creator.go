package controller

import (
	"FP-Sanbercode-Go-Batch-41/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCreator(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {

		res := helper.BuildResponse(true, "Get Data Success", gin.H{
			"Creator":    "Muhammad Ikhwan Fathulloh",
			"Portofolio": "https://ikhwanfathulloh.nocturnailed.tech/",
		})

		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildErrorResponse("Data not found", "Method Failed", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
}
