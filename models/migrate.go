package models

import "gorm.io/gorm"


func Migrate(db *gorm.DB) error {
	// function used to AutoMigrate {Will create table}
	// This function basically creates the database for you [Unlike mongodb which creates a db when not present, Postgres needs to explicitly create a db]
	err:=db.AutoMigrate(&Book{})

	// About the table name : GORM pluralizes the struct name for which it is creating a db, in this case book

	return err
}