package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sswastik02/GoLibraryAPI/models"
)


func(r* Repository) DeleteUser(context *fiber.Ctx) error {
	user := models.User{}
	context.BodyParser(&user)
	if user.Username == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message" : "Username cannot be blank",
			},
		)
		return nil
	}

	httpcode, err := r.AdminDBOperations(user.Username,nil)

	if err != nil {
		context.Status(httpcode).JSON(
			&fiber.Map{
				"message": err.Error(),
			},
		)
		return nil
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message":"Deleted User Successfully",
		},
	)
	return nil
}


func(r* Repository) MakeOrRemoveAdmin(makeAdmin bool) (func (*fiber.Ctx) error) {

return func (context *fiber.Ctx) error {
	user := models.User{}
	context.BodyParser(&user)

	if(user.Username == "") {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message" : "Username cannot be blank",
			},
		)
		return nil
	}

	desiredUser := models.User{
		Username: user.Username,
		Admin: makeAdmin,
	}

	httpcode , err := r.AdminDBOperations(user.Username,&desiredUser)

	if err != nil {
		context.Status(httpcode).JSON(
			&fiber.Map{
				"message" : err.Error(),
			},
		)
		return nil
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message":"Admin Change Operation Performed Successfuly",
		},
	)

	return nil

}
}

func(r* Repository) AdminSignup(context *fiber.Ctx) error{


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

	desiredUser := models.User{
		Username: user.Username,
		Password: user.Password,
		Admin: true,
	}
	
	httpcode, err := r.AdminDBOperations("",&desiredUser)

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