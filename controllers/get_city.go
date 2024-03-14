package controllers

import (
	"net/http"
	"shrzkhn/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func GetCity(c *gin.Context) {
	// Extract the city ID from the URL
	id := c.Param("id")

	// Fetch the city
	var city models.City
	if err := models.DB.Where("Id = ?", id).First(&city).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": city})
}
