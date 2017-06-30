package services

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"github.com/scmo/apayment-backend/models"
	"strings"
	"time"
)

func IssueToken(user *models.User) map[string]string {

	var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}
	//Expires the token and cookie in 1 hour
	expireToken := time.Now().Add(time.Hour * 5).Unix()
	claims := models.Claim{
		roles,
		jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expireToken,
			Issuer:    "localhost:9000",
			IssuedAt:  time.Now().Unix(),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := token.SignedString([]byte(beego.AppConfig.String("jwt_secret_secret")))
	return map[string]string{"token": signedToken}
}

// middleware to protect private pages
func Validate(signedTokenWithBearer string) bool {
	signedToken, err := stripBearerPrefixFromTokenString(signedTokenWithBearer)
	if err != nil {
		beego.Error("Error while stripBearerPrefixFromTokenString.", err.Error())
		return false
	}
	token, err := jwt.ParseWithClaims(signedToken, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(beego.AppConfig.String("jwt_secret_secret")), nil
	})
	if err != nil {
		beego.Error(err)
		return false
	}

	if _, ok := token.Claims.(*models.Claim); ok && token.Valid {
		return true
	}

	return false
}

func ParseToken(signedTokenWithBearer string) (models.Claim, error) {
	claims := models.Claim{}

	signedToken, err := stripBearerPrefixFromTokenString(signedTokenWithBearer)
	if err != nil {
		beego.Error("Error while stripBearerPrefixFromTokenString.", err.Error())
		return claims, err
	}

	token, err := jwt.ParseWithClaims(signedToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(beego.AppConfig.String("jwt_secret_secret")), nil
	})
	if err != nil {
		beego.Error("Error while parsing JWT Token.", err.Error())
		return claims, err
	}

	if _, ok := token.Claims.(*models.Claim); ok && token.Valid {
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
