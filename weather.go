package main

import (
	"net/http"
	"romanm/web-service-gin/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetWeatherForecast(c *gin.Context, session *gorm.DB) {
	city := c.Query("city")

	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City is required"})
		return
	}

	// Simulate fetching weather data
	weatherList := db.GetWeatherForecast(session, city)

	// Return the weather data as JSON
	c.JSON(http.StatusOK, weatherList)
}
