package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWeatherForecast(c *gin.Context) {
	// Simulate fetching weather data
	weather := Weather{
		Temperature: 25,
		Humidity:    60,
		Description: "Sunny",
	}

	// Return the weather data as JSON
	c.JSON(http.StatusOK, weather)
}
