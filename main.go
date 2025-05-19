package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Weather struct {
	Temperature int32  `json:"temperature"`
	Humidity    int32  `json:"humidity"`
	Description string `json:"description"`
}

type Frequency string

const (
	Hourly Frequency = "HOURLY"
	Daily  Frequency = "DAILY"
)

type Subscription struct {
	Email     string    `json:"email"`
	City      string    `json:"city"`
	Frequency Frequency `json:"frequency"`
	Token     string    `json:"token"`
	Confirmed bool      `json:"confirmed"`
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	} else {
		log.Print(".env file loaded")
	}
}

func main() {

	router := gin.Default()
	router.GET("/weather", GetWeatherForecast)

	router.POST("/subscribe", SubscribeWeatherUpdates)

	router.GET("/confirm/:token", ConfirmSubscription)

	router.GET("/unsubscribe/:token", UnsubscribeWeatherUpdates)

	router.Run(":8080")
}
