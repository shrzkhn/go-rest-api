package controllers

import (
	"net/http"
	"shrzkhn/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func GetCity(c *gin.Context) {
	// Extract the city ID from the URL
	id := c.Param("id")

	var city models.City

	// Fetch the city
	result := models.DB.Where("Id = ?", id).Find(&city)
	if result.Error != nil {
		panic("Failed to get city: " + result.Error.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": city})
}
