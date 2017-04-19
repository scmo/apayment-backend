package services

import (
	"time"
	"github.com/scmo/foodchain-backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/astaxie/beego"
	"errors"
	"strings"
)

func IssueToken(user *models.User) (map[string]string) {
	//Expires the token and cookie in 1 hour
	expireToken := time.Now().Add(time.Hour * 5).Unix()

	claims := models.Claims{
		user.Username,
		jwt.StandardClaims{
			//Subject: user.Id,
			ExpiresAt:expireToken,
			Issuer:"localhost:9000",
			IssuedAt:time.Now().Unix(),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := token.SignedString([]byte(beego.AppConfig.String("jwt_token")))
	return map[string]string{"token": signedToken}
}

// middleware to protect private pages
func Validate(signedToken string) (bool) {

	token, err := jwt.ParseWithClaims(signedToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(beego.AppConfig.String("jwt_token")), nil
	})
	if err != nil {
		return false
	}

	if _, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return true
	}

	return false
}


func ParseToken(signedTokenWithBearer string) (models.Claims, error){
	claims := models.Claims{}

	signedToken, err := stripBearerPrefixFromTokenString(signedTokenWithBearer)
	if err != nil {
		beego.Error("Error while stripBearerPrefixFromTokenString.", err.Error())
		return claims, err
	}

	token, err := jwt.ParseWithClaims(signedToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(beego.AppConfig.String("jwt_token")), nil
	})
	if err != nil {
		beego.Error("Error while parsing JWT Token.", err.Error())
		return claims, err
	}

	if _, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, err
	}
	return claims, errors.New("Error while Parsing token")
}


// Strips 'Bearer ' prefix from bearer token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}