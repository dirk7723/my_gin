package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("luka123")

//var mySecret = "luka123"

func GetToken(username string) (string, error) {
	cla := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "myProject",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	fmt.Println("secret: ", mySecret)
	fmt.Println("token: ", token)
	tmp, err := token.SignedString(mySecret)
	fmt.Println("生成的token令牌：", tmp)
	return tmp, err
}

func ParseToken(tokenSring string) (*MyClaims, error) {
	fmt.Println("dasd")
	token, err := jwt.ParseWithClaims(tokenSring, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if Claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return Claims, nil
	}
	return nil, errors.New("invalid token")
}
