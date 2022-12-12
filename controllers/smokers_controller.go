package controllers

import (
	"cigargo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SmokersIndex(ctx *gin.Context) {

	smokers := models.SmokerAll()
	ctx.HTML(
		http.StatusOK,
		"smokers/index.html",
		gin.H{
			"smokers": smokers,
		},
	)
}
