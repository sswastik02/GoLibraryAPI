package api

import (
	"net/http"

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

	

	app.Get("/",func(context *fiber.Ctx) error {
		context.Status(http.StatusOK).JSON(
			&fiber.Map{
				"message":"Access all endpoints at /api",
			},
		)
		return nil
	})
	
	api:= app.Group("/api") // All endpoints will contain /api as prefix

	lib:=api.Group("/library")

	lib.Use(jwtMiddleware())

	// Adding JWT auth to lib group

	
	lib.Get("/books",r.getAllBooks)
	lib.Get("/getBook/:id",r.getBookById)
	lib.Post("/entryBook",r.entryBook)
	lib.Delete("/removeBook/:id",r.removeBookById)
	
	// fiber handles the errors returned by these functions with a default internal server error
	
	auth:=api.Group("/auth")

	auth.Post("/signup",r.signup)
	auth.Post("/signin",r.signin)
	auth.Post("/refreshTokens",r.RefreshTokenPair)
	
	}
	
	// ============================================================================================================