package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sswastik02/Books-API/models"
)

type Book struct{
	Title string `json:"title"` // the backticks and json: specify how it will look when in json form
	Author string `json:"author"`
	Publisher string `json:"publisher"`
}

// needed a seperate structure because the SQL Model had id parameter which will not be there on Create


// ======================================= Methods for the Book API endpoints =====================================

// --------------------------------------- Entry Book Method --------------------------
func(r *Repository) entryBook(context *fiber.Ctx)error{
	
	book := Book{}
	err := context.BodyParser(&book) // This extracts the Book structure out of the json using the json rules specified the structure definiton

	if (err != nil){
		
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message":"request failed"},
		)
		return nil
	}

	response := r.DB.Create(&book)
	err = response.Error

	if (err != nil){
		
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"Could not Create Book"},
		)
		return nil
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message":"Book Added Successfully"},
	)
	
	return nil // we are supposed to return an error, In case of no error (nil)

}

// --------------------------------------- Get Book By ID Method --------------------------

func (r* Repository) getBookById(context *fiber.Ctx) error {
	bookModel := &models.Book{}
	id := context.Params("id")

	if(id == ""){
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message":"ID needs to be present",
			},
		)
		return nil
	}

	response := r.DB.Where("id = ?",id).First(bookModel)
	err:= response.Error

	if(err != nil) {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message":"Could Not Fetch with ID",
			},
		)
		return nil
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message":"Fetch with ID Successful",
			"data":bookModel,
		},
	)
	return nil
}

// --------------------------------------- Get All Books Method --------------------------

func (r *Repository) getAllBooks(context *fiber.Ctx) error {
	
	bookModels := &[]models.Book{}
	// It will contain the slice of the book model from models folder

	response := r.DB.Find(bookModels) 
	// first parameter in find is the destination variable, second parameter is conditions(which is blank)
	err:= response.Error
	if(err != nil){
		
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"All books fetch Failed"},
		)
		return nil
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message":"All books Fetched",
			"data":bookModels,
		},
	)
	return nil
}

// --------------------------------------- Remove Book Method --------------------------

func(r* Repository) removeBookById(context *fiber.Ctx) error{
	bookModel:= &models.Book{}
	// It will hold the Book to be deleted from the find
	id:= context.Params("id")

	if (id == ""){
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message":"ID needs to be present",
			},
		)
		return nil
	}

	response:= r.DB.Delete(bookModel,id)

	err:= response.Error
	if(err != nil){
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message":"Could Not Delete with ID",
			},
		)
		return nil;
	}

	if(response.RowsAffected < 1){ // This does not show as error while delete
		context.Status(http.StatusNotFound).JSON(
			&fiber.Map{
				"message":"ID does not exist",
			},
		)
		return nil;
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message":"Removed Book Successfully",
		},
	)
	return nil
}