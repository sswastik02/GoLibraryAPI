package api

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/sswastik02/GoLibraryAPI/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func(r* Repository) CreateUser(user *models.User) (int, error) {
	if(user.Username == "" || user.Password == "") {
		return http.StatusBadRequest, fmt.Errorf("username or password cannot be blank")
	}

	usernameReg , _:= regexp.Compile("^[A-Za-z][A-Za-z0-9_]{7,29}$") 
	// 8-30 characters starting should be alphabet, other can be alphanum or underscore

	passwordReg, _ := regexp.Compile("^[A-Za-z0-9_@$^&*]{8,30}$")
	// 8-30 characters starting should be alphabet, other can be alphanum or underscore

	
	if !(usernameReg.MatchString(user.Username)) {
		return http.StatusBadRequest ,fmt.Errorf("username should have 8-30 characters starting should be alphabet, other can be alphanumeric or underscore")
	}

	if !(passwordReg.MatchString(user.Password)) {
		return http.StatusBadRequest ,fmt.Errorf("password should have 8-30 characters, Containing alphanumeric, _ , @, $, ^, & or *")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),8) // 8 is chosen arbitrarily

	
	if err != nil {
		fmt.Printf("could not hash Password for %s",user.Username)
		return http.StatusInternalServerError ,fmt.Errorf("could not create user")
	}

	hashedUser := models.User{
		Username: user.Username,
		Password: string(hashedPassword),
		Admin: user.Admin,
	}

	response:= r.DB.Create(&hashedUser)
	err = response.Error

	if err != nil {

		if strings.Contains(err.Error(),"duplicate key value violates unique constraint") {
			return http.StatusConflict ,fmt.Errorf("user with that username already exists")
		}

		return http.StatusInternalServerError, fmt.Errorf("could not create user")
	}

	return http.StatusOK, nil
}




func (r* Repository) CheckUserExist(user *models.User) (*models.User,int, error) {
	dummyUser := &models.User{}
	if(user.Username == "" || user.Password == "") {
		return dummyUser,http.StatusBadRequest, fmt.Errorf("username or password cannot be blank")
	}

	hashedUser:= models.User{}
	response:=r.DB.Where("username = ?",user.Username).First(&hashedUser)
	err := response.Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
		return dummyUser, http.StatusConflict, fmt.Errorf("could not find user with that username")
		} 

		return dummyUser, http.StatusInternalServerError, fmt.Errorf("error occured in server")

	}

	return &hashedUser,http.StatusOK, nil
}




func(r* Repository) CheckUserCredentials(user *models.User,hashedUser *models.User) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedUser.Password),[]byte(user.Password))

	return err == nil
}