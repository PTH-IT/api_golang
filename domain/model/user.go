package model

type User struct {
	UserID     string `gorm:"UserID"`
	Password   string `gorm:"Password"`
	Status     string `gorm:"status"`
	TimeCreate string `gorm:"created_time"`
	TimeUpdate string `gorm:"updated_time"`
}
