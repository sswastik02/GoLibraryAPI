package api

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sswastik02/GoLibraryAPI/models"
)

func generateToken(user *models.User) (string,string,error) {

	jwtSecret := os.Getenv("JWT_SECRET")
	
	if jwtSecret == "" {
		log.Fatal("JWT Secret not found")
	}

	duration := 24 * time.Hour
	// duration := 2*time.Minute // for testing purposes


	token:= jwt.New(jwt.SigningMethodHS256)
	claims:= token.Claims.(jwt.MapClaims) // Type casting
	claims["sub"] = user.Username // sub needs to be unique (using username as sub)
	claims["exp"] = time.Now().Add(duration).Unix() // 3 minutes

	s, err := token.SignedString([]byte(jwtSecret))	

	return s,duration.String(),err

}

func getUsernameFromJWT(token *jwt.Token) string {
	claims:= token.Claims.(jwt.MapClaims)
	id:= claims["sub"].(string)
	return id
}