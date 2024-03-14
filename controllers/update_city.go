package controllers

import (
	"net/http"
	"shrzkhn/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

type updateCityRequest struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func UpdateCity(c *gin.Context) {
	// Extract the city ID from the URL
	id := c.Param("id")

	// Get city if exist
	var city models.City
	if err := models.DB.Where("Id = ?", id).First(&city).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate request
	var req updateCityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update
	models.DB.Model(&city).Updates(req)

	c.JSON(http.StatusOK, gin.H{"data": city})
}
