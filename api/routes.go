package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

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


	auth:=api.Group("/auth")
	lib:=api.Group("/library")
	admin := api.Group("/admin")

	lib.Use(jwtUserMiddleware(r))
	// Adding JWT auth to lib group
	admin.Use(jwtAdminMiddleware(r))



	auth.Post("/signup",r.signup)
	auth.Post("/signin",r.signin)
	auth.Post("/refreshTokens",r.RefreshTokenPair)

	
	lib.Get("/books",r.getAllBooks)
	lib.Get("/getBook/:id",r.getBookById)

	lib.Post("/entryBook",jwtAdminMiddleware(r),r.entryBook)
	lib.Delete("/removeBook/:id",jwtAdminMiddleware(r),r.removeBookById) // this is how you implement middleware in each 
	
	// fiber handles the errors returned by these functions with a default internal server error
	


	admin.Post("/create",r.AdminSignup)
	admin.Post("/makeAdmin",r.MakeOrRemoveAdmin(true))
	admin.Post("/removeAdmin",r.MakeOrRemoveAdmin(false))
	admin.Delete("/deleteUser",r.DeleteUser)
	}
	
	// ============================================================================================================