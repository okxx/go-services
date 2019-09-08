package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"go-services/global/common/utils/auth/md5"
)

var jwtSecret = []byte("iri9*w\x99")

type Claims struct {
	Pk		 string	`json:"pk"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func Build(pk,username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		pk, md5.EncodeMD5(username), md5.EncodeMD5(password),
		jwt.StandardClaims{
			Id			: pk,
			ExpiresAt	: expireTime.Unix(),
			Issuer		: "iri9",
			Subject		: username,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func Parse(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}