package main

import (
	"github.com/gin-gonic/gin"
)

type Weather struct {
	Temperature int32  `json:"temperature"`
	Humidity    int32  `json:"humidity"`
	Description string `json:"description"`
}

type Frequency int

const (
	Hourly Frequency = iota
	Daily
)

type Subscription struct {
	Email     string    `json:"email"`
	City      string    `json:"city"`
	Frequency Frequency `json:"frequency"`
	Confirmed bool      `json:"confirmed"`
}

func main() {
	router := gin.Default()
	router.GET("/weather", GetWeatherForecast)

	router.Run("localhost:8080")
}
