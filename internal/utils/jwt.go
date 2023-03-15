package utils

import (
	"cuckoo/internal/conf"
	"cuckoo/internal/domain/entity"
	"cuckoo/pkg/valobj"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	valobj.UserInfo
	jwt.StandardClaims
}

func GenerateToken(user *entity.User) (string, error) {
	jwtConf := conf.GetJWTConfig()
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserInfo: valobj.UserInfo{
			Id:       user.Id,
			NickName: user.NickName,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtConf.Expires).Unix(),
			Issuer:    jwtConf.Issuer,
		},
	})
	token, err := tokenClaims.SignedString(jwtConf.Secret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return conf.GetJWTConfig().Secret, nil
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
