package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// =========================================== Method to setup routes ============================================

func(r *Repository) SetupRoutes(app *fiber.App){ 
	// this is how a member function of struct looks(because class is not there in go)
	// It is a struct method
	
	api:= app.Group("/api") // All endpoints will contain /api as prefix
	
	api.Get("/books",r.getAllBooks)
	api.Get("/getBook/:id",r.getBookById)
	api.Post("/entryBook",r.entryBook)
	api.Delete("/removeBook/:id",r.removeBookById)
	
	// fiber handles the errors returned by these functions with a default internal server error
	
	auth:=api.Group("/auth")

	auth.Post("/signup",r.signup)
	auth.Post("/signin",r.signin)
	
	}
	
	// ============================================================================================================