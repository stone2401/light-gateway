package public

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stone2401/light-gateway/config"
)

type Claims struct {
	ID        int
	UserName  string
	LoginTime time.Time
	jwt.StandardClaims
}

// 生成token
func GenerateToken(id int, username string) (token string, err error) {
	expireTime := time.Now().Add(time.Minute * time.Duration(config.Config.JWT.TimeOut))
	claims := Claims{
		ID:        id,
		UserName:  username,
		LoginTime: time.Now(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Config.JWT.Issuer,
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Config.JWT.Key))
}

// 解析token
func ParseToken(token string) (claims *Claims, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWT.Key), nil
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
