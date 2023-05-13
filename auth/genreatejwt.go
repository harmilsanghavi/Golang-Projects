package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}
type LoginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

func SetJwt(name, password string) *LoginResponse {
	claims := CustomClaims{
		Name:     name,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "nameOfWebsiteHere",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("secureSecretText"))
	var loginResponse LoginResponse
	if err != nil {
		fmt.Println(err)
		signedToken = "0"
		loginResponse = LoginResponse{
			Status: "Login Failed",
			Token:  signedToken,
		}
	} else {
		fmt.Println(signedToken)
		loginResponse = LoginResponse{
			Status: "Login successful!",
			Token:  signedToken,
		}
	}
	return &loginResponse
}
