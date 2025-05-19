package main

import (
	"log"

	"romanm/web-service-gin/db"
	"romanm/web-service-gin/service"

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

	db := db.Init()

	subscriptionService := service.NewSubscriptionService()

	router := gin.Default()
	router.GET("/weather", func(ctx *gin.Context) {
		GetWeatherForecast(ctx, db)
	})

	router.POST("/subscribe", func(c *gin.Context) {
		subscriptionService.SubscribeWeatherUpdates(c, db)
	})

	router.GET("/confirm/:token", func(ctx *gin.Context) {
		subscriptionService.ConfirmSubscription(ctx, db)
	})

	router.GET("/unsubscribe/:token", func(ctx *gin.Context) {
		subscriptionService.UnsubscribeWeatherUpdates(ctx, db)
	})

	router.Run(":8080")
}
