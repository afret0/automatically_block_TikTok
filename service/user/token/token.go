package user

import (
	"backend/source"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID   string
	Name string
	jwt.StandardClaims
}

func GenerateToken(id, Name string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3600 * 24 * time.Second)
	issuer := "pancake"
	claims := new(Claims)
	claims.ID = id
	claims.Name = Name
	claims.StandardClaims.ExpiresAt = expireTime.Unix()
	claims.StandardClaims.Issuer = issuer

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SigningString()
	if err != nil {
		source.Logger.Errorln(id, Name, err)
	}
	return token, err
}
