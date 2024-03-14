package controllers

import (
	"net/http"
	"shrzkhn/go-rest-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type addCityRequest struct {
	Name    string `json:"name" binding:"required"`
	Country string `json:"country" binding:"required"`
}

func AddCity(c *gin.Context) {
	// Validate request
	var req addCityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add city
	city := models.City{
		Id:      uuid.NewString(),
		Name:    req.Name,
		Country: req.Country,
	}
	models.DB.Create(&city)

	c.JSON(http.StatusCreated, gin.H{"data": city})
}
