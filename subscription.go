package main

import (
	fmt "fmt"
	"log"
	"net/http"
	"romanm/web-service-gin/config"
	"romanm/web-service-gin/randomstring"
	"romanm/web-service-gin/service"

	"github.com/gin-gonic/gin"
)

var subscriptions []Subscription

var configs = config.New()

func SubscribeWeatherUpdates(c *gin.Context) {
	email := c.Query("email")
	city := c.Query("city")
	frequency := c.Query("frequency")

	log.Println("email: ", email)
	log.Println("city: ", city)
	log.Println("frequency: ", frequency)

	Subscription := Subscription{
		Email:     email,
		City:      city,
		Frequency: Frequency(frequency),
		Confirmed: false,
		Token:     randomstring.Generate(32),
	}

	subscriptions = append(subscriptions, Subscription)

	sendEmail(email, configs.Scheme+"://"+configs.Host+"/confirm/"+Subscription.Token)

	c.JSON(http.StatusOK, Subscription)
}

func sendEmail(to string, verificationLink string) {

	mailService := service.NewSGMailService(configs)
	subject := "Email Verification for Weather Updates Service"
	mailType := service.MailConfirmation
	from := configs.FromEmail
	mailData := &service.MailData{
		Username: "User",
		Link:     verificationLink,
	}

	mailReq := mailService.NewMail(from, to, subject, mailType, mailData)
	err := mailService.SendMail(mailReq)
	if err != nil {
		fmt.Println("Unable to send mail")
		return
	}
}

func ConfirmSubscription(c *gin.Context) {
	token := c.Param("token")
	log.Println("token: ", token)

	for i, sub := range subscriptions {
		if sub.Token == token {
			subscriptions[i].Confirmed = true
			c.JSON(http.StatusOK, "subscription confirmed")
			return
		}
	}

	c.JSON(http.StatusNotFound, "subscription not found")
}

func UnsubscribeWeatherUpdates(c *gin.Context) {
	token := c.Param("token")
	log.Println("token: ", token)

	for i, sub := range subscriptions {
		if sub.Token == token {
			subscriptions = append(subscriptions[:i], subscriptions[i+1:]...)
			c.JSON(http.StatusOK, "subscription deleted")
			return
		}
	}

	c.JSON(http.StatusNotFound, "subscription not found")
}
