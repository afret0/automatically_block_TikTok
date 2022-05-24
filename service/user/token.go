package user

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"service/source"
	"time"
)

var j *JWT
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	signKey          []byte = []byte("AllYourBase")
)

func init() {
	j = new(JWT)
	j.logger = source.GetLogger()
}

type JWT struct {
	logger *logrus.Logger
}

type Claims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func (j *JWT) GenerateToken(id, Name, email string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3600 * 16 * time.Second)
	issuer := "pancake"
	c := Claims{id, Name, email, jwt.StandardClaims{ExpiresAt: expireTime.Unix(), Issuer: issuer}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenS, err := token.SignedString(signKey)
	if err != nil {
		source.GetLogger().Errorln(id, Name, err)
	}
	return tokenS, err
}

func (j *JWT) ParseToken(token string) (*Claims, error) {
	c := new(Claims)
	tokenClaims, err := jwt.ParseWithClaims(token, c, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		j.logger.Errorln(token, err)
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func GetJWT() *JWT {
	return j
}
