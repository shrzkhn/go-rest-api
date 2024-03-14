package main

import (
	"shrzkhn/go-rest-api/controllers"
	"shrzkhn/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectToDatabase()

	router.POST("/api/city", controllers.AddCity)
	router.GET("/api/city/:id", controllers.GetCity)
	router.PATCH("/api/city/:id", controllers.UpdateCity)
	router.DELETE("api/city/:id", controllers.RemoveCity)

	router.Run()
}
