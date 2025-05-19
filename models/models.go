package models

type Weather struct {
	Id          int    `json:"id" gorm:"primaryKey"`
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
	Id        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	City      string    `json:"city"`
	Frequency Frequency `json:"frequency"`
	Token     string    `json:"token"`
	Confirmed bool      `json:"confirmed"`
}
