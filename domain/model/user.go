package model

type User struct {
	UserID      string `gorm:"column:user_id"`
	Password    string `gorm:"column:password"`
	Status      string `gorm:"column:status"`
	CreatedTime string `gorm:"column:created_time"`
	UpdatedTime string `gorm:"column:updated_time"`
}
type Login struct {
	UserID   string `json:"UserID"`
	Password string `json:"Password"`
}

type RegisterUser struct {
	UserID   string `json:"UserID"`
	Password string `json:"Password"`
	Email    string `json:"email"`
}
