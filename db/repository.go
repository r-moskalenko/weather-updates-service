package db

import (
	"log"
	"romanm/web-service-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:postgres@go_db:5432/postgres"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Weather{})
	db.AutoMigrate(&models.Subscription{})

	db.Create(&models.Weather{Temperature: 23, Humidity: 60, Description: "Sunny"})
	db.Create(&models.Weather{Temperature: 25, Humidity: 70, Description: "Cloudy"})
	db.Create(&models.Weather{Temperature: 30, Humidity: 80, Description: "Rainy"})
	db.Create(&models.Weather{Temperature: 28, Humidity: 65, Description: "Windy"})

	return db
}

func GetWeatherForecast(db *gorm.DB, city string) []models.Weather {
	var weatherList []models.Weather
	// Get all records
	result := db.Find(&weatherList)
	// SELECT * FROM users;
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return weatherList
}

func CreateSubscription(db *gorm.DB, subscription *models.Subscription) error {
	result := db.Create(subscription)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateSubscription(db *gorm.DB, subscription *models.Subscription) error {
	result := db.Save(subscription)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllSubscriptions(db *gorm.DB) ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	result := db.Find(&subscriptions)
	if result.Error != nil {
		return nil, result.Error
	}
	return subscriptions, nil
}

func GetSubscriptionsByToken(db *gorm.DB, token string) (*models.Subscription, error) {
	var subscription models.Subscription
	result := db.Where("token = ?", token).First(&subscription)
	if result.Error != nil {
		return nil, result.Error
	}
	return &subscription, nil
}

func GetSubscriptionsByEmail(db *gorm.DB, email string) (*models.Subscription, error) {
	var subscription models.Subscription
	result := db.Where("email = ?", email).First(&subscription)
	if result.Error != nil {
		return nil, result.Error
	}
	return &subscription, nil
}

func DeleteSubscription(db *gorm.DB, id int) error {
	var subscription models.Subscription
	result := db.Where("id = ?", id).Delete(&subscription)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
