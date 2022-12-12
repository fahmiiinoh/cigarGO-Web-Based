package controllers

import (
	"cigargo/models"
	"fmt"
	"net/http"
	"strconv"

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

func SmokersNew(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"smokers/new.html",
		gin.H{},
	)

}

func SmokersCreate(ctx *gin.Context) {
	name := ctx.PostForm("name")
	patientId := ctx.PostForm("patientId")
	healthUpdate := ctx.PostForm("healthUpdate")
	models.SmokersCreate(name, patientId, healthUpdate)
	ctx.Redirect(http.StatusMovedPermanently, "smokers")
}

func SmokersShow(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	smokers := models.SmokersFind(id)
	ctx.HTML(
		http.StatusOK,
		"smokers/show.html",
		gin.H{
			"smokers": smokers,
		},
	)
}
