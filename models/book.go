package models

import "gorm.io/gorm"

type Book struct{
	ID 			uint		`gorm:"primary key;autoincrement" json:"id"`
	Title 		*string		`json:"title"`
	Author 		*string		`json:"author"`
	Publisher 	*string		`json:"publisher"`

}

// reason for using *string is That field might be nil. (Otherwise the default value is assigned "" for string)

func Migrate(db *gorm.DB) error {
	// function used to AutoMigrate {Will create table}
	// This function basically creates the database for you [Unlike mongodb which creates a db when not present, Postgres needs to explicitly create a db]
	err:=db.AutoMigrate(&Book{})
	return err
}