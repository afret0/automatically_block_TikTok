package tokenManager

import (
	"backend/source"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID    string
	Name  string
	Phone string
	jwt.StandardClaims
}

func GenerateToken(id, Name, phone string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3600 * 16 * time.Second)
	issuer := "pancake"
	claims := new(Claims)
	claims.ID = id
	claims.Name = Name
	claims.Phone = phone
	claims.StandardClaims.ExpiresAt = expireTime.Unix()
	claims.StandardClaims.Issuer = issuer

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SigningString()
	if err != nil {
		source.GetLogger().Errorln(id, Name, err)
	}
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	c := new(Claims)
	tokenClaims, err := jwt.ParseWithClaims(token, c, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
