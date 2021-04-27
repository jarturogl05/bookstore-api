package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jarturogl05/go-api-gin/controllers"
	"github.com/jarturogl05/go-api-gin/models"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()
	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run()
}
