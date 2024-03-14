package main

import (
	"shrzkhn/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectToDatabase()

	router.Run()
}
