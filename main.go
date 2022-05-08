package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sswastik02/Books-API/api"
	"github.com/sswastik02/Books-API/models"
	"github.com/sswastik02/Books-API/storage"
)

func main(){

	// ----------------------------------- Load the dotenv file ----------------- 

	err:=godotenv.Load(".env")
	if(err != nil){
		log.Fatal("Could not import .env file")
	}

	// -------------------------------- Load Configuration from env file -----------------------

	config:= &storage.Config{ // & creates a pointer to a structure with data as declared
		Host: os.Getenv("DB_HOST") ,
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:os.Getenv("DB_NAME"),
		SSLMode:os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if(err != nil){
		log.Fatal("Could not load database")
	}

	// -------------------------------- Migrate the Database -----------------------

	err = models.Migrate(db)
	if(err != nil){
		log.Fatal("Could not Migrate Database")
	}

	// -----------------------------Initialise a Repository containing the Database--------------------

	r := api.Repository{
		DB: db,
	}

	// ------------------------- Initialise Fiber Framework and routes using the repository-------------------

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8000")

}