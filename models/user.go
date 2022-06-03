package models

type User struct {
	Username string `gorm:"primaryKey" json:"username"`
	Password string `json:"password"`
	Admin bool `gorm:"default:false" json:"admin"`
	IssuesRecord IssuesRecord `gorm:"foreignKey:Username;constraint:OnDelete:CASCADE"` // has one relationship
}