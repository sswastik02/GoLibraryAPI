package models

type User struct {
	Username string `gorm:"primaryKey" json:"username"`
	Password string `json:"password"`
	Admin bool `gorm:"type:bool;default:false" json:"admin"`
}