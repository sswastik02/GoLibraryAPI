package main

import (
	"log"
	"net/http"

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

// ======================================= Methods for the API endpoints =====================================

func(r *Repository) entryBook(context *fiber.Ctx)error{
	
	book := Book{}
	err := context.BodyParser(&book) // This extracts the Book structure out of the json using the json rules specified the structure definiton

	if (err != nil){
		log.Fatal("Could not Parse JSON for Book")
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message":"request failed"},
		)
		return err
	}

	response := r.DB.Create(&book)
	err = response.Error

	if (err != nil){
		log.Fatal("Could not insert into database")
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"Could not Create Book"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message":"Book Added Successfully"},
	)
	
	return nil // we are supposed to return an error, In case of no error (nil)

}

func (r *Repository) getAllBooks(context *fiber.Ctx) error {
	
	bookModels := &[]models.Book{}
	// It will contain the slice of the book model from models folder

	response := r.DB.Find(bookModels) 
	// first parameter in find is the destination variable, second parameter is conditions(which is blank)
	err:= response.Error
	if(err != nil){
		log.Fatal("Could not find all books")
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"All books fetch Failed"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message":"All books Fetched",
			"data":bookModels,
		},
	)
	return nil
}



// ====================================== Method to setup routes ============================================

func(r *Repository) SetupRoutes(app *fiber.App){ 
// this is how a member function of struct looks(because class is not there in go)
// It is a struct method

api:= app.Group("/api") // All endpoints will contain /api as prefix

api.Get("/books",r.getAllBooks)
api.Get("/getBook/:id",r.getBookById)
api.Post("/entryBook",r.entryBook)
api.Delete("/removeBook/:id",r.removeBookById)

}

// ============================================================================================================



func main(){

	// ----------------------------------- Load the dotenv file ----------------- 

	err:=godotenv.Load(".env")
	if(err != nil){
		log.Fatal("Could not import .env file")
	}

	// -------------------------------- Load Configuration from env file -----------------------

	db, err := storage.NewConnection(config)

	if(err != nil){
		log.Fatal("Could not load database")
	}

	// -----------------------------Initialise a Repository containing the Database--------------------

	r := Repository{
		DB: db,
	}

	// ------------------------- Initialise Fiber Framework and routes using the repository-------------------

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8000")

}