package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sswastik02/GoLibraryAPI/models"
)

// ======================================= Methods for the User API endpoints =====================================

// ------------------------------------------- Signup Method -----------------------------------------------------------

func(r* Repository) signup(context *fiber.Ctx) error{


	user:= models.User{}
	err:= context.BodyParser(&user)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message":"request failed",
			},
		)
		return nil
	}
	
	user.Admin = false

	httpcode, err := r.CreateUser(&user)

	if err != nil {
		context.Status(httpcode).JSON(
			&fiber.Map{
				"message":err.Error(),
			},
		)
		return nil
	}

	
	context.Path("/api/auth/signin") // signin after signup
	return context.RestartRouting()
	

}

// -------------------------------------------- Signin Method ---------------------------------------------------------

func (r* Repository) signin(context *fiber.Ctx) error {
	user:=models.User{}
	err:= context.BodyParser(&user)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message":"request failed",
			},
		)
		return nil
	}

	hashedUser, httpcode, err := r.CheckUserExist(&user)

	if err != nil {
		context.Status(httpcode).JSON(
			&fiber.Map{
				"message":err.Error(),
			},
		)
		return nil
	}

	ok := r.CheckUserCredentials(&user, hashedUser)

	if !ok  {
		context.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{
				"message":"Incorrect credentials",
			},
		)
		return nil
	}

	tokenPair,err:= generateTokenPair(hashedUser)

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message":"Could not generate JWT Token",
			},
		)
		return nil
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message":"User Signed in Successfully",
			"access_token":tokenPair["access_token"],
			"refresh_token":tokenPair["refresh_token"],
		},
	)
	return nil
}