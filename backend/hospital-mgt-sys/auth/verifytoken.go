package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(signedToken string) (claims *CustomClaimes, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&CustomClaimes{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secureText"), nil
		},
	)
	fmt.Println("error :- ", err)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*CustomClaimes)
	if !ok {
		var err error
		fmt.Println("erorr in claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	if (claims.Role_name != "patient") && (claims.Role_name != "doctor") {
		err = errors.New("you have no access for this route")
		return nil, err
	}
	return claims, nil
}
