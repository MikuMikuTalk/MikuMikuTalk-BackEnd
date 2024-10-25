package jwts

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtPayload struct {
	UserID   uint   `json:"userID"`
	Nickname string `json:"nickname"`
	Role     int8   `json:"role"` //1. 管理员 2.普通用户
}

type CustomClaims struct {
	JwtPayload
	jwt.RegisteredClaims
}

func GenerateJwtToken(payload JwtPayload, accessSecret string, expires int) (string, error) {
	cliams := CustomClaims{
		JwtPayload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	return token.SignedString([]byte(accessSecret))
}

func ParseToken(jwtToken string, accessSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("非法jwt token")
}
