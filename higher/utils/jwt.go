package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtSecret = []byte("jwtSecret")

const JWT_CONTEXT_KEY = "jwt_context_key"

type Token struct {
	Name string
	DcId int
	jwt.StandardClaims
}

func CreateJwtToken(name string, dcId int) (string, error) {
	var token Token
	token.StandardClaims = jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		Id:        "",
		IssuedAt:  time.Now().Unix(),
		Issuer:    "kit",
		NotBefore: time.Now().Unix(),
		Subject:   "login",
	}
	token.Name = name
	token.DcId = dcId
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	return tokenClaims.SignedString(jwtSecret)
}

func ParseToken(token string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || jwtToken == nil {
		return nil, err
	}
	claim, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid {
		return claim, nil
	} else {
		return nil, nil
	}
}
