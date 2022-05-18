package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sswastik02/GoLibraryAPI/database"
	"github.com/sswastik02/GoLibraryAPI/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
	Client *redis.Client
}


func(r* Repository) RdbClientOperations(username string, blacklist... string) (string,error) {
	if len(blacklist) == 0 {
		scmd := r.Client.Get(database.Ctx,username)
		if scmd.Err() == redis.Nil {
			return "",nil
		}

		if scmd.Err() != nil {
			return "",fmt.Errorf("unknown error occured")
		}

		return scmd.Result()
	}

	scmd := r.Client.Set(database.Ctx,username,blacklist[0],2*time.Minute)
	return scmd.Result()
}

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

	if desiredUser == nil {
		_,err := r.RdbClientOperations(username,"NoAccess")
		if err != nil {
			return http.StatusInternalServerError,fmt.Errorf("unknown error occured")
		}
	} else if !desiredUser.Admin {
		_,err := r.RdbClientOperations(username,"NoAdmin")
		if err != nil {
			return http.StatusInternalServerError, fmt.Errorf("unknown error occured")
		}
	}

	return http.StatusOK,nil
}
