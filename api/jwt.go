package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sswastik02/GoLibraryAPI/models"
	"gorm.io/gorm"
)

var jwtSecret string

func InitializeJwtSecret(secret string) {
	if secret == "" {
		log.Fatal("JWT Secret not found")
	}
	jwtSecret = secret
}





func verifySignature(refreshToken *jwt.Token) (interface{}, error) {
	_,ok := refreshToken.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil,fmt.Errorf("unexpected signing method : %s",refreshToken.Header["alg"])
	}	

	return []byte(jwtSecret),nil
}





func generateTokenPair(user *models.User) (map[string]string,error) {

	accessTokenDuration := 5 * time.Minute
	refreshTokenDuration := 24 * time.Hour
	// refreshTokenDuration := 5 * time.Minute // for testing purposes




	accessToken:= jwt.New(jwt.SigningMethodHS256)
	claims:= accessToken.Claims.(jwt.MapClaims) // Type casting
	claims["sub"] = user.Username // sub needs to be unique (using username as sub)
	claims["exp"] = time.Now().Add(accessTokenDuration).Unix()
	claims["admin"] = user.Admin
	claims["access"] = true

	accessTokenString, err := accessToken.SignedString([]byte(jwtSecret))	

	if err != nil {
		return map[string]string{},err
	}




	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rfclaims := refreshToken.Claims.(jwt.MapClaims)
	rfclaims["sub"] = user.Username
	rfclaims["exp"] = time.Now().Add(refreshTokenDuration).Unix()

	refreshTokenString, err := refreshToken.SignedString([]byte(jwtSecret))

	if err != nil {
		return map[string]string{},err
	}

	return map[string]string{
		"access_token":accessTokenString,
		"refresh_token":refreshTokenString,
	},nil

}

func getTokenFromContext(context *fiber.Ctx) (*jwt.Token,error) {
	bearer := context.GetReqHeaders()["Authorization"]
	tokenString := strings.Split(bearer, "Bearer ")[1]
	token,err := jwt.Parse(tokenString,verifySignature)
	
	if err != nil {
		return nil,fmt.Errorf("unexpected signing method")
	}

	return token,nil
}




func checkIfValidAccessToken(context *fiber.Ctx,checkIfAdmin bool) error {
	// Allow only JWT tokens with the access claims 


	token,err:= getTokenFromContext(context)

	if err != nil {
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	
	if claims["access"] == nil {
		return fmt.Errorf("malformed or invalid JWT")
	}

	if !claims["access"].(bool) {
		return fmt.Errorf("not a valid Access token")
	}

	if checkIfAdmin {
		admin :=claims["admin"].(bool)
		if !admin {
			return fmt.Errorf("only Administrators can access this endpoint")
		}
	}

	return nil
}



func jwtErrorHandler (context *fiber.Ctx, err error) error {
	context.Status(http.StatusUnauthorized).JSON(
		&fiber.Map{
			"message":"Unauthorized",
		},
	)
	return nil	
}

func(r *Repository) getBlackListAndAdminSuccessHandler(checkIfAdmin bool) (func (*fiber.Ctx) error) {

	return func (context *fiber.Ctx) error {
	err := r.BlackListAndAdminJWTFilter(context,checkIfAdmin)

	// Unfortunately adding this filter logic to Filter in the jwtware config will not work because
	// The filter present there is to accept and not reject

	if err != nil {
		context.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{
				"message": err.Error() ,
			},
		)
		return nil
	}
	return context.Next()
	}
}



func(r* Repository) BlackListAndAdminJWTFilter(context *fiber.Ctx,checkIfAdmin bool) error {
	err := checkIfValidAccessToken(context,checkIfAdmin)
	
	if err != nil {
		return err
	}


	token, err := getTokenFromContext(context)

	if err != nil {
		return err
	}

	user := getUserFromJWT(token)
	s,err := r.RdbClientOperations(user.Username)

	if err != nil {
		return err
	}

	if s == "NoAccess" {
		return fmt.Errorf("user is deleted")
	}

	if checkIfAdmin && s == "NoAdmin" {
		return fmt.Errorf("user has no administraive privileges anymore")
	}

	return nil
}




func jwtUserMiddleware(r *Repository) (func(*fiber.Ctx) error){
	return jwtware.New(jwtware.Config{
		SuccessHandler: r.getBlackListAndAdminSuccessHandler(false),
		ErrorHandler: jwtErrorHandler,
		
		SigningKey: []byte(jwtSecret),
	})
}




func jwtAdminMiddleware(r* Repository) (func(*fiber.Ctx) error ){
	return jwtware.New(jwtware.Config{
		SuccessHandler: r.getBlackListAndAdminSuccessHandler(true),
		ErrorHandler: jwtErrorHandler,
		SigningKey: []byte(jwtSecret),
	})
}





func getUserFromJWT(token *jwt.Token) models.User {
	claims:= token.Claims.(jwt.MapClaims)
	username:= claims["sub"].(string)
	admin:= claims["admin"].(bool)
	return models.User{
		Username: username,
		Admin: admin,
	}
}




func(r* Repository) RefreshTokenPair(context *fiber.Ctx) error {

	type RefreshTokenBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	refreshTokenBody := RefreshTokenBody{}
	context.BodyParser(&refreshTokenBody)



	token, err := jwt.Parse(refreshTokenBody.RefreshToken,verifySignature )

	if err != nil {
		context.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{
				"message":err.Error(),
			},
		)
		return nil
	}
	

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		context.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{
				"message":"Malformed or Invalid JWT",
			},
		)
		return nil
	}

	if claims["access"] != nil && claims["access"].(bool){
		context.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{
				"message":"Cannot refresh using access token",
			},
		)
		return nil
	}



	username:= getUserFromJWT(token)
	user := models.User{}
	response := r.DB.Where("username = ?",username).First(&user)
	err = response.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
		context.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{
				"message": "Malformed or Invalid JWT",
			},
		)
		
		return nil
	}
	context.Status(http.StatusInternalServerError).JSON(
		&fiber.Map{
			"message":"Could not Refresh Tokens",
		},
	)
	return nil
	}

	newTokenPair, err := generateTokenPair(&user)

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message":"Could not Refresh Tokens",
			},
		)
		return nil
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
		"access_token":newTokenPair["access_token"],
		"refresh_token":newTokenPair["refresh_token"],
		},
	)

	return nil

}