package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct{
	Title string `json:"title"` // the backticks and json: specify how it will look when in json form
	Author string `json:"author"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func(r *Repository) SetupRoutes(app *fiber.App){ 
// this is how a member function of struct looks(because class is not there in go)

api:= app.Group("/api") // All endpoints will contain /api as prefix

api.Get("/books",r.getAllBooks)
api.Get("/getBook/:id",r.getBookById)
api.Post("/entryBook",r.entryBook)
api.Delete("/removeBook/:id",r.removeBookById)

}

func main(){

	// ---------------- Load the dotenv file ----------------- 

	err:=godotenv.Load(".env")
	if(err != nil){
		log.Fatal("Could not import .env file")
	}

	// ---------------- Load Configuration from env file -----------------------

	db, err := storage.NewConnection(config)

	if(err != nil){
		log.Fatal("Could not load database")
	}

	// --------------- Initialise a Repository containing the Database--------------------

	r := Repository{
		DB: db,
	}

	// --------------- Initialise Fiber Framework and routes using the repository-------------------

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8000")

}