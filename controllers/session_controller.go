package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(ctx *gin.Context) {

	ctx.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{},
	)
}

func SignupPage(ctx *gin.Context) {

	ctx.HTML(
		http.StatusOK,
		"home/signup.html",
		gin.H{},
	)
}
