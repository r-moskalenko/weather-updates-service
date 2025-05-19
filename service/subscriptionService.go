package service

import (
	fmt "fmt"
	"log"
	"net/http"
	"regexp"
	"romanm/web-service-gin/config"
	"romanm/web-service-gin/db"
	"romanm/web-service-gin/models"
	"romanm/web-service-gin/randomstring"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SubscriptionService interface {
	SubscribeWeatherUpdates(c *gin.Context, session *gorm.DB)
	ConfirmSubscription(c *gin.Context, session *gorm.DB)
	UnsubscribeWeatherUpdates(c *gin.Context, session *gorm.DB)
}

// Define a concrete implementation of the interface
type subscriptionServiceImpl struct{}

type Subscription struct {
	Email     string
	City      string
	Frequency Frequency
	Confirmed bool
	Token     string
}

func NewSubscriptionService() SubscriptionService {
	return &subscriptionServiceImpl{}
}

// Subscription represents a user's subscription to weather updates.

type Frequency string

const (
	Hourly Frequency = "HOURLY"
	Daily  Frequency = "DAILY"
)

var configs = config.New()

func (s *subscriptionServiceImpl) SubscribeWeatherUpdates(c *gin.Context, session *gorm.DB) {
	email := c.Query("email")
	city := c.Query("city")
	frequency := c.Query("frequency")

	if !validateEmail(email) {
		log.Println("Invalid email format: ", email)
		c.JSON(http.StatusBadRequest, "Invalid input")
		return
	}

	log.Println("email: ", email)
	log.Println("city: ", city)
	log.Println("frequency: ", frequency)

	subscriptionFromDb, _ := db.GetSubscriptionsByEmail(session, email)

	if subscriptionFromDb != nil {
		log.Println("Subscription already exists: ", subscriptionFromDb)
		c.JSON(http.StatusConflict, "Email already subscribed")
		return
	}

	subscription := models.Subscription{
		Email:     email,
		City:      city,
		Frequency: models.Frequency(frequency),
		Confirmed: false,
		Token:     randomstring.Generate(32),
	}

	db.CreateSubscription(session, &subscription)

	s.sendEmail(email, configs.Scheme+"://"+configs.Host+"/confirm/"+subscription.Token)

	c.JSON(http.StatusOK, subscription)
}

func validateEmail(email string) bool {
	// Define the stricter email regex pattern (based on RFC 5322)
	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`

	// Compile the regex
	re := regexp.MustCompile(emailRegex)

	if re.MatchString(email) {
		return true
	} else {
		return false
	}
}

func (s *subscriptionServiceImpl) sendEmail(to string, verificationLink string) {

	mailService := NewSGMailService(configs)
	subject := "Email Verification for Weather Updates Service"
	mailType := MailConfirmation
	from := configs.FromEmail
	mailData := &MailData{
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

func (s *subscriptionServiceImpl) ConfirmSubscription(c *gin.Context, session *gorm.DB) {
	token := c.Param("token")
	log.Println("token: ", token)

	if len(token) != 32 {
		log.Println("Invalid token length: ", token)
		c.JSON(http.StatusBadRequest, "Invalid token")
		return
	}

	subscription, err := db.GetSubscriptionsByToken(session, token)

	if err != nil {
		log.Println("Error fetching subscriptions, token not found: ", err)
		c.JSON(http.StatusNotFound, "Token not found")
		return
	}

	subscription.Confirmed = true

	err = db.UpdateSubscription(session, subscription)

	if err != nil {
		log.Println("Error updating subscription: ", err)
		c.JSON(http.StatusInternalServerError, "error updating subscription")
		return
	}

	c.JSON(http.StatusOK, "Subscription confirmed successfully")
}

func (s *subscriptionServiceImpl) UnsubscribeWeatherUpdates(c *gin.Context, session *gorm.DB) {
	token := c.Param("token")
	log.Println("token: ", token)

	if len(token) != 32 {
		log.Println("Invalid token length: ", token)
		c.JSON(http.StatusBadRequest, "Invalid token")
		return
	}

	subscription, err := db.GetSubscriptionsByToken(session, token)

	if err != nil {
		log.Println("Error fetching subscriptions: ", err)
		c.JSON(http.StatusNotFound, "Token not found")
		return
	}

	err = db.DeleteSubscription(session, subscription.Id)
	if err != nil {
		log.Println("Error deleting subscription: ", err)
		c.JSON(http.StatusInternalServerError, "error deleting subscription")
		return
	}

	c.JSON(http.StatusOK, "Unsubscribed successfully")
}
