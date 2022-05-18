package database

import (
	"fmt"

	"gorm.io/driver/postgres" // Using Postgres Driver provided by gorm
	"gorm.io/gorm"
)

type PGConfig struct{
	Host		string
	Port		string
	User		string
	Password	string
	DBName		string
	SSLMode		string
}

// Config structure holds the information for the Database

func NewConnection(config *PGConfig)(*gorm.DB, error){
	// This function will open a new database using the config file provided
	dsn:= fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,config.Port,config.User,config.Password,config.DBName,config.SSLMode,
	)

	db, err:= gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db,err;
}
