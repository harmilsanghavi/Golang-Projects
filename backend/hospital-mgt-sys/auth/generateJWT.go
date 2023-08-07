package auth

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaimes struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Id        int    `json:"id"`
	Role_name string `json:"role_name"`
	jwt.StandardClaims
}

type Loginresp struct {
	Success      bool   `json:"success"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func SetJwt(id int, email, password string, Role_name string) *Loginresp {
	claims :=
		CustomClaimes{
			Email:     email,
			Password:  password,
			Id:        id,
			Role_name: Role_name,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
				Issuer:    "vivek.com",
			},
		}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("secureText"))

	if err != nil {
		log.Fatal(err)
	}

	//refresh token
	claims2 :=
		CustomClaimes{
			Email:     email,
			Password:  password,
			Id:        id,
			Role_name: Role_name,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
				Issuer:    "vivek.com",
			},
		}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)

	signedTokenForrefresh, err := refreshToken.SignedString([]byte("secureText"))

	var logresp Loginresp

	if err != nil {
		signedToken = "0"
		signedTokenForrefresh = "0"
		logresp = Loginresp{
			Success:      false,
			Token:        signedToken,
			RefreshToken: signedTokenForrefresh,
		}
	} else {
		logresp = Loginresp{
			Success:      true,
			Token:        signedToken,
			RefreshToken: signedTokenForrefresh,
		}
	}

	return &logresp
}
