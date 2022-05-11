package api

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sswastik02/GoLibraryAPI/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	if(user.Username == "" || user.Password == "") {
		context.Status(http.StatusOK).JSON(
			&fiber.Map{
				"message":"Username or Password cannot be blank",
			},
		)
		return nil
	}



	usernameReg , _:= regexp.Compile("^[A-Za-z][A-Za-z0-9_]{7,29}$") 
	// 8-30 characters starting should be alphabet, other can be alphanum or underscore

	passwordReg, _ := regexp.Compile("^[A-Za-z0-9_@$^&*]{8,30}$")
	// 8-30 characters starting should be alphabet, other can be alphanum or underscore

	
	if !(usernameReg.MatchString(user.Username)) {
		context.Status(http.StatusOK).JSON(
			&fiber.Map{
				"message":"Username should have 8-30 characters starting should be alphabet, other can be alphanumeric or underscore",
			},
		)
		return nil
	}

	if !(passwordReg.MatchString(user.Password)) {
		context.Status(http.StatusOK).JSON(
			&fiber.Map{
				"message":"Password should have 8-30 characters, Containing alphanumeric, _ , @, $, ^, & or *",
			},
		)
		return nil
	}





	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),8) // 8 is chosen arbitrarily

	
	if err != nil {
		fmt.Printf("Could not hash Password for %s",user.Username)
		return nil;
	}

	hashedUser := models.User{
		Username: user.Username,
		Password: string(hashedPassword),
	}

	response:= r.DB.Create(&hashedUser)
	err = response.Error

	if err != nil {

		if strings.Contains(err.Error(),"duplicate key value violates unique constraint") {
			context.Status(http.StatusOK).JSON(
				&fiber.Map{
					"message":"User with that username already exists",
				},
			)
			return nil
		}

		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message":"Could not create user",
			},
		)
		return nil
	}

	
	context.Redirect("/signin",http.StatusOK) // signin after signup
	return nil

}

// -------------------------------------------- Signin Method ---------------------------------------------------------

func (r* Repository) signin(context *fiber.Ctx) error {
	user:=models.User{}
	err:= context.BodyParser(&user)

	if(user.Username == "" || user.Password == "") {
		context.Status(http.StatusOK).JSON(
			&fiber.Map{
				"message":"Username or Password cannot be blank",
			},
		)
		return nil
	}

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message":"request failed",
			},
		)
		return nil
	}

	hashedUser:= models.User{}
	response:=r.DB.Where("username = ?",user.Username).First(&hashedUser)
	err = response.Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
		context.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{
				"message":"Could not find user with that username",
			},
		)
		return nil
		} 
		
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message":"Could not login user",
			},
		)

		return nil

	}

	pass := bcrypt.CompareHashAndPassword([]byte(hashedUser.Password),[]byte(user.Password))

	if pass != nil {
		context.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{
				"message":"Incorrect credentials",
			},
		)
		return nil
	}

	token,duration,err:= generateToken(&user)

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
			"token":token,
			"duration":duration,
		},
	)
	return nil
}