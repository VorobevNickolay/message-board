package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"message-board/internal/pkg/user"
	"os"
	"time"
)

func CreateToken(userid uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	user.OnlineUsers = append(user.OnlineUsers, user.Users[userid-1])
	if err != nil {
		return "", err
	}
	return token, nil
}
