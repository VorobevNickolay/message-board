package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var mySigningKey = []byte("BlaBlaBla123")
var ErrJwtParse = errors.New("jwt parse error")

type UserClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

func CreateToken(userid string) (string, error) {
	claims := UserClaims{
		userid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims.UserID, nil
	}
	return "", ErrJwtParse
}
