package main

import (
	"cigargo/controllers"
	"cigargo/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.Use(gin.Logger())

	route.Static("/vendor", "./static/vendor")

	route.LoadHTMLGlob("templates/**/**")

	models.ConnectDB()
	models.DBMigrate()

	route.GET("/smokers", controllers.SmokersIndex)

	route.GET("/smokers/new", controllers.SmokersNew)

	route.GET("/smokers:id", controllers.SmokersShow)

	route.GET("/login", controllers.LoginPage)

	route.GET("/signup", controllers.SignupPage)

	route.POST("/smokers", controllers.SmokersCreate)

	route.GET("/", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "home/index.html", gin.H{

			"title": "Welcome to Cigarette Monitor System",
		})
	})
	log.Println("Server has started!")
	route.Run() //default port :8080
}
