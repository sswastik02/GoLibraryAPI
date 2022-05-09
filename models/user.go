package models

type User struct {
	Username string `gorm:"primaryKey" json:"username"`
	Password string `json:"password"`
}