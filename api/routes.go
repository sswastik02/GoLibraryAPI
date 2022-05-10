package api

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// =========================================== Method to setup routes ============================================

func(r *Repository) SetupRoutes(app *fiber.App){ 
	// this is how a member function of struct looks(because class is not there in go)
	// It is a struct method

	jwtSecret:=os.Getenv("JWT_SECRET")
	
	api:= app.Group("/api") // All endpoints will contain /api as prefix

	lib:=api.Group("/library")

	lib.Use(jwtware.New(jwtware.Config{
		ErrorHandler: func(context *fiber.Ctx, err error) error {
			context.Status(http.StatusUnauthorized).JSON(
				&fiber.Map{
					"message":"Unauthorized",
				},
			)
			return nil	
		},
		SigningKey: []byte(jwtSecret),
	}))

	// Adding JWT auth to lib group

	
	lib.Get("/books",r.getAllBooks)
	lib.Get("/getBook/:id",r.getBookById)
	lib.Post("/entryBook",r.entryBook)
	lib.Delete("/removeBook/:id",r.removeBookById)
	
	// fiber handles the errors returned by these functions with a default internal server error
	
	auth:=api.Group("/auth")

	auth.Post("/signup",r.signup)
	auth.Post("/signin",r.signin)
	
	}
	
	// ============================================================================================================