package models

import "time"

type Issue struct {
	IssueID uint `gorm:"primaryKey;autoIncrement" json:"issue_id"` // the primary key here is neccessary because that will serve as OnConflict Column, otherwise blank onConflict is not allowed
	Username string `gorm:"index" json:"username"`
	BookID uint `gorm:"unique" json:"book_id"` // unique constraint is asserted for it to be a foreign key
	Book Book `gorm:"foreignKey:BookID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"book"` // Belongs to Relationship
	IssuedAt time.Time `gorm:"autoCreateTime" json:"issued_at"`
}

type IssuesRecord struct {
	Username string `gorm:"index;primaryKey" json:"username"`
	Issues []Issue	`gorm:"foreignKey:Username;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"issues"` // Has many relationship
}