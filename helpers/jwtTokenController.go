package helpers

import (
	"fmt"
	"golang_gin_gorm_jwt/models"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var expiryTime = time.Now().Local().Add(5 * time.Minute).Unix()
var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateTokens(username, usertype string) string {
	claims := &models.SignedDetails{
		Username:  username,
		User_type: usertype,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println(err)
	}
	//erro.CheckError(err)

	return token
}

func ValidateTokens(signedtoken string) bool {

	token, err := jwt.ParseWithClaims(
		signedtoken,
		&models.SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		})

	if err != nil {
		fmt.Println(err)
	}

	claims, ok := token.Claims.(*models.SignedDetails)

	if !ok && claims.ExpiresAt < time.Now().Local().Unix() && !token.Valid {
		return false
	}

	return true

}
