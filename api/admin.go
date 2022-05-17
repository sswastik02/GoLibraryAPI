package api

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sswastik02/GoLibraryAPI/models"
	"gorm.io/gorm"
)

func(r* Repository) AdminDBOperations(username string, desiredUser *models.User) (int, error) {
	if username == "" {
		httpcode, err := r.CreateUser(desiredUser)
		return httpcode,err
	}
	var response *gorm.DB
	if desiredUser == nil {
		response = r.DB.Where("username = ?",username).Delete(&models.User{})
	} else {
		// response = r.DB.Where("username = ?",username).Updates(
		// 	models.User{
		// 		Username: desiredUser.Username,
		// 		Admin: desiredUser.Admin,
		// 	},
		// ) 
		//  Unfortunately the above method does not work to update many fields because it ignores values like 0 and false
		
		desiredUserMap := make(map[string]interface {})
		desiredUserMap["username"] = desiredUser.Username
		desiredUserMap["admin"] = desiredUser.Admin
		response = r.DB.Model(&models.User{}).Where("username = ?",username).Updates(&desiredUserMap	)
	}
		err := response.Error

		if err != nil {
			return http.StatusInternalServerError,err
		}

		if response.RowsAffected < 1 {
			return http.StatusBadRequest, fmt.Errorf("user with that username not found")
		}

		return http.StatusOK,nil
}

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